package models

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"create_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Response struct {
	StatusCode int
	Message    string
	Error      error
}
