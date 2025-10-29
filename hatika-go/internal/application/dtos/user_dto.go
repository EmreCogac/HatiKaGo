package dtos

// UserDto represents a user data transfer object
type UserDto struct {
	FullAuditedEntityDto

	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Name           string    `json:"name"`
	Surname        string    `json:"surname"`
	FullName       string    `json:"fullName"`
	IsActive       bool      `json:"isActive"`
	EmailConfirmed bool      `json:"emailConfirmed"`
	PhoneNumber    string    `json:"phoneNumber,omitempty"`
	Roles          []RoleDto `json:"roles,omitempty"`
}

type CreateUserDto struct {
	Username    string   `json:"username" binding:"required,min=3,max=32"`
	Email       string   `json:"email" binding:"required,email"`
	Password    string   `json:"password" binding:"required,min=6"`
	Name        string   `json:"name" binding:"required"`
	Surname     string   `json:"surname,omitempty"`
	PhoneNumber string   `json:"phoneNumber,omitempty"`
	IsActive    bool     `json:"isActive"`
	RoleNames   []string `json:"roleNames,omitempty"`
}

type UpdateUserDto struct {
	Username    string   `json:"username" binding:"required,min=3,max=32"`
	Email       string   `json:"email" binding:"required,email"`
	Name        string   `json:"name" binding:"required"`
	Surname     string   `json:"surname,omitempty"`
	PhoneNumber string   `json:"phoneNumber,omitempty"`
	IsActive    bool     `json:"isActive"`
	RoleNames   []string `json:"roleNames,omitempty"`
}

type PagedUserResultRequestDto struct {
	PagedResultRequestDto

	Username string `form:"username" json:"username,omitempty"`
	Email    string `form:"email" json:"email,omitempty"`
	IsActive *bool  `form:"isActive" json:"isActive,omitempty"`
}
