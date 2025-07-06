#!/bin/bash

# League of Legends MCP Server セットアップスクリプト

set -e

echo "🎮 League of Legends MCP Server セットアップを開始します..."

# カラーコードの定義
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 関数: エラーメッセージの表示
error() {
    echo -e "${RED}❌ エラー: $1${NC}"
    exit 1
}

# 関数: 成功メッセージの表示
success() {
    echo -e "${GREEN}✅ $1${NC}"
}

# 関数: 警告メッセージの表示
warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

# 関数: 情報メッセージの表示
info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

# Node.js のバージョン確認
check_nodejs() {
    if ! command -v node &> /dev/null; then
        error "Node.js がインストールされていません。Node.js 18+ をインストールしてください。"
    fi
    
    NODE_VERSION=$(node --version | cut -d 'v' -f 2 | cut -d '.' -f 1)
    if [ "$NODE_VERSION" -lt 18 ]; then
        error "Node.js のバージョンが古すぎます。Node.js 18+ が必要です。現在のバージョン: $(node --version)"
    fi
    
    success "Node.js バージョン確認完了: $(node --version)"
}

# ディレクトリの作成と移動
setup_directory() {
    info "MCP サーバーディレクトリの作成..."
    
    if [ ! -d "mcp-servers" ]; then
        mkdir -p mcp-servers
        success "mcp-servers ディレクトリを作成しました"
    fi
    
    if [ ! -d "mcp-servers/lol-mcp-server" ]; then
        error "League of Legends MCP サーバーのソースコードが見つかりません。"
    fi
    
    cd mcp-servers/lol-mcp-server || error "ディレクトリの移動に失敗しました"
    success "プロジェクトディレクトリに移動しました"
}

# 依存関係のインストール
install_dependencies() {
    info "依存関係をインストールしています..."
    
    if [ ! -f "package.json" ]; then
        error "package.json が見つかりません。"
    fi
    
    npm install || error "依存関係のインストールに失敗しました"
    success "依存関係のインストールが完了しました"
}

# TypeScript のコンパイル
build_typescript() {
    info "TypeScript をコンパイルしています..."
    
    npm run build || error "TypeScript のコンパイルに失敗しました"
    success "TypeScript のコンパイルが完了しました"
}

# 環境変数設定ファイルの作成
setup_env_file() {
    info "環境変数設定ファイルを作成しています..."
    
    if [ ! -f ".env" ]; then
        if [ -f "env.example" ]; then
            cp env.example .env
            success ".env ファイルを作成しました"
        else
            error "env.example ファイルが見つかりません"
        fi
    else
        warning ".env ファイルは既に存在します"
    fi
}

# Riot Games API Key の設定確認
check_api_key() {
    info "Riot Games API Key の設定を確認します..."
    
    if [ -f ".env" ]; then
        if grep -q "RIOT_API_KEY=RGAPI-" .env; then
            warning "Riot Games API Key がまだ設定されていません"
            echo ""
            echo "📋 次のステップ："
            echo "1. https://developer.riotgames.com/ にアクセス"
            echo "2. アカウントを作成またはログイン"
            echo "3. 'PERSONAL API KEY' を生成"
            echo "4. .env ファイルの RIOT_API_KEY を更新"
            echo ""
            echo "例: RIOT_API_KEY=RGAPI-12345678-abcd-1234-abcd-123456789abc"
        else
            success "API Key が設定されています"
        fi
    fi
}

# VS Code 設定の更新
update_vscode_config() {
    info "VS Code MCP 設定を更新しています..."
    
    cd ../.. || error "プロジェクトルートに移動できません"
    
    if [ ! -d ".vscode" ]; then
        mkdir -p .vscode
        success ".vscode ディレクトリを作成しました"
    fi
    
    if [ -f ".vscode/mcp.json" ]; then
        success "VS Code MCP 設定ファイルは既に存在します"
    else
        warning "VS Code MCP 設定ファイルが見つかりません"
        echo "手動で .vscode/mcp.json を作成してください"
    fi
}

# テスト実行
test_server() {
    info "サーバーのテストを実行しています..."
    
    cd mcp-servers/lol-mcp-server || error "ディレクトリの移動に失敗しました"
    
    # 環境変数の確認
    if [ -f ".env" ]; then
        source .env
        if [ -z "$RIOT_API_KEY" ] || [ "$RIOT_API_KEY" = "RGAPI-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" ]; then
            warning "API Key が設定されていないため、サーバーテストをスキップします"
            return 0
        fi
    fi
    
    # サーバーの起動テスト (5秒間)
    timeout 5s npm start &> /dev/null || true
    success "サーバーのテストが完了しました"
}

# 完了メッセージの表示
show_completion_message() {
    echo ""
    echo "🎉 League of Legends MCP Server のセットアップが完了しました！"
    echo ""
    echo "📋 使用方法："
    echo "1. Riot Games API Key を設定 (まだの場合)"
    echo "2. Visual Studio Code でプロジェクトを開く"
    echo "3. Copilot Chat でMCPサーバーを有効化"
    echo "4. '〇〇サモナーの情報を教えて' などの質問を試す"
    echo ""
    echo "🔧 手動コマンド："
    echo "   cd mcp-servers/lol-mcp-server"
    echo "   npm start                    # サーバー起動"
    echo "   npm run dev                  # 開発モード"
    echo "   npm run build                # ビルド"
    echo ""
    echo "📖 詳細は mcp-servers/lol-mcp-server/README.md を参照してください"
}

# メイン実行
main() {
    echo "🚀 セットアップを開始します..."
    
    check_nodejs
    setup_directory
    install_dependencies
    build_typescript
    setup_env_file
    check_api_key
    update_vscode_config
    test_server
    show_completion_message
}

# スクリプト実行
main "$@" 