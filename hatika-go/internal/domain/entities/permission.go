package entities

// Permission represents a system permission
type Permission struct {
	BaseEntity
	
	Name         string  `gorm:"size:128;uniqueIndex;not null" json:"name" binding:"required"`
	DisplayName  string  `gorm:"size:256;not null" json:"displayName" binding:"required"`
	Description  string  `gorm:"type:text" json:"description,omitempty"`
	
	// Navigation properties
	Roles        []Role  `gorm:"many2many:role_permissions;" json:"roles,omitempty"`
}

// TableName overrides the table name
func (Permission) TableName() string {
	return "permissions"
}

// Permission names (similar to PermissionNames.cs)
const (
	// Pages
	PagesUsers      = "Pages.Users"
	PagesRoles      = "Pages.Roles"
	PagesProjects   = "Pages.Projects"
	PagesOcrProjects = "Pages.OcrProjects"
	PagesTenants    = "Pages.Tenants"
	
	// Actions
	UsersCreate     = "Pages.Users.Create"
	UsersEdit       = "Pages.Users.Edit"
	UsersDelete     = "Pages.Users.Delete"
	
	RolesCreate     = "Pages.Roles.Create"
	RolesEdit       = "Pages.Roles.Edit"
	RolesDelete     = "Pages.Roles.Delete"
	
	ProjectsCreate  = "Pages.Projects.Create"
	ProjectsEdit    = "Pages.Projects.Edit"
	ProjectsDelete  = "Pages.Projects.Delete"
)
