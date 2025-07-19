package usecases

import (
	"context"
	"reverse-engineering-backend/domain/services"
)

// DocumentationUseCase handles documentation generation operations
type DocumentationUseCase struct {
	llmService services.LLMService
}

// NewDocumentationUseCase creates a new documentation use case
func NewDocumentationUseCase(llmService services.LLMService) *DocumentationUseCase {
	return &DocumentationUseCase{
		llmService: llmService,
	}
}

// Execute generates documentation for code
func (uc *DocumentationUseCase) Execute(ctx context.Context, code, language string) (string, error) {
	return uc.llmService.GenerateDocumentation(ctx, code, language)
}
