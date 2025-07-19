package rag

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
)

// RAGService provides RAG functionality
type RAGService struct {
	vectorStore    VectorStore
	openaiClient   *openai.Client
	collectionName string
}

// QueryResult represents the result of a RAG query
type QueryResult struct {
	Answer     string     `json:"answer"`
	Sources    []Document `json:"sources"`
	Confidence float64    `json:"confidence"`
}

// NewRAGService creates a new RAG service instance
func NewRAGService(vectorStore VectorStore, openaiClient *openai.Client, collectionName string) *RAGService {
	return &RAGService{
		vectorStore:    vectorStore,
		openaiClient:   openaiClient,
		collectionName: collectionName,
	}
}

// Initialize sets up the RAG service
func (r *RAGService) Initialize(ctx context.Context) error {
	// Create collection if it doesn't exist
	if err := r.vectorStore.CreateCollection(ctx, r.collectionName); err != nil {
		log.Printf("Warning: Failed to create collection (might already exist): %v", err)
		// Don't return error as collection might already exist
	}
	
	log.Printf("RAG service initialized with collection: %s", r.collectionName)
	return nil
}

// Query performs a RAG query
func (r *RAGService) Query(ctx context.Context, question string, maxResults int) (*QueryResult, error) {
	if maxResults <= 0 {
		maxResults = 5
	}

	// Step 1: Search for relevant documents
	documents, err := r.vectorStore.Search(ctx, question, maxResults)
	if err != nil {
		return nil, fmt.Errorf("failed to search vector store: %w", err)
	}

	if len(documents) == 0 {
		return &QueryResult{
			Answer:     "申し訳ございませんが、関連する情報が見つかりませんでした。",
			Sources:    []Document{},
			Confidence: 0.0,
		}, nil
	}

	// Step 2: Build context from documents
	context := r.buildContext(documents)

	// Step 3: Generate answer using LLM
	answer, err := r.generateAnswer(ctx, question, context)
	if err != nil {
		return nil, fmt.Errorf("failed to generate answer: %w", err)
	}

	// Step 4: Calculate confidence (simple heuristic)
	confidence := r.calculateConfidence(documents, len(context))

	return &QueryResult{
		Answer:     answer,
		Sources:    documents,
		Confidence: confidence,
	}, nil
}

// buildContext creates a context string from relevant documents
func (r *RAGService) buildContext(documents []Document) string {
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

// generateAnswer uses OpenAI to generate an answer based on context
func (r *RAGService) generateAnswer(ctx context.Context, question, context string) (string, error) {
	// For now, we'll return a mock response
	// In a production environment, you'd want to use the actual OpenAI API
	log.Printf("Generating answer for question: %s", question)
	
	// Return mock response based on question
	if strings.Contains(strings.ToLower(question), "技術スタック") {
		return `このプロジェクトの技術スタックについて説明いたします：

## フロントエンド
- **Next.js 14**: React フレームワーク
- **TypeScript**: 型安全性の確保
- **Tailwind CSS**: モダンなスタイリング
- **React Query**: サーバーステート管理

## バックエンド
- **Go 1.21+**: 高性能バックエンド
- **Gin**: Webフレームワーク
- **GORM**: ORM
- **PostgreSQL**: メインデータベース
- **Redis**: キャッシュ・セッション管理

## 開発環境
- **Docker Compose**: 統一開発環境
- **ChromaDB**: ベクトルデータベース（RAG機能用）
- **OpenAI API**: AI機能統合

この技術スタックは、スケーラブルで保守性の高いWebアプリケーション開発に最適化されています。`, nil
	}
	
	return `プロジェクトの知識ベースを参考に、以下の情報をお答えします：

このプロジェクトは、Go + Next.js + PostgreSQL + Redis による現代的なWebアプリケーション開発のためのテンプレートプロジェクトです。

## 主要な特徴
- チーム開発対応の統一環境
- クリーンアーキテクチャパターンの実装
- AI駆動開発支援機能
- 包括的なテスト戦略

詳細な情報については、プロジェクトのドキュメントをご参照ください。`, nil
}

// calculateConfidence calculates a simple confidence score
func (r *RAGService) calculateConfidence(documents []Document, contextLength int) float64 {
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

// AddDocuments adds documents to the knowledge base
func (r *RAGService) AddDocuments(ctx context.Context, documents []Document) error {
	if len(documents) == 0 {
		return nil
	}

	// Generate embeddings for documents without embeddings
	documentsWithEmbeddings := make([]Document, len(documents))
	for i, doc := range documents {
		documentsWithEmbeddings[i] = doc
		if len(doc.Embedding) == 0 {
			embedding, err := r.generateEmbedding(ctx, doc.Content)
			if err != nil {
				return fmt.Errorf("failed to generate embedding for document %s: %w", doc.ID, err)
			}
			documentsWithEmbeddings[i].Embedding = embedding
		}
	}

	return r.vectorStore.AddDocuments(ctx, documentsWithEmbeddings)
}

// generateEmbedding generates an embedding for text using OpenAI
func (r *RAGService) generateEmbedding(ctx context.Context, text string) ([]float64, error) {
	resp, err := r.openaiClient.CreateEmbeddings(
		ctx,
		openai.EmbeddingRequest{
			Input: []string{text},
			Model: "text-embedding-ada-002",
		},
	)
	
	if err != nil {
		return nil, fmt.Errorf("failed to create embedding: %w", err)
	}
	
	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no embedding data received")
	}
	
	// Convert []float32 to []float64
	embedding := make([]float64, len(resp.Data[0].Embedding))
	for i, v := range resp.Data[0].Embedding {
		embedding[i] = float64(v)
	}
	
	return embedding, nil
}

// GetCollectionInfo returns information about the current collection
func (r *RAGService) GetCollectionInfo(ctx context.Context) (map[string]interface{}, error) {
	// This is a placeholder - in a real implementation, you'd query ChromaDB for collection info
	return map[string]interface{}{
		"collection_name": r.collectionName,
		"service_type":    "RAG Service",
		"vector_store":    "ChromaDB",
		"llm_provider":    "OpenAI",
		"created_at":      time.Now().Format(time.RFC3339),
	}, nil
}
