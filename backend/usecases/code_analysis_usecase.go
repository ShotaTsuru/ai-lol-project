package usecases

import (
	"context"
	"reverse-engineering-backend/domain/entities"
	"reverse-engineering-backend/domain/services"
)

// CodeAnalysisUseCase handles code analysis operations
type CodeAnalysisUseCase struct {
	llmService services.LLMService
}

// NewCodeAnalysisUseCase creates a new code analysis use case
func NewCodeAnalysisUseCase(llmService services.LLMService) *CodeAnalysisUseCase {
	return &CodeAnalysisUseCase{
		llmService: llmService,
	}
}

// Execute analyzes code and returns analysis result
func (uc *CodeAnalysisUseCase) Execute(ctx context.Context, code, language string) (*entities.AnalysisResult, error) {
	return uc.llmService.AnalyzeCode(ctx, code, language)
}
