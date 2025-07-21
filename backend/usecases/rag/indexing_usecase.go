package rag

import (
	"context"
	"fmt"
	"reverse-engineering-backend/domain/entities"
	"reverse-engineering-backend/domain/repositories"
	"reverse-engineering-backend/domain/services"
)

// RAGIndexingUseCase handles RAG indexing operations
type RAGIndexingUseCase struct {
	vectorRepo repositories.VectorRepository
	llmService services.LLMService
}

// NewRAGIndexingUseCase creates a new RAG indexing use case
func NewRAGIndexingUseCase(vectorRepo repositories.VectorRepository, llmService services.LLMService) *RAGIndexingUseCase {
	return &RAGIndexingUseCase{
		vectorRepo: vectorRepo,
		llmService: llmService,
	}
}

// Execute adds documents to the knowledge base
func (uc *RAGIndexingUseCase) Execute(ctx context.Context, documents []entities.Document) error {
	if len(documents) == 0 {
		return nil
	}

	// Generate embeddings for documents without embeddings
	documentsWithEmbeddings := make([]entities.Document, len(documents))
	for i, doc := range documents {
		documentsWithEmbeddings[i] = doc
		if len(doc.Embedding) == 0 {
			embedding, err := uc.llmService.GenerateEmbedding(ctx, doc.Content)
			if err != nil {
				return fmt.Errorf("failed to generate embedding for document %s: %w", doc.ID, err)
			}
			documentsWithEmbeddings[i].Embedding = embedding
		}
	}

	return uc.vectorRepo.AddDocuments(ctx, documentsWithEmbeddings)
}
