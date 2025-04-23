package main

import (
	"encoding/json"
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

	if isUserExist != nil {
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

func SigninHandler(w http.ResponseWriter , r *http.Request){
	if r.Method != http.MethodPost {
		respondError(w , http.StatusMethodNotAllowed , "Only POST request allowed for this endpoint");
		return ;
	}

	var req SigninRequest ;
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

	//Fetch user by email
	user , err := GetUserByEmail(req.Email);
	if err != nil {
		respondError(w , http.StatusInternalServerError , "Database Error during login");
		return ;
	}
	if user == nil {
		respondError(w , http.StatusUnauthorized , "Invalid email or password");
		return ;
	}

	//Check password
	if !CheckPasswordHash(req.Password , user.Password) {
		respondError(w , http.StatusUnauthorized , "Invalide email or password");
		return ;
	}

	//Generate JWT Token
	tokenString , err := GenerateJWT(user);
	if err != nil {
		respondError(w, http.StatusInternalServerError , "Failed to generated authentication token");
		return ;
	}

	log.Printf("User signed in successfully");
	response := &SigninResponse{
		Token:		tokenString,
		Email:		user.Email,
		Name: 		user.Name,
	}
	respondJSON(w , http.StatusOK , response);
}