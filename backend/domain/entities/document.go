package entities

// Document represents a document in the knowledge base
type Document struct {
	ID        string                 `json:"id"`
	Content   string                 `json:"content"`
	Metadata  map[string]interface{} `json:"metadata"`
	Embedding []float64              `json:"embedding,omitempty"`
}

// QueryResult represents the result of a RAG query
type QueryResult struct {
	Answer     string     `json:"answer"`
	Sources    []Document `json:"sources"`
	Confidence float64    `json:"confidence"`
}

// AnalysisResult represents the result of code analysis
type AnalysisResult struct {
	Summary         string                 `json:"summary"`
	Functions       []string               `json:"functions"`
	Patterns        []string               `json:"patterns"`
	Issues          []string               `json:"issues"`
	Dependencies    map[string]interface{} `json:"dependencies"`
	Recommendations []string               `json:"recommendations"`
}

// FileInfo represents information about a file
type FileInfo struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Content  string `json:"content"`
}
