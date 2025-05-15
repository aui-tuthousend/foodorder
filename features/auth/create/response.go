package create

type CreateUserResponse struct {
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
}