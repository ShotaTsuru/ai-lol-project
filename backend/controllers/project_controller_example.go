package controllers

import (
	"net/http"
	"reverse-engineering-backend/controllers/dto"
	"reverse-engineering-backend/domain/entities"
	"reverse-engineering-backend/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProjectControllerExample 生成されたDTOを使用するコントローラーの例
type ProjectControllerExample struct {
	createProjectUseCase *usecases.CreateProjectUseCase
	getProjectUseCase    *usecases.GetProjectUseCase
}

// NewProjectControllerExample 新しいプロジェクトコントローラーを作成
func NewProjectControllerExample(
	createProjectUseCase *usecases.CreateProjectUseCase,
	getProjectUseCase *usecases.GetProjectUseCase,
) *ProjectControllerExample {
	return &ProjectControllerExample{
		createProjectUseCase: createProjectUseCase,
		getProjectUseCase:    getProjectUseCase,
	}
}

// CreateProject プロジェクト作成（生成されたDTO使用）
func (pc *ProjectControllerExample) CreateProject(c *gin.Context) {
	var req dto.CreateProjectRequest

	// 生成されたバリデーションを使用
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.BadRequest{
			Error: "Invalid request format: " + err.Error(),
		})
		return
	}

	// 追加のバリデーション（生成された関数を使用）
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, dto.BadRequest{
			Error: "Validation failed: " + err.Error(),
		})
		return
	}

	// DTO → Entity変換
	project := &entities.Project{
		Name:        req.Name,
		Description: req.Description,
		UserID:      req.UserId,
		Status:      "pending",
	}

	// Usecase実行
	if err := pc.createProjectUseCase.Execute(c.Request.Context(), project); err != nil {
		c.JSON(http.StatusInternalServerError, dto.InternalServerError{
			Error: "Failed to create project: " + err.Error(),
		})
		return
	}

	// Entity → DTO変換
	response := dto.ProjectResponse{
		Id:          int(project.ID),
		Name:        project.Name,
		Description: project.Description,
		UserId:      int(project.UserID),
		Status:      project.Status,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
		FileCount:   0, // 新規作成時は0
	}

	c.JSON(http.StatusCreated, response)
}

// GetProject プロジェクト取得（生成されたDTO使用）
func (pc *ProjectControllerExample) GetProject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.BadRequest{
			Error: "Invalid project ID",
		})
		return
	}

	// Usecase実行
	project, err := pc.getProjectUseCase.Execute(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, dto.NotFound{
			Error: "Project not found",
		})
		return
	}

	// Entity → DTO変換
	response := dto.ProjectResponse{
		Id:          int(project.ID),
		Name:        project.Name,
		Description: project.Description,
		UserId:      int(project.UserID),
		Status:      project.Status,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
		FileCount:   len(project.Files), // 実際のファイル数を取得
	}

	c.JSON(http.StatusOK, response)
}

// GetProjects プロジェクト一覧取得（生成されたDTO使用）
func (pc *ProjectControllerExample) GetProjects(c *gin.Context) {
	// クエリパラメータの取得とバリデーション
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	// Usecase実行（実際の実装は省略）
	// projects, total, err := pc.listProjectsUseCase.Execute(c.Request.Context(), page, limit)

	// 仮のデータ
	projects := []dto.ProjectResponse{
		{
			Id:          1,
			Name:        "Sample Project",
			Description: "This is a sample project",
			UserId:      1,
			Status:      "completed",
			CreatedAt:   "2024-01-01T00:00:00Z",
			UpdatedAt:   "2024-01-01T00:00:00Z",
			FileCount:   5,
		},
	}

	response := dto.ProjectListResponse{
		Projects: projects,
		Total:    1,
		Page:     page,
		Limit:    limit,
	}

	c.JSON(http.StatusOK, response)
}
