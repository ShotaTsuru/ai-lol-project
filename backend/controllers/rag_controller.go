package controllers

import (
	"net/http"
	"strconv"

	"reverse-engineering-backend/services/rag"

	"github.com/gin-gonic/gin"
)

// RAGController handles RAG-related HTTP requests
type RAGController struct {
	ragService *rag.RAGService
}

// NewRAGController creates a new RAG controller
func NewRAGController(ragService *rag.RAGService) *RAGController {
	return &RAGController{
		ragService: ragService,
	}
}

// QueryRequest represents a RAG query request
type QueryRequest struct {
	Question   string `json:"question" binding:"required"`
	MaxResults int    `json:"max_results"`
}

// AddDocumentRequest represents a request to add documents
type AddDocumentRequest struct {
	Documents []rag.Document `json:"documents" binding:"required"`
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

	result, err := rc.ragService.Query(c.Request.Context(), req.Question, req.MaxResults)
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

	if err := rc.ragService.AddDocuments(c.Request.Context(), req.Documents); err != nil {
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

// GetCollectionInfo returns information about the current collection
func (rc *RAGController) GetCollectionInfo(c *gin.Context) {
	info, err := rc.ragService.GetCollectionInfo(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get collection info: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    info,
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

	// Get vector store from RAG service (this would need to be exposed)
	// For now, we'll use the RAG service's query method
	result, err := rc.ragService.Query(c.Request.Context(), query, limit)
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
	info, err := rc.ragService.GetCollectionInfo(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "unhealthy",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"info":   info,
	})
}
