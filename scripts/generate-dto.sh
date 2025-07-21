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
mkdir -p backend/dto/generated

# DTO型の生成
echo "📝 DTO型を生成しています..."
oapi-codegen \
    -package generated \
    -generate types \
    openapi.yaml > backend/dto/generated/types.go

# バリデーション関数の生成
echo "🔍 バリデーション関数を生成しています..."
oapi-codegen \
    -package generated \
    -generate spec \
    openapi.yaml > backend/dto/generated/spec.go

# サーバーコードの生成（オプション）
echo "🖥️ サーバーコードを生成しています..."
oapi-codegen \
    -package generated \
    -generate server \
    openapi.yaml > backend/dto/generated/server.go

echo "✅ DTO生成が完了しました！"
echo "📁 生成されたファイル:"
echo "   - backend/dto/generated/types.go"
echo "   - backend/dto/generated/spec.go"
echo "   - backend/dto/generated/server.go"

# 生成されたコードの確認
echo "🔍 生成されたコードを確認しています..."
go fmt backend/dto/generated/*.go
go vet backend/dto/generated/*.go

echo "🎉 すべて完了しました！" 