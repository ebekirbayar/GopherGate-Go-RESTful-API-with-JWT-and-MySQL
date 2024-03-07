package user

// UserDTOResponse represents the response DTO for user-related operations.
type UserDTOResponse struct {
	Message string  `json:"message"`
	User    UserDTO `json:"user"`
	Status  bool    `json:"status"`
}

// ToDTO converts a UserRegister object to a UserDTOResponse object.
func (u *UserRegister) ToDTO() *UserDTOResponse {
	return &UserDTOResponse{
		Message: "User created successfully",
		User: UserDTO{
			ID:        u.ID,
			UserEmail: u.UserEmail,
			UserName:  u.UserName,
			UserRole:  u.UserRole,
		},
		Status: true,
	}
}

// ToDTO converts a User object to a UserDTOResponse object.
func (u *User) ToDTO() *UserDTOResponse {
	return &UserDTOResponse{
		Message: "User data retrieved successfully!",
		User: UserDTO{
			ID:          u.ID,
			UserEmail:   u.UserEmail,
			UserName:    u.UserName,
			UserRole:    u.UserRole,
			AccessToken: u.AccessToken,
		},
		Status: true,
	}
}

// ToDTO converts a UserLogin object to a UserLoginDTO object.
func (u *UserLogin) ToDTO() *UserLoginDTO {
	return &UserLoginDTO{
		UserEmail: u.UserEmail,
	}
}
