package repositories

import (
	"context"

	"reverse-engineering-backend/domain/entities"
)

// VectorRepository defines the interface for vector store operations
type VectorRepository interface {
	AddDocuments(ctx context.Context, documents []entities.Document) error
	Search(ctx context.Context, query string, limit int) ([]entities.Document, error)
	DeleteCollection(ctx context.Context, collectionName string) error
	CreateCollection(ctx context.Context, collectionName string) error
}
