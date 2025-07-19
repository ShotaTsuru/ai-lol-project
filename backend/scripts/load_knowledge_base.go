package main

import (
	"context"
	"fmt"
	"log"

	"reverse-engineering-backend/services/rag"
)

// KnowledgeDocument represents a document in the knowledge base
type KnowledgeDocument struct {
	Title       string            `json:"title"`
	Content     string            `json:"content"`
	Category    string            `json:"category"`
	Tags        []string          `json:"tags"`
	LastUpdated string            `json:"last_updated"`
}

func main() {
	// Initialize vector store
	vectorStore := rag.NewChromaDBStore("http://localhost:8000", "project_knowledge_base")
	
	// Create collection
	ctx := context.Background()
	if err := vectorStore.CreateCollection(ctx, "project_knowledge_base"); err != nil {
		log.Printf("Warning: Failed to create collection: %v", err)
	}
	
	// Load knowledge documents
	documents := loadKnowledgeDocuments()
	
	// Convert to RAG documents
	ragDocuments := make([]rag.Document, len(documents))
	for i, doc := range documents {
		ragDocuments[i] = rag.Document{
			ID: fmt.Sprintf("doc_%d", i+1),
			Content: fmt.Sprintf("タイトル: %s\n\n%s", doc.Title, doc.Content),
			Metadata: map[string]interface{}{
				"title":        doc.Title,
				"category":     doc.Category,
				"tags":         doc.Tags,
				"last_updated": doc.LastUpdated,
			},
		}
	}
	
	// Add documents to vector store
	if err := vectorStore.AddDocuments(ctx, ragDocuments); err != nil {
		log.Fatalf("Failed to add documents: %v", err)
	}
	
	log.Printf("Successfully loaded %d documents into knowledge base", len(documents))
}

func loadKnowledgeDocuments() []KnowledgeDocument {
	return []KnowledgeDocument{
		{
			Title: "プロジェクト概要とアーキテクチャ",
			Content: `このプロジェクトは、Go + Next.js + PostgreSQL + Redis による現代的なWebアプリケーション開発のためのテンプレートプロジェクトです。

## アーキテクチャ

### フロントエンド (React + Next.js)
- Next.js 14ベースのWebアプリケーション
- TypeScript + Tailwind CSS による型安全でモダンなUI開発
- React Query によるサーバーステート管理

### バックエンド (Go)
- Goベースの高性能RESTful APIサーバー
- クリーンアーキテクチャパターンの実装
- ルーティング、コントローラー、サービス層の明確な分離

## 主要機能・特徴
- チーム開発対応: Docker Compose による統一開発環境
- スケーラブル: PostgreSQL + Redis によるデータ層
- セキュリティ: CORS設定、環境変数管理
- 高速開発: ホットリロード対応開発サーバー
- テスト準備: フロントエンド・バックエンド両方のテスト環境
- AI Issue作成システム: GitHub Actionsとテンプレートによる効率的なissue管理`,
			Category: "アーキテクチャ",
			Tags: []string{"アーキテクチャ", "概要", "技術スタック"},
			LastUpdated: "2024-01-01",
		},
		{
			Title: "技術スタック選定ガイド",
			Content: `## 技術選定のフレームワーク

### 1. 要件からの逆算手法
技術選定は要件から逆算して行います：
- 機能要件 → 技術要件 → 技術選定
- 非機能要件 → パフォーマンス要件 → 技術選定
- 運用要件 → 運用技術選定

### 2. 比較分析ツール
- 客観的な比較指標の設定
- 他技術との比較分析
- トレードオフの明確化

### 3. ROI分析手法
- 投資対効果の定量的評価
- 学習コストの考慮
- 長期運用コストの計算

### 4. 意思決定フロー
- 判断プロセスの可視化
- チーム合意形成の方法
- リスク評価と対策`,
			Category: "技術選定",
			Tags: []string{"技術選定", "フレームワーク", "ROI分析"},
			LastUpdated: "2024-01-01",
		},
		{
			Title: "AI駆動開発のベストプラクティス",
			Content: `## AI駆動開発の成熟度レベル

### レベル1: 基本的なコード生成
- 単純なコード生成
- コメント生成
- 基本的なリファクタリング

### レベル2: 設計支援
- アーキテクチャ設計支援
- データベース設計支援
- API設計支援

### レベル3: 品質向上支援
- テストケース生成
- コードレビュー支援
- パフォーマンス分析

### レベル4: プロジェクト管理支援
- 要件定義支援
- タスク分解支援
- 進捗管理支援

## 人間-AI協働パターン
- 設計駆動アプローチ
- TDD協働アプローチ
- 段階的検証プロセス
- 批判的思考の維持`,
			Category: "AI駆動開発",
			Tags: []string{"AI", "開発手法", "ベストプラクティス"},
			LastUpdated: "2024-01-01",
		},
		{
			Title: "クリーンアーキテクチャ実装パターン",
			Content: `## Go クリーンアーキテクチャ実装

### レイヤー構成
controllers/     # インターフェースアダプター - HTTPハンドラー
    ↓ (依存関係逆転)
services/        # アプリケーションビジネスルール - ユースケース
    ↓ (依存関係逆転)
domain/          # エンタープライズビジネスルール - エンティティ
    ↑ (インターフェース実装)
repositories/    # データアクセス抽象化
    ↑ (実装)
infrastructure/  # 外部レイヤー - データベース、外部API

### 実装原則
- 依存関係は内側に向かってのみ流れる
- 内側の層は外側の層を知らない
- インターフェースによる依存関係逆転
- ドメインロジックはドメイン層に集約

### エラーハンドリングパターン
- 統一エラーレスポンス形式
- 適切なHTTPステータスコード
- エラーログの記録
- クライアントへの適切なエラー情報提供`,
			Category: "実装パターン",
			Tags: []string{"クリーンアーキテクチャ", "Go", "設計パターン"},
			LastUpdated: "2024-01-01",
		},
		{
			Title: "テスト戦略と実装",
			Content: `## 包括的テスト戦略

### テストピラミッド
1. **単体テスト** (基盤)
   - 関数・メソッドレベルのテスト
   - 高速実行、高カバレッジ
   - モック・スタブの活用

2. **結合テスト** (中間)
   - コンポーネント間の統合テスト
   - データベース・外部API統合
   - エンドツーエンドの一部機能

3. **E2Eテスト** (頂点)
   - ユーザーシナリオベース
   - 実際のブラウザ・環境でのテスト
   - 重要なユーザーフロー

### テスト実装のベストプラクティス
- AAA パターン (Arrange, Act, Assert)
- テストデータの管理
- テスト環境の分離
- CI/CDパイプライン統合

### 品質保証プロセス
- コードレビュー
- 静的解析
- セキュリティスキャン
- パフォーマンステスト`,
			Category: "テスト",
			Tags: []string{"テスト", "品質保証", "CI/CD"},
			LastUpdated: "2024-01-01",
		},
	}
} 