package controllers

import (
	"net/http"
	"reverse-engineering-backend/domain/entities"
	"reverse-engineering-backend/usecases/rag"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RAGController handles RAG-related HTTP requests
type RAGController struct {
	queryUseCase    *rag.RAGQueryUseCase
	indexingUseCase *rag.RAGIndexingUseCase
}

// NewRAGController creates a new RAG controller
func NewRAGController(queryUseCase *rag.RAGQueryUseCase, indexingUseCase *rag.RAGIndexingUseCase) *RAGController {
	return &RAGController{
		queryUseCase:    queryUseCase,
		indexingUseCase: indexingUseCase,
	}
}

// QueryRequest represents a RAG query request
type QueryRequest struct {
	Question   string `json:"question" binding:"required"`
	MaxResults int    `json:"max_results"`
}

// AddDocumentRequest represents a request to add documents
type AddDocumentRequest struct {
	Documents []entities.Document `json:"documents" binding:"required"`
}

// Query handles RAG queries
func (rc *RAGController) Query(c *gin.Context) {
	var req QueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format: " + err.Error(),
		})
		return
	}

	if req.MaxResults <= 0 {
		req.MaxResults = 5
	}

	result, err := rc.queryUseCase.Execute(c.Request.Context(), req.Question, req.MaxResults)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process query: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}

// AddDocuments handles adding documents to the knowledge base
func (rc *RAGController) AddDocuments(c *gin.Context) {
	var req AddDocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format: " + err.Error(),
		})
		return
	}

	if err := rc.indexingUseCase.Execute(c.Request.Context(), req.Documents); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add documents: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Documents added successfully",
		"count":   len(req.Documents),
	})
}

// Search handles direct vector search
func (rc *RAGController) Search(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter 'q' is required",
		})
		return
	}

	limitStr := c.DefaultQuery("limit", "5")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 5
	}

	// Use the query use case for search
	result, err := rc.queryUseCase.Execute(c.Request.Context(), query, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to search: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"query":   query,
			"sources": result.Sources,
			"count":   len(result.Sources),
		},
	})
}

// HealthCheck checks if the RAG service is healthy
func (rc *RAGController) HealthCheck(c *gin.Context) {
	// Simple health check - try to execute a test query
	_, err := rc.queryUseCase.Execute(c.Request.Context(), "test", 1)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "RAG service is operational",
	})
}
