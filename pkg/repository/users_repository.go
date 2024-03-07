package repository

import (
	"GopherGate/pkg/model"
	"GopherGate/pkg/model/user"
	"errors"

	"gorm.io/gorm"
)

var (
	usersTable = "users"
)

// UsersRepository handles database operations related to users.
type UsersRepository struct {
	DB *gorm.DB
}

// NewUsersRepository creates a new instance of UsersRepository.
func NewUsersRepository(db *gorm.DB) *UsersRepository {
	return &UsersRepository{
		DB: db,
	}
}

// CreateUser creates a new user.
func (u *UsersRepository) CreateUser(user *user.UserRegister) (*model.BaseResponse, error) {
	hashedPass, err := HashPassword(user.UserPassword)
	if err != nil {
		return nil, errors.New("error hashing password: " + err.Error())
	}
	user.UserPassword = hashedPass
	if err := u.DB.Create(&user).Error; err != nil {
		return nil, errors.New("error creating user: " + err.Error())
	}
	user.UserPassword = ""
	return &model.BaseResponse{
		Message: "User created successfully!",
		Data:    user,
		Status:  true,
	}, nil
}

// LoginUser performs user login.
func (u *UsersRepository) LoginUser(login *user.UserLogin) (*model.BaseResponse, error) {
	if login.UserEmail == "" || login.UserPassword == "" {
		return nil, errors.New("please provide email and password")
	}
	var loginedUser *user.UserRegister
	err := u.DB.Table("users").Where("user_email = ?", login.UserEmail).First(&loginedUser).Error
	if err != nil {
		return nil, errors.New("error finding user")
	}

	if !CheckPasswordHash(login.UserPassword, loginedUser.UserPassword) {
		return nil, errors.New("invalid login credentials")
	}

	accessToken, err := CreateToken(loginedUser.UserName, loginedUser.ID, loginedUser.UserRole)

	if err != nil {
		return nil, errors.New("error creating accessToken")
	}

	if err := u.DB.Table("users").Where("user_email = ?", login.UserEmail).Update("access_token", accessToken).Error; err != nil {
		return nil, errors.New("error updating user accessToken")
	}

	dto := &user.UserLoginDTO{
		UserEmail:   login.UserEmail,
		UserName:    loginedUser.UserName,
		UserRole:    loginedUser.UserRole,
		AccessToken: *accessToken,
	}
	baseResponse := model.BaseResponse{
		Message: "User logged in successfully!",
		Data:    dto,
		Status:  true,
	}
	return &baseResponse, nil
}

// GetUsers retrieves all users (admin only).
func (u *UsersRepository) GetUsers(userRole string) (*[]user.User, error) {
	if userRole != "admin" {
		return nil, errors.New("you are not authorized to access this resource")
	}

	var users []user.User
	err := u.DB.Table(usersTable).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

// GetUser retrieves a specific user.
func (u *UsersRepository) GetUser(id string) (*user.User, error) {
	var user user.User
	err := u.DB.Table(usersTable).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user.
func (u *UsersRepository) UpdateUser(updateUser *user.UserDTO, id string) (*user.User, error) {
	existingUser := &user.User{}
	if err := u.DB.First(existingUser, id).Error; err != nil {
		return nil, errors.New("user not found repo hatası")
	}

	// Güncellenmiş kullanıcı bilgilerini oluştur
	updatedUser := &user.User{
		Model:     existingUser.Model,
		UserEmail: updateUser.UserEmail,
		UserName:  updateUser.UserName,
		UserRole:  updateUser.UserRole,
	}

	// Kullanıcıyı güncelle
	if err := u.DB.Model(existingUser).Updates(updatedUser).Error; err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// DeleteUser deletes a user.
func (u *UsersRepository) DeleteUser(id uint) error {
	err := u.DB.Delete(&user.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
