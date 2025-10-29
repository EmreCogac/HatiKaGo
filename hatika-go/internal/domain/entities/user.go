package entities

import "time"

// User represents a system user
type User struct {
	FullAuditedEntity
	MultiTenantEntity
	
	Username            string     `gorm:"size:256;uniqueIndex;not null" json:"username" binding:"required"`
	Email               string     `gorm:"size:256;uniqueIndex;not null" json:"email" binding:"required,email"`
	PasswordHash        string     `gorm:"size:512;not null" json:"-"`
	Name                string     `gorm:"size:64" json:"name,omitempty"`
	Surname             string     `gorm:"size:64" json:"surname,omitempty"`
	IsActive            bool       `gorm:"default:true" json:"isActive"`
	EmailConfirmed      bool       `gorm:"default:false" json:"emailConfirmed"`
	PhoneNumber         string     `gorm:"size:32" json:"phoneNumber,omitempty"`
	PhoneNumberConfirmed bool      `gorm:"default:false" json:"phoneNumberConfirmed"`
	LockoutEnabled      bool       `gorm:"default:false" json:"lockoutEnabled"`
	LockoutEndDate      *time.Time `json:"lockoutEndDate,omitempty"`
	AccessFailedCount   int        `gorm:"default:0" json:"accessFailedCount"`
	
	// Navigation properties
	Roles               []Role     `gorm:"many2many:user_roles;" json:"roles,omitempty"`
}

// TableName overrides the table name
func (User) TableName() string {
	return "users"
}

// FullName returns the full name of the user
func (u *User) FullName() string {
	if u.Surname != "" {
		return u.Name + " " + u.Surname
	}
	return u.Name
}
