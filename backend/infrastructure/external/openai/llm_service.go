package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reverse-engineering-backend/domain/entities"
	"reverse-engineering-backend/domain/services"

	"github.com/sashabaranov/go-openai"
)

// OpenAIService implements LLMService using OpenAI API
type OpenAIService struct {
	client *openai.Client
}

// NewOpenAIService creates a new OpenAI service instance
func NewOpenAIService() services.LLMService {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return &OpenAIService{client: nil}
	}

	return &OpenAIService{
		client: openai.NewClient(apiKey),
	}
}

// GenerateAnswer generates an answer using OpenAI
func (o *OpenAIService) GenerateAnswer(ctx context.Context, question, context string) (string, error) {
	if o.client == nil {
		return o.mockAnswer(question), nil
	}

	prompt := fmt.Sprintf(`
以下のプロジェクト知識ベースを参考にしてください：

%s

質問：%s

プロジェクトの知識ベースに基づいて回答してください。
`, context, question)

	resp, err := o.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: 2000,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

// GenerateEmbedding generates embeddings using OpenAI
func (o *OpenAIService) GenerateEmbedding(ctx context.Context, text string) ([]float64, error) {
	if o.client == nil {
		// Return mock embedding
		return make([]float64, 1536), nil
	}

	resp, err := o.client.CreateEmbeddings(
		ctx,
		openai.EmbeddingRequest{
			Input: text,
			Model: "text-embedding-ada-002",
		},
	)

	if err != nil {
		return nil, err
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

// AnalyzeCode analyzes code using OpenAI
func (o *OpenAIService) AnalyzeCode(ctx context.Context, code, language string) (*entities.AnalysisResult, error) {
	if o.client == nil {
		return o.mockCodeAnalysis(code, language), nil
	}

	prompt := fmt.Sprintf(`
以下の%sコードを解析して、以下の情報をJSON形式で提供してください：

1. コードの概要と目的
2. 主要な関数・メソッドの一覧
3. 使用されているデザインパターン
4. 潜在的な問題点や改善提案
5. 依存関係の分析

コード：
%s

JSON形式で回答してください。
`, language, code)

	resp, err := o.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: 2000,
		},
	)

	if err != nil {
		return nil, err
	}

	// Parse JSON response
	var result entities.AnalysisResult
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GenerateDocumentation generates documentation using OpenAI
func (o *OpenAIService) GenerateDocumentation(ctx context.Context, code, language string) (string, error) {
	if o.client == nil {
		return o.mockDocumentation(code, language), nil
	}

	prompt := fmt.Sprintf(`
以下の%sコードの技術文書を作成してください。以下の要素を含めてください：

1. API仕様（関数・メソッドの説明）
2. アーキテクチャ概要
3. 使用方法の例
4. 設定方法
5. トラブルシューティング

コード：
%s

Markdown形式で回答してください。
`, language, code)

	resp, err := o.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: 3000,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

// DetectPatterns detects patterns using OpenAI
func (o *OpenAIService) DetectPatterns(ctx context.Context, code, language string) (*entities.AnalysisResult, error) {
	if o.client == nil {
		return o.mockPatternDetection(code, language), nil
	}

	prompt := fmt.Sprintf(`
以下の%sコードを分析して、使用されているデザインパターンやアンチパターンを特定してください：

1. デザインパターン（Singleton, Factory, Observer, etc.）
2. アンチパターン（God Object, Spaghetti Code, etc.）
3. コード品質の評価
4. リファクタリング提案

コード：
%s

JSON形式で回答してください。
`, language, code)

	resp, err := o.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: 2000,
		},
	)

	if err != nil {
		return nil, err
	}

	var result entities.AnalysisResult
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// AnalyzeDependencies analyzes dependencies using OpenAI
func (o *OpenAIService) AnalyzeDependencies(ctx context.Context, files []entities.FileInfo) (*entities.AnalysisResult, error) {
	if o.client == nil {
		return o.mockDependencyAnalysis(files), nil
	}

	fileList := "ファイル一覧:\n"
	for _, file := range files {
		fileList += fmt.Sprintf("- %s (%s)\n", file.Name, file.Language)
	}

	prompt := fmt.Sprintf(`
以下のファイル群の依存関係を分析して、プロジェクト構造を可視化してください：

%s

以下の情報をJSON形式で提供してください：
1. ファイル間の依存関係マップ
2. モジュール構造の分析
3. 循環依存の検出
4. アーキテクチャの改善提案

JSON形式で回答してください。
`, fileList)

	resp, err := o.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: 2000,
		},
	)

	if err != nil {
		return nil, err
	}

	var result entities.AnalysisResult
	if err := json.Unmarshal([]byte(resp.Choices[0].Message.Content), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Mock methods for development
func (o *OpenAIService) mockAnswer(question string) string {
	return "プロジェクトの知識ベースを参考にした回答です。"
}

func (o *OpenAIService) mockCodeAnalysis(code, language string) *entities.AnalysisResult {
	return &entities.AnalysisResult{
		Summary:         "コード解析結果",
		Functions:       []string{"main", "handler"},
		Patterns:        []string{"MVC", "Repository"},
		Issues:          []string{"改善の余地あり"},
		Dependencies:    map[string]interface{}{"framework": "gin"},
		Recommendations: []string{"テストの追加", "エラーハンドリングの改善"},
	}
}

func (o *OpenAIService) mockDocumentation(code, language string) string {
	return "# ドキュメント\n\nこのコードのドキュメントです。"
}

func (o *OpenAIService) mockPatternDetection(code, language string) *entities.AnalysisResult {
	return &entities.AnalysisResult{
		Summary:         "パターン検出結果",
		Functions:       []string{},
		Patterns:        []string{"Repository Pattern", "Dependency Injection"},
		Issues:          []string{},
		Dependencies:    map[string]interface{}{},
		Recommendations: []string{"パターンの適用を継続"},
	}
}

func (o *OpenAIService) mockDependencyAnalysis(files []entities.FileInfo) *entities.AnalysisResult {
	return &entities.AnalysisResult{
		Summary:         "依存関係分析結果",
		Functions:       []string{},
		Patterns:        []string{},
		Issues:          []string{},
		Dependencies:    map[string]interface{}{"files": len(files)},
		Recommendations: []string{"依存関係の整理"},
	}
}
