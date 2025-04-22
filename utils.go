package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitUtils() {
	validate = validator.New();
}

func respondJSON(w http.ResponseWriter , statusCode int , payload interface{}) {
	response , err := json.Marshal(payload);
	
	if err != nil {
		log.Printf("Error marshaling JSON response: %v", err);
		w.WriteHeader(http.StatusInternalServerError);
		w.Write([]byte(`{"error": "Internal Server Error"}`));
		return ;
	}

	w.Header().Set("Content-Type" , "application/json");
	w.WriteHeader(statusCode);
	w.Write(response);
}

func respondError(w http.ResponseWriter , statusCode int , message string) {
	log.Printf("Responding with error statusCode: %d message: %s \n" , statusCode , message);
	respondJSON(w , statusCode , map[string]string{"error" : message});
}

type ValidationError struct {
	Errors map[string]string
}

func (e ValidationError) Error() string {
	return "validation error"
}

func getValidationErrorMsg(fe validator.FieldError) string {
	switch fe.Field() {
	case "Name":
		switch fe.Tag() {
			case "required":
				return "Name is rquired."
			case "min":
				return "Name must be atleast 3 characters long."
			case "max":
				return "Name must be less than 100 characters long."
		}
	case "Email":
		switch fe.Tag() {
			case "required":
				return "Email is required."
			case "email":
				return "Email must be a valid email address."
		}
	case "Password":
		switch fe.Tag() {
		case "required":
			return "Password is required."
		case "min":
			return "Password must be atleast 8 characters long."
		}
	}
	return fe.Field() + " is not valid";
}

func decodeAndValidate(r *http.Request , v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return fmt.Errorf("invalid request payload: %w", err)
	}

	defer r.Body.Close();

	if err := validate.Struct(v) ; err != nil {
		if validationErrors , ok := err.(validator.ValidationErrors); ok {
			errMap := make(map[string]string);
			for _ , ve := range validationErrors {
				errMap[ve.Field()] = getValidationErrorMsg(ve)
			}
			return ValidationError{Errors: errMap}
		}
		return err ;
	}

	return nil ;
}