package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func SignupHandler(w http.ResponseWriter , r *http.Request){
	if r.Method != http.MethodPost {
		respondError(w , http.StatusMethodNotAllowed , "Only POST method for this endpoint");
		return;
	}

	var req SignupRequest ;
	
	//Handling the validation errors that may occur
	if err := decodeAndValidate(r , &req); err != nil {
		if ve , ok := err.(ValidationError); ok {
			w.WriteHeader(http.StatusBadRequest);
			json.NewEncoder(w).Encode(map[string]interface{} {
				"error" : "Validation failed",
				"detailed" : ve.Errors ,
			})
			return ;
		}

		w.WriteHeader(http.StatusBadRequest);
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		});
		return
	}

	isUserExist , err := GetUserByEmail(req.Email);
	if err != nil {
		respondError(w , http.StatusInternalServerError , "Database error while cheking user");
		return ;
	}

	fmt.Println("User exist");
	fmt.Println(isUserExist);

	if isUserExist == true {
		respondError(w , http.StatusConflict , "Email address is already registered");
		return ;
	}

	//Hash the password 
	hashedPassword , err := HashPassword(req.Password);

	newUser := &User {
		Name: 			req.Name,
		Email:			req.Email ,
		Password: 		hashedPassword,		
	}

	userID , err := CreateUser(newUser);
	if err != nil {
		respondError(w , http.StatusInternalServerError , "Failed to create user , Try Again!");
		return ;
	}

	log.Printf("User registered successfully");
	respondJSON(w , http.StatusCreated , map[string]string{"message" : "User created successfully" , "userID" : strconv.Itoa(userID)});
}

func SigninHandler(w http.ResponseWriter , r *http.Request){}