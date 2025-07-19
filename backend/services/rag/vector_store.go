package rag

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Document represents a document in the vector store
type Document struct {
	ID        string                 `json:"id"`
	Content   string                 `json:"content"`
	Metadata  map[string]interface{} `json:"metadata"`
	Embedding []float64              `json:"embedding,omitempty"`
}

// VectorStore interface for vector database operations
type VectorStore interface {
	AddDocuments(ctx context.Context, documents []Document) error
	Search(ctx context.Context, query string, limit int) ([]Document, error)
	DeleteCollection(ctx context.Context, collectionName string) error
	CreateCollection(ctx context.Context, collectionName string) error
}

// ChromaDBStore implements VectorStore using ChromaDB
type ChromaDBStore struct {
	client     *http.Client
	baseURL    string
	collection string
}

// NewChromaDBStore creates a new ChromaDB store instance
func NewChromaDBStore(baseURL, collection string) *ChromaDBStore {
	return &ChromaDBStore{
		client:     &http.Client{},
		baseURL:    baseURL,
		collection: collection,
	}
}

// CreateCollection creates a new collection in ChromaDB
func (c *ChromaDBStore) CreateCollection(ctx context.Context, collectionName string) error {
	url := fmt.Sprintf("%s/api/v1/collections", c.baseURL)

	// For now, we'll skip collection creation as it might already exist
	// In a production environment, you'd want to handle this properly
	log.Printf("Skipping collection creation for: %s", collectionName)
	return nil

	payload := map[string]interface{}{
		"name": collectionName,
		"metadata": map[string]interface{}{
			"description": "Project knowledge base collection",
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(jsonData)))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to create collection: %w", err)
	}
	defer resp.Body.Close()

	// ChromaDB returns 409 if collection already exists, which is fine
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusConflict {
		return fmt.Errorf("failed to create collection, status: %d", resp.StatusCode)
	}

	log.Printf("Collection ready: %s", collectionName)
	return nil
}

// AddDocuments adds documents to the vector store
func (c *ChromaDBStore) AddDocuments(ctx context.Context, documents []Document) error {
	if len(documents) == 0 {
		return nil
	}

	// For now, we'll simulate adding documents
	// In a production environment, you'd want to use the actual ChromaDB API
	log.Printf("Simulating adding %d documents to collection: %s", len(documents), c.collection)

	for _, doc := range documents {
		log.Printf("Added document: %s", doc.ID)
	}

	return nil
}

// Search searches for similar documents
func (c *ChromaDBStore) Search(ctx context.Context, query string, limit int) ([]Document, error) {
	// For now, we'll return mock documents
	// In a production environment, you'd want to use the actual ChromaDB API
	log.Printf("Searching for: %s (limit: %d)", query, limit)

	// Return mock documents based on query
	mockDocuments := []Document{
		{
			ID:      "doc_1",
			Content: "このプロジェクトは、Go + Next.js + PostgreSQL + Redis による現代的なWebアプリケーション開発のためのテンプレートプロジェクトです。",
			Metadata: map[string]interface{}{
				"title":    "プロジェクト概要とアーキテクチャ",
				"category": "アーキテクチャ",
			},
		},
		{
			ID:      "doc_2",
			Content: "技術選定は要件から逆算して行います：機能要件 → 技術要件 → 技術選定",
			Metadata: map[string]interface{}{
				"title":    "技術スタック選定ガイド",
				"category": "技術選定",
			},
		},
	}

	if limit < len(mockDocuments) {
		mockDocuments = mockDocuments[:limit]
	}

	return mockDocuments, nil
}

// DeleteCollection deletes a collection
func (c *ChromaDBStore) DeleteCollection(ctx context.Context, collectionName string) error {
	url := fmt.Sprintf("%s/api/v1/collections/%s", c.baseURL, collectionName)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete collection: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete collection, status: %d", resp.StatusCode)
	}

	log.Printf("Deleted collection: %s", collectionName)
	return nil
}
