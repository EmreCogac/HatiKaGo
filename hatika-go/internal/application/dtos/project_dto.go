package dtos

// ProjectDto represents a project data transfer object
type ProjectDto struct {
	FullAuditedEntityDto
	
	ProjectName             string        `json:"projectName" binding:"required"`
	ProjectCode             string        `json:"projectCode" binding:"required"`
	ProjectComment          string        `json:"projectComment,omitempty"`
	ProjectMuellef          string        `json:"projectMuellef,omitempty"`
	Ada                     *int          `json:"ada,omitempty"`
	Parsel                  *int          `json:"parsel,omitempty"`
	TalepGucu               *int          `json:"talepGucu,omitempty"`
	KuruluGuc               *int          `json:"kuruluGuc,omitempty"`
	BagimsizBS              *int          `json:"bagimsizBS,omitempty"`
	BlokS                   *int          `json:"blokS,omitempty"`
	YapiYuksekligi          *float64      `json:"yapiYuksekligi,omitempty"`
	RuhsatGecerlilikDate    string        `json:"ruhsatGecerlilikDate,omitempty"`
	YapiSahibi              string        `json:"yapiSahibi,omitempty"`
	Adress                  string        `json:"adress,omitempty"`
	GroupID                 *int          `json:"groupId,omitempty"`
	BildirimNo              string        `json:"bildirimNo,omitempty"`
	OcrProjects             []OcrProjectDto `json:"ocrProjects,omitempty"`
}

// CreateProjectDto represents the input for creating a project
type CreateProjectDto struct {
	ProjectName             string   `json:"projectName" binding:"required"`
	ProjectCode             string   `json:"projectCode" binding:"required"`
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
	GroupID                 *int     `json:"groupId,omitempty"`
	BildirimNo              string   `json:"bildirimNo,omitempty"`
}

// UpdateProjectDto represents the input for updating a project
type UpdateProjectDto struct {
	CreateProjectDto
}

// PagedProjectResultRequestDto represents paged request for projects with filters
type PagedProjectResultRequestDto struct {
	PagedResultRequestDto
	
	GroupID        int      `form:"groupId" json:"groupId,omitempty"`
	BildirimNo     string   `form:"bildirimNo" json:"bildirimNo,omitempty"`
	ProjectCode    string   `form:"projectCode" json:"projectCode,omitempty"`
	ProjectName    string   `form:"projectName" json:"projectName,omitempty"`
	ProjectMuellef string   `form:"projectMuellef" json:"projectMuellef,omitempty"`
	IdList         []int    `form:"idList" json:"idList,omitempty"`
}
