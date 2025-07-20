#!/bin/bash

# DTO生成スクリプト
# 使用方法: ./scripts/generate-dto.sh

set -e

echo "🚀 DTO生成を開始します..."

# oapi-codegenのインストール確認
if ! command -v oapi-codegen &> /dev/null; then
    echo "📦 oapi-codegenをインストールしています..."
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
fi

# 出力ディレクトリの作成
mkdir -p backend/controllers/dto

# DTO型の生成
echo "📝 DTO型を生成しています..."
oapi-codegen \
    -package dto \
    -generate types \
    openapi.yaml > backend/controllers/dto/generated_types.go

# バリデーション関数の生成
echo "🔍 バリデーション関数を生成しています..."
oapi-codegen \
    -package dto \
    -generate spec \
    openapi.yaml > backend/controllers/dto/generated_spec.go

# サーバーコードの生成（オプション）
echo "🖥️ サーバーコードを生成しています..."
oapi-codegen \
    -package dto \
    -generate server \
    openapi.yaml > backend/controllers/dto/generated_server.go

echo "✅ DTO生成が完了しました！"
echo "📁 生成されたファイル:"
echo "   - backend/controllers/dto/generated_types.go"
echo "   - backend/controllers/dto/generated_spec.go"
echo "   - backend/controllers/dto/generated_server.go"

# 生成されたコードの確認
echo "🔍 生成されたコードを確認しています..."
go fmt backend/controllers/dto/*.go
go vet backend/controllers/dto/*.go

echo "🎉 すべて完了しました！" 