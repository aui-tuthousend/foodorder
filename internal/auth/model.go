package auth

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
}

type UserResponse struct {
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
}

// CreateUserRequestBody
// @Description Create user request body
type AuthRequest struct {
    // Your Fukin Name
    Name     string `json:"name"`
    // Your Fukin Email
    Email    string `json:"email"`
    // Your Fukin Password
    Password string `json:"password"`
}
