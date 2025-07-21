package requests

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

// FileUploadRequest ファイルアップロードリクエスト
type FileUploadRequest struct {
	ProjectID uint   `json:"project_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Language  string `json:"language"`
}

// AnalysisRequest 解析開始リクエスト
type AnalysisRequest struct {
	ProjectID uint     `json:"project_id" binding:"required"`
	Types     []string `json:"types" binding:"required"`
}

// RAGQueryRequest RAG検索リクエスト
type RAGQueryRequest struct {
	Question   string `json:"question" binding:"required"`
	MaxResults int    `json:"max_results"`
}
