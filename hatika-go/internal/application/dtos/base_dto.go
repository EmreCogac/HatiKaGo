package dtos

import "time"

// PagedResultRequestDto is the base class for paged request DTOs
type PagedResultRequestDto struct {
	PageNumber int    `form:"pageNumber" json:"pageNumber" binding:"required,min=1"`
	PageSize   int    `form:"pageSize" json:"pageSize" binding:"required,min=1,max=100"`
	Sorting    string `form:"sorting" json:"sorting,omitempty"`
}

// PagedResultDto represents a paged result
type PagedResultDto[T any] struct {
	TotalCount int `json:"totalCount"`
	Items      []T `json:"items"`
}

// EntityDto represents a basic entity DTO
type EntityDto struct {
	ID int `json:"id"`
}

// AuditedEntityDto includes audit information
type AuditedEntityDto struct {
	EntityDto
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	CreatorUserID  *int      `json:"creatorUserId,omitempty"`
	LastModifierID *int      `json:"lastModifierId,omitempty"`
}

// FullAuditedEntityDto includes all audit fields
type FullAuditedEntityDto struct {
	AuditedEntityDto
	DeleterUserID *int       `json:"deleterUserId,omitempty"`
	DeletionTime  *time.Time `json:"deletionTime,omitempty"`
	IsDeleted     bool       `json:"isDeleted"`
}
