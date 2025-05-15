package create

import (
    e "foodorder/entities"
	r "foodorder/infrastructure/repositories"


	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Register godoc
//	@Summary		
//	@Description	Register a new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateUserRequest	true	"Create user request body"
//	@Success		200		{object}	CreateUserResponse
//	@Failure		500		{object}	map[string]string
//	@Router			/api/auth/register [post]
func Register(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var req CreateUserRequest
        if err := c.BodyParser(&req); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
        }

        hashedPassword, _ := r.HashPassword(req.Password)
        user := e.User{
            Name:     req.Name,
            Email:    req.Email,
            Password: hashedPassword,
        }

        if err := r.CreateUser(db, &user); err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Email already exists"})
        }
        return c.JSON(CreateUserResponse{Name: user.Name, Email: user.Email})
    }
}