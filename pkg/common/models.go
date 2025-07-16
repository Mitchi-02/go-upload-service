package common

type PaginatedResponse struct {
	Results interface{} `json:"results"`
	Total   int         `json:"total"`
}

type APIResponse struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

type JWTClaims struct {
	UserID string `json:"user_id"`
	Exp    int64  `json:"exp"`
}

const (
	UserIDContextKey = "userID"
)
