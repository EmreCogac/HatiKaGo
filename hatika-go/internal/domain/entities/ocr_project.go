package entities

type OcrProjectType int

const (
	ProjeAntenti OcrProjectType = iota
	YapiRuhsati
	YapiKullanimBelgesi
	Tapu
)

func (t OcrProjectType) String() string {
	return [...]string{"ProjeAntenti", "YapiRuhsati", "YapiKullanimBelgesi", "Tapu"}[t]
}

type OcrProject struct {
	FullAuditedEntity
	MultiTenantEntity

	ProjectName          string         `gorm:"size:255" json:"projectName,omitempty"`
	ProjectCode          string         `gorm:"size:100" json:"projectCode,omitempty"`
	ProjectComment       string         `gorm:"type:text" json:"projectComment,omitempty"`
	ProjectMuellef       string         `gorm:"size:255" json:"projectMuellef,omitempty"`
	Ada                  *int           `json:"ada,omitempty"`
	Parsel               *int           `json:"parsel,omitempty"`
	TalepGucu            *int           `json:"talepGucu,omitempty"`
	KuruluGuc            *int           `json:"kuruluGuc,omitempty"`
	BagimsizBS           *int           `json:"bagimsizBS,omitempty"`
	BlokS                *int           `json:"blokS,omitempty"`
	YapiYuksekligi       *float64       `json:"yapiYuksekligi,omitempty"`
	RuhsatGecerlilikDate string         `gorm:"size:50" json:"ruhsatGecerlilikDate,omitempty"`
	YapiSahibi           string         `gorm:"size:255" json:"yapiSahibi,omitempty"`
	Adress               string         `gorm:"type:text" json:"adress,omitempty"`
	Type                 OcrProjectType `gorm:"type:int;not null" json:"type"`
	ProjectID            int            `gorm:"not null;index" json:"projectId"`
	PdfPath              string         `gorm:"size:500" json:"pdfPath,omitempty"`

	Project *Project `gorm:"foreignKey:ProjectID" json:"project,omitempty"`
}

func (OcrProject) TableName() string {
	return "ocr_projects"
}
