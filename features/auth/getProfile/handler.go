package getprofile

import (
	"strconv"
	r "foodorder/infrastructure/repositories"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Get Profile godoc
//	@Summary
//	@Description	Get Logged user profile
//	@Tags			auth
//	@Produce		json
// @Security     BearerAuth
//	@Success		200		{object}	GetUserResponse
//	@Failure		404		{object}	map[string]string
//	@Router			/api/auth/profile [get]
func Profile(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userToken := c.Locals("user").(*jwt.Token)
        claims := userToken.Claims.(jwt.MapClaims)
        userID := uint(claims["sub"].(float64))

		user, err := r.GetUser(db, strconv.FormatUint(uint64(userID), 10))
        if err != nil {
            return c.Status(404).JSON(fiber.Map{"error": "User not found"})
        }

        return c.JSON(GetUserResponse{
            Id:    user.ID,
            Name:  user.Name,
            Email: user.Email,
        })
    }
}