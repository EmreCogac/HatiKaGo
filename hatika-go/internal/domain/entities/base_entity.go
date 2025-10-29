package entities

import (
	"time"
)

type BaseEntity struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

type FullAuditedEntity struct {
	BaseEntity
	CreatorUserID  *int       `gorm:"index" json:"creatorUserId,omitempty"`
	LastModifierID *int       `gorm:"index" json:"lastModifierId,omitempty"`
	DeleterUserID  *int       `gorm:"index" json:"deleterUserId,omitempty"`
	DeletionTime   *time.Time `json:"deletionTime,omitempty"`
	IsDeleted      bool       `gorm:"default:false;index" json:"isDeleted"`
}

func (e *FullAuditedEntity) SoftDelete(userID int) {
	now := time.Now()
	e.IsDeleted = true
	e.DeletionTime = &now
	e.DeleterUserID = &userID
}

type MultiTenantEntity struct {
	TenantID *int `gorm:"index" json:"tenantId,omitempty"`
}

type IMayHaveTenant interface {
	GetTenantID() *int
	SetTenantID(tenantID *int)
}

func (e *MultiTenantEntity) GetTenantID() *int {
	return e.TenantID
}

func (e *MultiTenantEntity) SetTenantID(tenantID *int) {
	e.TenantID = tenantID
}
