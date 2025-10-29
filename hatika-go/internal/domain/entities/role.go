package entities

// Role represents a user role for authorization
type Role struct {
	FullAuditedEntity
	MultiTenantEntity
	
	Name         string  `gorm:"size:128;uniqueIndex;not null" json:"name" binding:"required"`
	DisplayName  string  `gorm:"size:256;not null" json:"displayName" binding:"required"`
	Description  string  `gorm:"type:text" json:"description,omitempty"`
	IsStatic     bool    `gorm:"default:false" json:"isStatic"`
	IsDefault    bool    `gorm:"default:false" json:"isDefault"`
	
	// Navigation properties
	Users        []User  `gorm:"many2many:user_roles;" json:"users,omitempty"`
	Permissions  []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}

// TableName overrides the table name
func (Role) TableName() string {
	return "roles"
}

// Static role names
const (
	AdminRoleName = "Admin"
	UserRoleName  = "User"
)
