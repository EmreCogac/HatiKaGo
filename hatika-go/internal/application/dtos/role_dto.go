package dtos

// RoleDto represents a role data transfer object
type RoleDto struct {
	FullAuditedEntityDto
	
	Name        string          `json:"name"`
	DisplayName string          `json:"displayName"`
	Description string          `json:"description,omitempty"`
	IsStatic    bool            `json:"isStatic"`
	IsDefault   bool            `json:"isDefault"`
	Permissions []PermissionDto `json:"permissions,omitempty"`
}

// CreateRoleDto represents the input for creating a role
type CreateRoleDto struct {
	Name            string   `json:"name" binding:"required"`
	DisplayName     string   `json:"displayName" binding:"required"`
	Description     string   `json:"description,omitempty"`
	IsDefault       bool     `json:"isDefault"`
	PermissionNames []string `json:"permissionNames,omitempty"`
}

// UpdateRoleDto represents the input for updating a role
type UpdateRoleDto struct {
	DisplayName     string   `json:"displayName" binding:"required"`
	Description     string   `json:"description,omitempty"`
	IsDefault       bool     `json:"isDefault"`
	PermissionNames []string `json:"permissionNames,omitempty"`
}

// PagedRoleResultRequestDto represents paged request for roles
type PagedRoleResultRequestDto struct {
	PagedResultRequestDto
	
	Name string `form:"name" json:"name,omitempty"`
}
