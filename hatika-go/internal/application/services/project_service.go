package services

import (
	"context"
	"fmt"

	"hatika-go/internal/application/dtos"
	"hatika-go/internal/domain/entities"
	"hatika-go/internal/infrastructure/persistence"
)

// ProjectService handles project business logic
type ProjectService struct {
	projectRepo *persistence.ProjectRepository
}

// NewProjectService creates a new project service
func NewProjectService(projectRepo *persistence.ProjectRepository) *ProjectService {
	return &ProjectService{
		projectRepo: projectRepo,
	}
}

func (s *ProjectService) GetAll(ctx context.Context, request *dtos.PagedProjectResultRequestDto) (*dtos.PagedResultDto[dtos.ProjectDto], error) {
	filters := make(map[string]interface{})

	if request.GroupID != 0 {
		filters["groupId"] = request.GroupID
	}
	if request.BildirimNo != "" {
		filters["bildirimNo"] = request.BildirimNo
	}
	if request.ProjectCode != "" {
		filters["projectCode"] = request.ProjectCode
	}
	if request.ProjectName != "" {
		filters["projectName"] = request.ProjectName
	}
	if request.ProjectMuellef != "" {
		filters["projectMuellef"] = request.ProjectMuellef
	}
	if len(request.IdList) > 0 {
		filters["idList"] = request.IdList
	}

	projects, totalCount, err := s.projectRepo.GetAllIncludingOcrProjects(
		ctx,
		request.PageNumber,
		request.PageSize,
		filters,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get projects: %w", err)
	}

	projectDtos := make([]dtos.ProjectDto, len(projects))
	for i, project := range projects {
		projectDtos[i] = s.mapToDto(&project)
	}

	return &dtos.PagedResultDto[dtos.ProjectDto]{
		TotalCount: int(totalCount),
		Items:      projectDtos,
	}, nil
}

func (s *ProjectService) GetByID(ctx context.Context, id int) (*dtos.ProjectDto, error) {
	project, err := s.projectRepo.GetByIDIncludingOcrProjects(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}

	dto := s.mapToDto(project)
	return &dto, nil
}

func (s *ProjectService) Create(ctx context.Context, input *dtos.CreateProjectDto) (*dtos.ProjectDto, error) {
	project := &entities.Project{
		ProjectName:          input.ProjectName,
		ProjectCode:          input.ProjectCode,
		ProjectComment:       input.ProjectComment,
		ProjectMuellef:       input.ProjectMuellef,
		Ada:                  input.Ada,
		Parsel:               input.Parsel,
		TalepGucu:            input.TalepGucu,
		KuruluGuc:            input.KuruluGuc,
		BagimsizBS:           input.BagimsizBS,
		BlokS:                input.BlokS,
		YapiYuksekligi:       input.YapiYuksekligi,
		RuhsatGecerlilikDate: input.RuhsatGecerlilikDate,
		YapiSahibi:           input.YapiSahibi,
		Adress:               input.Adress,
		GroupID:              input.GroupID,
		BildirimNo:           input.BildirimNo,
	}

	// TODO: Set CreatorUserID from context/JWT

	if err := s.projectRepo.CreateWithOcrProjects(ctx, project); err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	// Fetch the created project with OCR projects
	createdProject, err := s.projectRepo.GetByIDIncludingOcrProjects(ctx, project.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch created project: %w", err)
	}

	dto := s.mapToDto(createdProject)
	return &dto, nil
}

// Update updates an existing project
func (s *ProjectService) Update(ctx context.Context, id int, input *dtos.UpdateProjectDto) (*dtos.ProjectDto, error) {
	project, err := s.projectRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}

	// Update fields
	project.ProjectName = input.ProjectName
	project.ProjectCode = input.ProjectCode
	project.ProjectComment = input.ProjectComment
	project.ProjectMuellef = input.ProjectMuellef
	project.Ada = input.Ada
	project.Parsel = input.Parsel
	project.TalepGucu = input.TalepGucu
	project.KuruluGuc = input.KuruluGuc
	project.BagimsizBS = input.BagimsizBS
	project.BlokS = input.BlokS
	project.YapiYuksekligi = input.YapiYuksekligi
	project.RuhsatGecerlilikDate = input.RuhsatGecerlilikDate
	project.YapiSahibi = input.YapiSahibi
	project.Adress = input.Adress
	project.GroupID = input.GroupID
	project.BildirimNo = input.BildirimNo

	// TODO: Set LastModifierID from context/JWT

	if err := s.projectRepo.Update(ctx, project); err != nil {
		return nil, fmt.Errorf("failed to update project: %w", err)
	}

	dto := s.mapToDto(project)
	return &dto, nil
}

// Delete deletes a project (soft delete)
func (s *ProjectService) Delete(ctx context.Context, id int, userID int) error {
	if err := s.projectRepo.SoftDelete(ctx, id, userID); err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	return nil
}

// mapToDto converts a project entity to DTO
func (s *ProjectService) mapToDto(project *entities.Project) dtos.ProjectDto {
	dto := dtos.ProjectDto{
		FullAuditedEntityDto: dtos.FullAuditedEntityDto{
			AuditedEntityDto: dtos.AuditedEntityDto{
				EntityDto: dtos.EntityDto{
					ID: project.ID,
				},
				CreatedAt:      project.CreatedAt,
				UpdatedAt:      project.UpdatedAt,
				CreatorUserID:  project.CreatorUserID,
				LastModifierID: project.LastModifierID,
			},
			DeleterUserID: project.DeleterUserID,
			DeletionTime:  project.DeletionTime,
			IsDeleted:     project.IsDeleted,
		},
		ProjectName:          project.ProjectName,
		ProjectCode:          project.ProjectCode,
		ProjectComment:       project.ProjectComment,
		ProjectMuellef:       project.ProjectMuellef,
		Ada:                  project.Ada,
		Parsel:               project.Parsel,
		TalepGucu:            project.TalepGucu,
		KuruluGuc:            project.KuruluGuc,
		BagimsizBS:           project.BagimsizBS,
		BlokS:                project.BlokS,
		YapiYuksekligi:       project.YapiYuksekligi,
		RuhsatGecerlilikDate: project.RuhsatGecerlilikDate,
		YapiSahibi:           project.YapiSahibi,
		Adress:               project.Adress,
		GroupID:              project.GroupID,
		BildirimNo:           project.BildirimNo,
	}

	// Map OCR projects if loaded
	if project.OcrProjects != nil {
		dto.OcrProjects = make([]dtos.OcrProjectDto, len(project.OcrProjects))
		for i, ocrProj := range project.OcrProjects {
			dto.OcrProjects[i] = dtos.OcrProjectDto{
				FullAuditedEntityDto: dtos.FullAuditedEntityDto{
					AuditedEntityDto: dtos.AuditedEntityDto{
						EntityDto: dtos.EntityDto{
							ID: ocrProj.ID,
						},
						CreatedAt:      ocrProj.CreatedAt,
						UpdatedAt:      ocrProj.UpdatedAt,
						CreatorUserID:  ocrProj.CreatorUserID,
						LastModifierID: ocrProj.LastModifierID,
					},
					DeleterUserID: ocrProj.DeleterUserID,
					DeletionTime:  ocrProj.DeletionTime,
					IsDeleted:     ocrProj.IsDeleted,
				},
				ProjectName:          ocrProj.ProjectName,
				ProjectCode:          ocrProj.ProjectCode,
				ProjectComment:       ocrProj.ProjectComment,
				ProjectMuellef:       ocrProj.ProjectMuellef,
				Ada:                  ocrProj.Ada,
				Parsel:               ocrProj.Parsel,
				TalepGucu:            ocrProj.TalepGucu,
				KuruluGuc:            ocrProj.KuruluGuc,
				BagimsizBS:           ocrProj.BagimsizBS,
				BlokS:                ocrProj.BlokS,
				YapiYuksekligi:       ocrProj.YapiYuksekligi,
				RuhsatGecerlilikDate: ocrProj.RuhsatGecerlilikDate,
				YapiSahibi:           ocrProj.YapiSahibi,
				Adress:               ocrProj.Adress,
				Type:                 int(ocrProj.Type),
				TypeName:             ocrProj.Type.String(),
				ProjectID:            ocrProj.ProjectID,
				PdfPath:              ocrProj.PdfPath,
			}
		}
	}

	return dto
}
