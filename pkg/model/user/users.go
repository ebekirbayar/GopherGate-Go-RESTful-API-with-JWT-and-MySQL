package user

import "gorm.io/gorm"

// User represents the user model stored in the database.
type User struct {
	gorm.Model
	UserEmail    string `gorm:"not null" json:"userEmail"`
	UserPassword string `gorm:"not null" json:"userPassword"`
	UserName     string `gorm:"not null" json:"userName"`
	UserRole     string `gorm:"not null" json:"userRole"`
	AccessToken  string `json:"accessToken,omitempty"`
}

// UserDTO represents the data transfer object for user-related operations.
type UserDTO struct {
	ID          uint
	UserEmail   string `gorm:"not null" json:"userEmail"`
	UserName    string `gorm:"not null" json:"userName"`
	UserRole    string `gorm:"not null" json:"userRole"`
	AccessToken string `json:"accessToken,omitempty"`
}

// UserLogin represents the user login model.
type UserLogin struct {
	UserEmail    string `json:"userEmail" validate:"required"`
	UserPassword string `json:"userPassword" validate:"required"`
}

// UserLoginDTO represents the data transfer object for user login.
type UserLoginDTO struct {
	ID          uint   `json:"id,omitempty"`
	UserEmail   string `json:"userEmail"`
	UserName    string `json:"userName"`
	UserRole    string `json:"userRole"`
	AccessToken string `json:"accessToken,omitempty"`
}

// Tabler is an interface for defining table names.
type Tabler interface {
	TableName() string
}

// TableName returns the table name for the UserRegister model.
func (UserRegister) TableName() string {
	return "users"
}

// UserRegister represents the user registration model stored in the database.
type UserRegister struct {
	gorm.Model
	UserEmail    string `json:"userEmail" gorm:"unique"`
	UserPassword string `json:"userPassword"`
	UserName     string `json:"userName" gorm:"unique"`
	UserRole     string `json:"userRole" gorm:"default:'user'"`
	AccessToken  string `json:"accessToken,omitempty"`
}
