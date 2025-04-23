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
	cachedUser , err := GetUserInCache(email);
	if err != nil {
		log.Printf("Warning: Redis cache read error for email %s: %v. Attempting DB.", email, err);
	} 

	//Cache Hit , return the user cached in redis.
	//NOTE : 
	//We ignore to store hashed password in redis so it is empty 
	//and we have to add it manually by fetching from db 
	if cachedUser != nil {
		var passwordHash string;
		query := `SELECT password FROM app_users WHERE email = $1`;
		err := db.QueryRow(query , email).Scan(&passwordHash);
		if err != nil {
			return nil , fmt.Errorf("Error fetching password hash: %w" , err);
		}
		cachedUser.Password = passwordHash;
		return cachedUser , nil ;
	}

	//Cache Miss , lets get the user from DB .
    query := `SELECT id, name, email, password, created_at FROM app_users WHERE email = $1`;
    user := &User{};
    dbErr := db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt);
	fmt.Println("LOL:");
	fmt.Println(dbErr);
    if dbErr != nil {
        if dbErr == sql.ErrNoRows {
            log.Printf("User with email %s is not exist", email)
            return nil, nil // Return nil, nil to indicate not exist, not an error
        }
        log.Printf("Error fetching user by email %s: %v", email, err)
        return nil, fmt.Errorf("database error fetching user: %w", err)
    }

	//Now store the user in redis cache.
	err = SetUserInCache(user);
	if err != nil {
		log.Printf("Failed to SET user in cached after DB fetch: %v" , err);
	}

    return user, nil
}