package dtos

// LoginDto represents login credentials
type LoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResultDto represents the result of a login attempt
type LoginResultDto struct {
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken,omitempty"`
	ExpiresIn    int64    `json:"expiresIn"`
	User         UserDto  `json:"user"`
}

// RegisterDto represents user registration data
type RegisterDto struct {
	Username        string `json:"username" binding:"required,min=3,max=32"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	Name            string `json:"name" binding:"required"`
	Surname         string `json:"surname,omitempty"`
	TenancyName     string `json:"tenancyName,omitempty"`
}

// RegisterResultDto represents the result of registration
type RegisterResultDto struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	UserID  int    `json:"userId,omitempty"`
}

// ChangePasswordDto represents password change request
type ChangePasswordDto struct {
	CurrentPassword string `json:"currentPassword" binding:"required"`
	NewPassword     string `json:"newPassword" binding:"required,min=6"`
}

// ResetPasswordDto represents password reset request
type ResetPasswordDto struct {
	UserID      int    `json:"userId" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
}
