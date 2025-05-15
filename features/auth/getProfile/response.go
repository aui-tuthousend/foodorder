package getprofile

type GetUserResponse struct {
    Id     uint `json:"Id"`
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
}