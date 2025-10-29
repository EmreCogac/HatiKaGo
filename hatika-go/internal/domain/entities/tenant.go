package entities

// Tenant represents a tenant in multi-tenant system
type Tenant struct {
	FullAuditedEntity
	
	TenancyName      string `gorm:"size:128;uniqueIndex;not null" json:"tenancyName" binding:"required"`
	Name             string `gorm:"size:256;not null" json:"name" binding:"required"`
	ConnectionString string `gorm:"size:1024" json:"connectionString,omitempty"`
	IsActive         bool   `gorm:"default:true" json:"isActive"`
	EditionID        *int   `gorm:"index" json:"editionId,omitempty"`
}

// TableName overrides the table name
func (Tenant) TableName() string {
	return "tenants"
}
