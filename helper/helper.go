package helper

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetUserID extracts the user ID from the Fiber context and converts it to an unsigned integer.
func GetUserID(c *fiber.Ctx) (uint, error) {
	id := c.Get("userID")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return uint(userID), nil
}
