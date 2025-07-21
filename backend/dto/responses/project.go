package responses

import "time"

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

// RAGQueryResponse RAG検索レスポンス
type RAGQueryResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Answer     string     `json:"answer"`
		Sources    []Document `json:"sources"`
		Confidence float64    `json:"confidence"`
	} `json:"data"`
}

// Document ドキュメント
type Document struct {
	ID       string                 `json:"id"`
	Content  string                 `json:"content"`
	Metadata map[string]interface{} `json:"metadata"`
}

// ErrorResponse エラーレスポンス
type ErrorResponse struct {
	Error string `json:"error"`
}
