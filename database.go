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

// GetUserByEmail retrieves a user from the database by email
func GetUserByEmail(email string) (*User, error) {
    query := `SELECT id, name, email, password, created_at FROM app_users WHERE email = $1`
    user := &User{}
    err := db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Printf("User with email %s not found", email)
            return nil, nil // Return nil, nil to indicate not found, not an error
        }
        log.Printf("Error fetching user by email %s: %v", email, err)
        return nil, fmt.Errorf("database error fetching user: %w", err)
    }
    return user, nil
}