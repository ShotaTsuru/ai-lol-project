package dto

import "time"

// CreateProjectRequest プロジェクト作成リクエスト
type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id" binding:"required"`
}

// UpdateProjectRequest プロジェクト更新リクエスト
type UpdateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// ProjectResponse プロジェクトレスポンス
type ProjectResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      uint      `json:"user_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	FileCount   int       `json:"file_count"`
}

// ProjectListResponse プロジェクト一覧レスポンス
type ProjectListResponse struct {
	Projects []ProjectResponse `json:"projects"`
	Total    int               `json:"total"`
	Page     int               `json:"page"`
	Limit    int               `json:"limit"`
}

// FileUploadRequest ファイルアップロードリクエスト
type FileUploadRequest struct {
	ProjectID uint   `json:"project_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Language  string `json:"language"`
}

// FileResponse ファイルレスポンス
type FileResponse struct {
	ID        uint      `json:"id"`
	ProjectID uint      `json:"project_id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Size      int64     `json:"size"`
	MimeType  string    `json:"mime_type"`
	Language  string    `json:"language"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AnalysisRequest 解析開始リクエスト
type AnalysisRequest struct {
	ProjectID uint     `json:"project_id" binding:"required"`
	Types     []string `json:"types" binding:"required"`
}

// AnalysisResponse 解析レスポンス
type AnalysisResponse struct {
	ID        uint      `json:"id"`
	ProjectID uint      `json:"project_id"`
	FileID    *uint     `json:"file_id,omitempty"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	Result    string    `json:"result,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
