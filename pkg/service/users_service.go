package service

import (
	"GopherGate/pkg/model"
	"GopherGate/pkg/model/user"
	"GopherGate/pkg/repository"
	"errors"
	"strconv"
)

// UsersService provides methods for user-related operations.
type UsersService struct {
	repository *repository.UsersRepository
}

// NewUsersService creates a new instance of UsersService.
func NewUsersService(repository *repository.UsersRepository) *UsersService {
	return &UsersService{
		repository: repository,
	}
}

// createErrorResponse creates an error response.
func createErrorResponse(message string, err error) *model.BaseErrorResponse {
	return &model.BaseErrorResponse{
		Message: message,
		Errors:  []error{err},
		Status:  false,
	}
}

// CreateUser creates a new user.
func (s *UsersService) CreateUser(user *user.UserRegister) (*model.BaseResponse, *model.BaseErrorResponse) {
	registerUser, err := s.repository.CreateUser(user)
	if err != nil {
		return nil, createErrorResponse("User could not be created!", err)
	}
	return registerUser, nil
}

// LoginUser handles user login.
func (s *UsersService) LoginUser(user *user.UserLogin) (*model.BaseResponse, *model.BaseErrorResponse) {
	loginUser, err := s.repository.LoginUser(user)
	if err != nil {
		return nil, createErrorResponse("User failed to log in!", err)
	}
	return loginUser, nil
}

// GetUsers retrieves all users.
func (s *UsersService) GetUsers(userRole string) (*[]user.UserDTOResponse, *model.BaseErrorResponse) {
	// Only admin users can retrieve all users
	if userRole != "admin" {
		return nil, createErrorResponse("You are not authorized to access this resource", nil)
	}

	// If the user is an admin, retrieve all users
	users, err := s.repository.GetUsers(userRole)
	if err != nil {
		return nil, createErrorResponse("Failed to fetch users!", err)
	}

	var usersDTO []user.UserDTOResponse
	for _, user := range *users {
		usersDTO = append(usersDTO, *user.ToDTO())
	}
	return &usersDTO, nil
}

// GetUser retrieves a user by ID.
func (s *UsersService) GetUser(id string) (*user.UserDTOResponse, error) {
	user, err := s.repository.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user.ToDTO(), nil
}

// UpdateUser updates a user.
func (s *UsersService) UpdateUser(updateUser *user.UserDTO, id string) (*user.UserDTOResponse, error) {
	updatedUser, err := s.repository.UpdateUser(updateUser, id)
	if err != nil {
		return nil, err
	}
	return updatedUser.ToDTO(), nil
}

// DeleteUser deletes a user.
func (s *UsersService) DeleteUser(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("Invalid user ID!")
	}
	if err := s.repository.DeleteUser(uint(intID)); err != nil {
		return err
	}
	return nil
}
