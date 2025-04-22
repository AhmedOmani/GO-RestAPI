package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	err := godotenv.Load();
	if err != nil {
		log.Println("Warning: Could not load .env file")
    }

	connStr := os.Getenv("DATABASE_URL");

	db , err = sql.Open("postgres" , connStr)
	if err != nil {
		log.Fatal("Failed to open database connection:", err)
	}

	err = db.Ping() 
	if err != nil {
		log.Fatal("Failed to ping database: " , err)
	}

	log.Println("Database connection established successfully.")
}

//TODO: Create new user in database
func CreateUser(user *User) (int , error) {
	query := `INSERT INTO app_users (name , email , password) VALUES ($1 , $2 , $3) RETURNING id`;
	var id int ;
	err := db.QueryRow(query , user.Name , user.Email , user.Password).Scan(&id);
	fmt.Println(err);
	if err != nil {
		return 0 , fmt.Errorf("could not create user: %w", err);
	}
	user.ID = id ;
	return id , nil ;
}

func GetUserByEmail(email string) (bool , error) {
	var exists bool ;
	query := `SELECT 1 FROM app_users WHERE email = $1 LIMIT 1`;
	err := db.QueryRow(query , email).Scan(&exists);
	fmt.Println(err);
	if err != nil {
		return false , fmt.Errorf("Error while checking if email exists: %w" , err);
	}
	return exists , nil ;
}


