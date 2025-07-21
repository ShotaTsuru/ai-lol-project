package rag

import (
	"context"
	"fmt"
	"log"
	"reverse-engineering-backend/domain/entities"
	"reverse-engineering-backend/domain/repositories"
	"reverse-engineering-backend/domain/services"
	"strings"
)

// RAGQueryUseCase handles RAG query operations
type RAGQueryUseCase struct {
	vectorRepo     repositories.VectorRepository
	llmService     services.LLMService
	collectionName string
}

// NewRAGQueryUseCase creates a new RAG query use case
func NewRAGQueryUseCase(vectorRepo repositories.VectorRepository, llmService services.LLMService, collectionName string) *RAGQueryUseCase {
	return &RAGQueryUseCase{
		vectorRepo:     vectorRepo,
		llmService:     llmService,
		collectionName: collectionName,
	}
}

// Initialize sets up the RAG query use case
func (uc *RAGQueryUseCase) Initialize(ctx context.Context) error {
	// Create collection if it doesn't exist
	if err := uc.vectorRepo.CreateCollection(ctx, uc.collectionName); err != nil {
		log.Printf("Warning: Failed to create collection (might already exist): %v", err)
		// Don't return error as collection might already exist
	}

	log.Printf("RAG query use case initialized with collection: %s", uc.collectionName)
	return nil
}

// Execute performs a RAG query
func (uc *RAGQueryUseCase) Execute(ctx context.Context, question string, maxResults int) (*entities.QueryResult, error) {
	if maxResults <= 0 {
		maxResults = 5
	}

	// Step 1: Search for relevant documents
	documents, err := uc.vectorRepo.Search(ctx, question, maxResults)
	if err != nil {
		return nil, fmt.Errorf("failed to search vector store: %w", err)
	}

	if len(documents) == 0 {
		return &entities.QueryResult{
			Answer:     "申し訳ございませんが、関連する情報が見つかりませんでした。",
			Sources:    []entities.Document{},
			Confidence: 0.0,
		}, nil
	}

	// Step 2: Build context from documents
	context := uc.buildContext(documents)

	// Step 3: Generate answer using LLM
	answer, err := uc.llmService.GenerateAnswer(ctx, question, context)
	if err != nil {
		return nil, fmt.Errorf("failed to generate answer: %w", err)
	}

	// Step 4: Calculate confidence (simple heuristic)
	confidence := uc.calculateConfidence(documents, len(context))

	return &entities.QueryResult{
		Answer:     answer,
		Sources:    documents,
		Confidence: confidence,
	}, nil
}

// buildContext creates a context string from relevant documents
func (uc *RAGQueryUseCase) buildContext(documents []entities.Document) string {
	var contextBuilder strings.Builder

	contextBuilder.WriteString("以下のプロジェクト知識ベースを参考にしてください：\n\n")

	for i, doc := range documents {
		contextBuilder.WriteString(fmt.Sprintf("--- ドキュメント %d ---\n", i+1))
		contextBuilder.WriteString(doc.Content)
		contextBuilder.WriteString("\n\n")

		// Add metadata if available
		if len(doc.Metadata) > 0 {
			contextBuilder.WriteString("メタデータ: ")
			for key, value := range doc.Metadata {
				contextBuilder.WriteString(fmt.Sprintf("%s=%v, ", key, value))
			}
			contextBuilder.WriteString("\n\n")
		}
	}

	return contextBuilder.String()
}

// calculateConfidence calculates a simple confidence score
func (uc *RAGQueryUseCase) calculateConfidence(documents []entities.Document, contextLength int) float64 {
	// Simple heuristic: more documents and longer context = higher confidence
	baseConfidence := 0.5

	// Bonus for number of relevant documents
	docBonus := float64(len(documents)) * 0.1
	if docBonus > 0.3 {
		docBonus = 0.3
	}

	// Bonus for context length
	contextBonus := float64(contextLength) / 10000.0
	if contextBonus > 0.2 {
		contextBonus = 0.2
	}

	confidence := baseConfidence + docBonus + contextBonus
	if confidence > 1.0 {
		confidence = 1.0
	}

	return confidence
}
