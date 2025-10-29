package persistence

import (
	"context"
	"fmt"

	"hatika-go/internal/domain/entities"

	"gorm.io/gorm"
)

// ProjectRepository implements project-specific repository operations
type ProjectRepository struct {
	*BaseRepository[entities.Project, int]
}

// NewProjectRepository creates a new project repository
func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		BaseRepository: NewBaseRepository[entities.Project, int](db),
	}
}

// GetAllIncludingOcrProjects retrieves projects with OCR projects and pagination
func (r *ProjectRepository) GetAllIncludingOcrProjects(
	ctx context.Context,
	pageNumber, pageSize int,
	filters map[string]interface{},
) ([]entities.Project, int64, error) {
	query := r.GetDB().WithContext(ctx).Preload("OcrProjects")

	// Apply filters
	if groupID, ok := filters["groupId"].(int); ok && groupID != 0 {
		query = query.Where("group_id = ?", groupID)
	}

	if bildirimNo, ok := filters["bildirimNo"].(string); ok && bildirimNo != "" {
		query = query.Where("bildirim_no LIKE ?", "%"+bildirimNo+"%")
	}

	if projectCode, ok := filters["projectCode"].(string); ok && projectCode != "" {
		query = query.Where("project_code LIKE ?", "%"+projectCode+"%")
	}

	if projectName, ok := filters["projectName"].(string); ok && projectName != "" {
		query = query.Where("project_name LIKE ?", "%"+projectName+"%")
	}

	if projectMuellef, ok := filters["projectMuellef"].(string); ok && projectMuellef != "" {
		query = query.Where("project_muellef LIKE ?", "%"+projectMuellef+"%")
	}

	if idList, ok := filters["idList"].([]int); ok && len(idList) > 0 {
		query = query.Where("id IN ?", idList)
	}

	var totalCount int64
	if err := query.Model(&entities.Project{}).Count(&totalCount).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count projects: %w", err)
	}
	//paged
	var projects []entities.Project
	offset := (pageNumber - 1) * pageSize
	if err := query.
		Offset(offset).
		Limit(pageSize).
		Order("id ASC").
		Find(&projects).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch projects: %w", err)
	}

	return projects, totalCount, nil
}

func (r *ProjectRepository) GetByIDIncludingOcrProjects(ctx context.Context, id int) (*entities.Project, error) {
	var project entities.Project
	result := r.GetDB().WithContext(ctx).
		Preload("OcrProjects").
		First(&project, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("project with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to fetch project: %w", result.Error)
	}

	return &project, nil
}

func (r *ProjectRepository) CreateWithOcrProjects(ctx context.Context, project *entities.Project) error {
	return r.GetDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(project).Error; err != nil {
			return fmt.Errorf("failed to create project: %w", err)
		}

		ocrProjects := make([]entities.OcrProject, 4)
		for i := 0; i < 4; i++ {
			ocrProjects[i] = entities.OcrProject{
				ProjectID:   project.ID,
				ProjectName: project.ProjectName,
				ProjectCode: fmt.Sprintf("%s-OCR-%d", project.ProjectCode, i+1),
				Type:        entities.OcrProjectType(i),
			}
		}

		if err := tx.Create(&ocrProjects).Error; err != nil {
			return fmt.Errorf("failed to create OCR projects: %w", err)
		}

		return nil
	})
}

func (r *ProjectRepository) SoftDelete(ctx context.Context, id int, userID int) error {
	return r.GetDB().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var project entities.Project
		if err := tx.First(&project, id).Error; err != nil {
			return fmt.Errorf("project not found: %w", err)
		}

		project.SoftDelete(userID)

		if err := tx.Save(&project).Error; err != nil {
			return fmt.Errorf("failed to soft delete project: %w", err)
		}

		return nil
	})
}
