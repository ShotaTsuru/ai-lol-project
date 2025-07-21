package main

import (
	"log"
	"os"
	"reverse-engineering-backend/config"
	"reverse-engineering-backend/controllers"
	"reverse-engineering-backend/infrastructure/external/chromadb"
	"reverse-engineering-backend/infrastructure/external/openai"
	"reverse-engineering-backend/routes"

	"reverse-engineering-backend/usecases/rag"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 環境変数の読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// データベース接続
	db, err := config.InitDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Redis接続
	redis, err := config.InitRedis()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	// インフラストラクチャ層の初期化
	llmService := openai.NewOpenAIService()
	vectorRepo := chromadb.NewChromaDBVectorRepository(
		os.Getenv("CHROMADB_URL"),
		"project_knowledge_base",
	)

	// ユースケース層の初期化
	ragQueryUseCase := rag.NewRAGQueryUseCase(vectorRepo, llmService, "project_knowledge_base")
	ragIndexingUseCase := rag.NewRAGIndexingUseCase(vectorRepo, llmService)

	// RAGサービスの初期化
	if err := ragQueryUseCase.Initialize(gin.Context{}.Request.Context()); err != nil {
		log.Printf("Warning: Failed to initialize RAG service: %v", err)
	}

	// コントローラー層の初期化
	ragController := controllers.NewRAGController(ragQueryUseCase, ragIndexingUseCase)

	// Ginエンジンの初期化
	if os.Getenv("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// CORS設定
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		"http://localhost:3000",
		"http://127.0.0.1:3000",
	}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	corsConfig.AllowCredentials = true

	r.Use(cors.New(corsConfig))

	// ルートの設定
	routes.SetupRoutes(r, db, redis, ragController)

	// サーバー起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
