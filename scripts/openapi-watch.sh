#!/bin/bash

# OpenAPIファイル監視スクリプト
# 使用方法: ./scripts/openapi-watch.sh

set -e

echo "👀 OpenAPIファイルの変更を監視しています..."
echo "📁 監視対象: openapi.yaml"
echo "🔄 変更を検出すると自動的にコードを生成します"
echo "⏹️  停止するには Ctrl+C を押してください"
echo ""

# fswatchがインストールされているかチェック
if ! command -v fswatch &> /dev/null; then
    echo "📦 fswatchをインストールしています..."
    if [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        brew install fswatch
    elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # Linux
        sudo apt-get update && sudo apt-get install -y fswatch
    else
        echo "❌ fswatchのインストールが必要です"
        echo "macOS: brew install fswatch"
        echo "Linux: sudo apt-get install fswatch"
        exit 1
    fi
fi

# 初回生成
echo "🚀 初回生成を実行..."
make generate-openapi

echo ""
echo "✅ 監視を開始しました"

# ファイル変更を監視
fswatch -o openapi.yaml | while read f; do
    echo ""
    echo "🔄 変更を検出しました: $(date)"
    make generate-openapi
    echo "✅ 生成完了: $(date)"
    echo ""
done 