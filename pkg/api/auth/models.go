package auth

import "upload-service/pkg/api/users"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string             `json:"token"`
	User  users.UserResource `json:"user"`
}
