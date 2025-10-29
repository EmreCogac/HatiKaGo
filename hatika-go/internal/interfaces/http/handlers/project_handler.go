package handlers

import (
	"net/http"
	"strconv"

	"hatika-go/internal/application/dtos"
	"hatika-go/internal/application/services"
	"hatika-go/pkg/utils"

	"github.com/gin-gonic/gin"
)

// ProjectHandler handles HTTP requests for projects
type ProjectHandler struct {
	projectService *services.ProjectService
}

// NewProjectHandler creates a new project handler
func NewProjectHandler(projectService *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

// GetAll godoc
// @Summary Get all projects
// @Description Get all projects with pagination and filters
// @Tags projects
// @Accept json
// @Produce json
// @Param pageNumber query int true "Page number" minimum(1)
// @Param pageSize query int true "Page size" minimum(1) maximum(100)
// @Param groupId query int false "Group ID filter"
// @Param bildirimNo query string false "Bildirim No filter"
// @Param projectCode query string false "Project Code filter"
// @Param projectName query string false "Project Name filter"
// @Param projectMuellef query string false "Project Muellef filter"
// @Security BearerAuth
// @Success 200 {object} object "Paged result with projects"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /projects [get]
func (h *ProjectHandler) GetAll(c *gin.Context) {
	var request dtos.PagedProjectResultRequestDto

	if err := c.ShouldBindQuery(&request); err != nil {
		utils.RespondWithValidationError(c, err.Error())
		return
	}

	result, err := h.projectService.GetAll(c.Request.Context(), &request)
	if err != nil {
		utils.RespondInternalError(c, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, result, "")
}

// GetByID godoc
// @Summary Get project by ID
// @Description Get a single project by its ID including OCR projects
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Security BearerAuth
// @Success 200 {object} dtos.ProjectDto
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /projects/{id} [get]
func (h *ProjectHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid project ID", nil)
		return
	}

	result, err := h.projectService.GetByID(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "project not found" {
			utils.RespondNotFound(c, "Project not found")
			return
		}
		utils.RespondInternalError(c, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, result, "")
}

// Create godoc
// @Summary Create a new project
// @Description Create a new project with 4 OCR projects automatically
// @Tags projects
// @Accept json
// @Produce json
// @Param project body dtos.CreateProjectDto true "Project data"
// @Security BearerAuth
// @Success 201 {object} dtos.ProjectDto
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /projects [post]
func (h *ProjectHandler) Create(c *gin.Context) {
	var input dtos.CreateProjectDto

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondWithValidationError(c, err.Error())
		return
	}

	result, err := h.projectService.Create(c.Request.Context(), &input)
	if err != nil {
		utils.RespondInternalError(c, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, result, "Project created successfully")
}

// Update godoc
// @Summary Update a project
// @Description Update an existing project
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param project body dtos.UpdateProjectDto true "Project data"
// @Security BearerAuth
// @Success 200 {object} dtos.ProjectDto
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /projects/{id} [put]
func (h *ProjectHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid project ID", nil)
		return
	}

	var input dtos.UpdateProjectDto
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondWithValidationError(c, err.Error())
		return
	}

	result, err := h.projectService.Update(c.Request.Context(), id, &input)
	if err != nil {
		if err.Error() == "project not found" {
			utils.RespondNotFound(c, "Project not found")
			return
		}
		utils.RespondInternalError(c, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, result, "Project updated successfully")
}

// Delete godoc
// @Summary Delete a project
// @Description Soft delete a project
// @Tags projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Security BearerAuth
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /projects/{id} [delete]
func (h *ProjectHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid project ID", nil)
		return
	}

	// TODO: Get user ID from JWT context
	userID := 1 // Placeholder

	if err := h.projectService.Delete(c.Request.Context(), id, userID); err != nil {
		if err.Error() == "project not found" {
			utils.RespondNotFound(c, "Project not found")
			return
		}
		utils.RespondInternalError(c, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, nil, "Project deleted successfully")
}
