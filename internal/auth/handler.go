package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)


// Register godoc
//	@Summary		Register a new user
//	@Description	Register a new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		AuthRequest	true	"Create user request body"
//	@Success		200		{object}	UserResponse
//	@Failure		500		{string}	error
//	@Router			/auth/register [post]
func Register(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var req AuthRequest
        if err := c.BodyParser(&req); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
        }

        hashedPassword, _ := HashPassword(req.Password)
        user := User{
            Name:     req.Name,
            Email:    req.Email,
            Password: hashedPassword,
        }

        if err := CreateUser(db, &user); err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Email already exists"})
        }

        token, _ := GenerateJWT(user.ID)
        return c.JSON(fiber.Map{"token": token})
    }
}

func Login(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var req AuthRequest
        if err := c.BodyParser(&req); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
        }

        user, err := FindUserByEmail(db, req.Email)
        if err != nil || !CheckPasswordHash(req.Password, user.Password) {
            return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
        }

        token, _ := GenerateJWT(user.ID)
        return c.JSON(fiber.Map{"token": token})
    }
}

func Profile(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userToken := c.Locals("user").(*jwt.Token)
        claims := userToken.Claims.(jwt.MapClaims)
        userID := uint(claims["sub"].(float64))

        var user User
        if err := db.First(&user, userID).Error; err != nil {
            return c.Status(404).JSON(fiber.Map{"error": "User not found"})
        }

        return c.JSON(fiber.Map{
            "id":    user.ID,
            "name":  user.Name,
            "email": user.Email,
        })
    }
}
