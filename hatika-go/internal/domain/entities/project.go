package entities

// Project represents a construction project
type Project struct {
	FullAuditedEntity
	MultiTenantEntity
	
	ProjectName             string        `gorm:"size:255;not null" json:"projectName" binding:"required"`
	ProjectCode             string        `gorm:"size:100;uniqueIndex" json:"projectCode" binding:"required"`
	ProjectComment          string        `gorm:"type:text" json:"projectComment,omitempty"`
	ProjectMuellef          string        `gorm:"size:255" json:"projectMuellef,omitempty"`
	Ada                     *int          `json:"ada,omitempty"`
	Parsel                  *int          `json:"parsel,omitempty"`
	TalepGucu               *int          `json:"talepGucu,omitempty"`
	KuruluGuc               *int          `json:"kuruluGuc,omitempty"`
	BagimsizBS              *int          `json:"bagimsizBS,omitempty"`
	BlokS                   *int          `json:"blokS,omitempty"`
	YapiYuksekligi          *float64      `json:"yapiYuksekligi,omitempty"`
	RuhsatGecerlilikDate    string        `gorm:"size:50" json:"ruhsatGecerlilikDate,omitempty"`
	YapiSahibi              string        `gorm:"size:255" json:"yapiSahibi,omitempty"`
	Adress                  string        `gorm:"type:text" json:"adress,omitempty"`
	GroupID                 *int          `gorm:"index" json:"groupId,omitempty"`
	BildirimNo              string        `gorm:"size:100" json:"bildirimNo,omitempty"`
	
	// Navigation property
	OcrProjects             []OcrProject  `gorm:"foreignKey:ProjectID" json:"ocrProjects,omitempty"`
}

// TableName overrides the table name
func (Project) TableName() string {
	return "projects"
}
