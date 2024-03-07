package api

import (
	"GopherGate/pkg/model"
	"GopherGate/pkg/model/user"
	"GopherGate/pkg/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// UsersAPI is a struct that represents the API endpoints related to users.
type UsersAPI struct {
	usersService *service.UsersService
}

// NewUsersAPI creates a new instance of the UsersAPI.
func NewUsersAPI(service *service.UsersService) *UsersAPI {
	return &UsersAPI{
		usersService: service,
	}
}

// CreateUser handles the creation of a new user.
func (api *UsersAPI) CreateUser(c *fiber.Ctx) error {
	user := new(user.UserRegister)
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	createdUser, err := api.usersService.CreateUser(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(createdUser)
}

// LoginUser handles the user login process.
func (api *UsersAPI) LoginUser(c *fiber.Ctx) error {
	user := new(user.UserLogin)
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	loginUser, err := api.usersService.LoginUser(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(loginUser)
}

// GetUsers retrieves all users based on the user's role.
func (api *UsersAPI) GetUsers(c *fiber.Ctx) error {
	// Retrieve user role from JWT token
	userRole, ok := c.Locals("userRole").(string)
	if !ok || userRole != "admin" {
		return c.Status(http.StatusUnauthorized).JSON(&model.BaseErrorResponse{
			Message: "Unauthorized access",
			Status:  false,
		})
	}

	// If user role is admin, retrieve all users
	users, err := api.usersService.GetUsers(userRole)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&model.BaseErrorResponse{
			Message: "Failed to fetch users",
			Status:  false,
		})
	}
	return c.Status(http.StatusOK).JSON(users)
}

// GetUser retrieves a specific user based on the user ID.
func (api *UsersAPI) GetUser(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(http.StatusUnauthorized).JSON(&model.BaseErrorResponse{
			Message: "User ID not found in JWT",
			Status:  false,
		})
	}

	userDTO, err := api.usersService.GetUser(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&model.BaseErrorResponse{
			Message: "Failed to fetch user!",
			Errors:  []error{err},
			Status:  false,
		})
	}
	return c.Status(http.StatusOK).JSON(userDTO)
}

// UpdateUser updates user information.
func (api *UsersAPI) UpdateUser(c *fiber.Ctx) error {
	userRole, ok := c.Locals("userRole").(string)
	if !ok || userRole != "admin" {
		return c.Status(http.StatusUnauthorized).JSON(&model.BaseErrorResponse{
			Message: "Unauthorized access!",
			Status:  false,
		})
	}

	id := c.Params("id")

	// Gelen JSON verisini user.UserDTO'ya çevir
	updateUser := new(user.UserDTO)
	if err := c.BodyParser(updateUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&model.BaseErrorResponse{
			Message: "Invalid request body!",
			Status:  false,
		})
	}

	// Kullanıcı güncelleme işlemini yap
	updatedUser, err := api.usersService.UpdateUser(updateUser, id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&model.BaseErrorResponse{
			Message: "Failed to update user: api hatası",
			Status:  false,
		})
	}

	return c.Status(http.StatusOK).JSON(updatedUser)
}

// DeleteUser deletes a user.
func (api *UsersAPI) DeleteUser(c *fiber.Ctx) error {

	userID := c.Locals("userID").(string)
	requestedUserID := c.Params("id")

	userRole, ok := c.Locals("userRole").(string)
	if !ok || userRole != "admin" {
		return c.Status(http.StatusUnauthorized).JSON(&model.BaseErrorResponse{
			Message: "Only admin users can delete users!",
			Status:  false,
		})
	}

	if userID == requestedUserID {
		return c.Status(http.StatusBadRequest).JSON(&model.BaseErrorResponse{
			Message: "Admin cannot delete their own account!",
			Status:  false,
		})
	}

	if err := api.usersService.DeleteUser(requestedUserID); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&model.BaseErrorResponse{
			Message: err.Error(),
			Status:  false,
		})
	}

	// Başarılı silme işlemini bildir ve silinen kullanıcının bilgilerini göster
	response := &model.BaseResponse{
		Message: "User successfully deleted!",
		Status:  true,
	}
	return c.Status(http.StatusOK).JSON(response)
}
