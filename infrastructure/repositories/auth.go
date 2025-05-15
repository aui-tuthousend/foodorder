package repositories

import (
    e "foodorder/entities"
    "os"
    "time"

    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GenerateJWT(userID uint) (string, error) {
    claims := jwt.MapClaims{
        "sub": userID,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func FindUserByEmail(db *gorm.DB, email string) (*e.User, error) {
    var user e.User
    result := db.Raw("SELECT * FROM users WHERE email = ? LIMIT 1", email).Scan(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func GetUser(db *gorm.DB, id string) (*e.User, error) {
    var user e.User
    result := db.Raw("SELECT * FROM users WHERE id = ? LIMIT 1", id).Scan(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

func CreateUser(db *gorm.DB, user *e.User) error {
    result := db.Exec(`
        INSERT INTO users (name, email, password, created_at, updated_at)
        VALUES (?, ?, ?, NOW(), NOW())
    `, user.Name, user.Email, user.Password)

    return result.Error
}