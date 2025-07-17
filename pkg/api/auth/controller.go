package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"upload-service/configs"
	"upload-service/pkg/api/users"
	"upload-service/pkg/common"

	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	repo      *users.Repository
	appConfig configs.AppConfig
}

func AuthController(db *sql.DB, appConfig configs.AppConfig) *Controller {
	return &Controller{repo: users.UserRepository(db), appConfig: appConfig}
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Invalid request body"})
		return
	}

	if req.Email == "" || req.Password == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Email and password are required"})
		return
	}

	user, err := c.repo.GetByEmail(req.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Invalid credentials"})
		return
	}

	token, err := common.GenerateJWT(user.ID, c.appConfig.JWTSecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Failed to generate token"})
		return
	}

	response := common.APIResponse{
		Data: AuthResponse{
			Token: token,
			User: users.UserResource{
				ID:    user.ID,
				Email: user.Email,
			},
		},
	}

	json.NewEncoder(w).Encode(response)
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Invalid request body"})
		return
	}

	if req.Email == "" || req.Password == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Email and password are required"})
		return
	}

	if _, err := c.repo.GetByEmail(req.Email); err == nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "User already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Failed to hash password"})
		return
	}

	user := users.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := c.repo.Create(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Failed to create user"})
		return
	}

	token, err := common.GenerateJWT(user.ID, c.appConfig.JWTSecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(common.APIResponse{Error: "Failed to generate token"})
		return
	}

	response := common.APIResponse{
		Data: AuthResponse{
			Token: token,
			User: users.UserResource{
				ID:    user.ID,
				Email: user.Email,
			},
		},
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
