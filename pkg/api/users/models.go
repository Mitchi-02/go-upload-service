package users

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// api response model
type UserResource struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
