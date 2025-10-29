package dtos

// OcrProjectDto represents an OCR project data transfer object
type OcrProjectDto struct {
	FullAuditedEntityDto
	
	ProjectName             string   `json:"projectName,omitempty"`
	ProjectCode             string   `json:"projectCode,omitempty"`
	ProjectComment          string   `json:"projectComment,omitempty"`
	ProjectMuellef          string   `json:"projectMuellef,omitempty"`
	Ada                     *int     `json:"ada,omitempty"`
	Parsel                  *int     `json:"parsel,omitempty"`
	TalepGucu               *int     `json:"talepGucu,omitempty"`
	KuruluGuc               *int     `json:"kuruluGuc,omitempty"`
	BagimsizBS              *int     `json:"bagimsizBS,omitempty"`
	BlokS                   *int     `json:"blokS,omitempty"`
	YapiYuksekligi          *float64 `json:"yapiYuksekligi,omitempty"`
	RuhsatGecerlilikDate    string   `json:"ruhsatGecerlilikDate,omitempty"`
	YapiSahibi              string   `json:"yapiSahibi,omitempty"`
	Adress                  string   `json:"adress,omitempty"`
	Type                    int      `json:"type"`
	TypeName                string   `json:"typeName"`
	ProjectID               int      `json:"projectId"`
	PdfPath                 string   `json:"pdfPath,omitempty"`
}

// UpdateOcrProjectDto represents the input for updating an OCR project
type UpdateOcrProjectDto struct {
	ProjectName             string   `json:"projectName,omitempty"`
	ProjectCode             string   `json:"projectCode,omitempty"`
	ProjectComment          string   `json:"projectComment,omitempty"`
	ProjectMuellef          string   `json:"projectMuellef,omitempty"`
	Ada                     *int     `json:"ada,omitempty"`
	Parsel                  *int     `json:"parsel,omitempty"`
	TalepGucu               *int     `json:"talepGucu,omitempty"`
	KuruluGuc               *int     `json:"kuruluGuc,omitempty"`
	BagimsizBS              *int     `json:"bagimsizBS,omitempty"`
	BlokS                   *int     `json:"blokS,omitempty"`
	YapiYuksekligi          *float64 `json:"yapiYuksekligi,omitempty"`
	RuhsatGecerlilikDate    string   `json:"ruhsatGecerlilikDate,omitempty"`
	YapiSahibi              string   `json:"yapiSahibi,omitempty"`
	Adress                  string   `json:"adress,omitempty"`
	PdfPath                 string   `json:"pdfPath,omitempty"`
}

// PagedOcrProjectResultRequestDto represents paged request for OCR projects
type PagedOcrProjectResultRequestDto struct {
	PagedResultRequestDto
	
	ProjectID   int    `form:"projectId" json:"projectId,omitempty"`
	Type        int    `form:"type" json:"type,omitempty"`
	ProjectCode string `form:"projectCode" json:"projectCode,omitempty"`
}

// ProcessedDataModel represents processed OCR data
type ProcessedDataModel struct {
	FieldName  string      `json:"fieldName"`
	Value      interface{} `json:"value"`
	Confidence float64     `json:"confidence"`
}
