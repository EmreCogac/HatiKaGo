package dtos

// PermissionDto represents a permission data transfer object
type PermissionDto struct {
	EntityDto
	
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Description string `json:"description,omitempty"`
}

// FlatPermissionDto represents a flat permission structure
type FlatPermissionDto struct {
	ParentName  string `json:"parentName,omitempty"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Description string `json:"description,omitempty"`
	IsGranted   bool   `json:"isGranted"`
}
