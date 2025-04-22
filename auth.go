package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret []byte 

func InitAuth() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file")
	}
	secret := os.Getenv("JWT_SECRET");
	jwtSecret = []byte(secret);
}

func HashPassword(password string) (string , error) {
	bytes , err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost);
	if err != nil {
		log.Printf("Error generating bcrypt hash: %v", err);
        return "", fmt.Errorf("could not hash password: %w", err);
	}
	return string(bytes) , nil ;
}

func CheckPasswordHash(password , hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash) , []byte(password));
	return err == nil ;
}

func GenerateJWT(user *User) (string , error) {
	//Create the claim
	claims := jwt.MapClaims {
		"sub":			user.ID,
		"email":		user.Email,
		"name":			user.Name,	
		"exp": 			time.Now().Add(time.Hour * 24).Unix(),
		"iat":			time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256 , claims);
	tokenString , err := token.SignedString(jwtSecret);

	if err != nil {
        log.Printf("Error signing JWT for user %s: %v", user.Email, err)
        return "", fmt.Errorf("could not generate token: %w", err)
    }

	return tokenString , nil ;
} 

