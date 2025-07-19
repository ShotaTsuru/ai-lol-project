package services

import (
	"context"
	"reverse-engineering-backend/domain/entities"
)

// LLMService defines the interface for LLM operations
type LLMService interface {
	GenerateAnswer(ctx context.Context, question, context string) (string, error)
	GenerateEmbedding(ctx context.Context, text string) ([]float64, error)
	AnalyzeCode(ctx context.Context, code, language string) (*entities.AnalysisResult, error)
	GenerateDocumentation(ctx context.Context, code, language string) (string, error)
	DetectPatterns(ctx context.Context, code, language string) (*entities.AnalysisResult, error)
	AnalyzeDependencies(ctx context.Context, files []entities.FileInfo) (*entities.AnalysisResult, error)
}
