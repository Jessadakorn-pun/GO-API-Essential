package domain

// Dummy User for Example
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var MemberUser = User{
	Email:    "user@example.com",
	Password: "password123",
}
