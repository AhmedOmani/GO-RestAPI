package main

import "time"

type User struct {
	ID						int					`json:"id"`
	Name					string				`json:"name"`
	Email					string				`json:"email"`
	Password				string				`json:"-"`
	CreatedAt				time.Time			`json:"created_at"`
}

type SignupRequest struct {
	Name					string				`json:"name" validate:"required,min=3,max=100"`
	Email					string 				`json:"email" validate:"required,email"`
	Password				string				`json:"password" validate:"required,min=8"`
}

type SigninRequest struct {
	Email 					string				`json:"email" validate:"required,email"`
	Password				string				`json:"password" validate:"required"`
}

type SigninResponse struct {
	Token 					string 				`json:"token"`
	Email					string				`json:"Email"`
	Name					string				`json:"name"`
}