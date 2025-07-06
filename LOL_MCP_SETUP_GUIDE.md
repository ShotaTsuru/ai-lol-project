# League of Legends MCP サーバー設定ガイド 🎮

League of Legends専用のModel Context Protocol (MCP) サーバーを使用して、AI アシスタントがLoL関連の情報を自然言語で取得・分析できるようにします。

## 🎯 概要

このMCPサーバーは以下の機能を提供します：

- **サモナー情報取得**: プレイヤーの基本情報とランク情報
- **マッチ履歴分析**: 最近の試合結果とパフォーマンス
- **チャンピオン情報**: 詳細な統計とスキル情報
- **アイテム情報**: 最新のアイテムデータ
- **パフォーマンス分析**: 特定チャンピオンでの成績評価
- **ゲーム情報**: 現在のパッチバージョンとメタ情報

## 📋 前提条件

- **Node.js 18+**
- **Riot Games API Key**
- **Visual Studio Code**
- **GitHub Copilot**

## 🚀 クイックスタート

### 1. 自動セットアップ

```bash
# 自動セットアップスクリプトを実行
make lol-mcp-setup
```

### 2. 手動セットアップ

```bash
# 1. プロジェクトディレクトリに移動
cd mcp-servers/lol-mcp-server

# 2. 依存関係をインストール
npm install

# 3. 環境変数を設定
cp env.example .env
# .envファイルを編集してRIOT_API_KEYを設定

# 4. TypeScriptをコンパイル
npm run build

# 5. サーバーを起動
npm start
```

## 🔑 Riot Games API Key の取得

### 1. Riot Developer Portal へのアクセス

1. [Riot Developer Portal](https://developer.riotgames.com/) にアクセス
2. Riot Games アカウントでログイン（League of Legends アカウントと同じ）

### 2. API Key の生成

1. **「PERSONAL API KEY」** セクションを見つける
2. 「**REGENERATE API KEY**」ボタンをクリック
3. 生成されたAPI Keyをコピー（例：`RGAPI-12345678-abcd-1234-abcd-123456789abc`）

### 3. API Key の設定

`.env` ファイルを編集：

```env
RIOT_API_KEY=RGAPI-あなたのAPIキー
```

⚠️ **注意**: Personal API Keyは24時間で期限切れになります。本格的な開発には Production API Key の申請が必要です。

## 🔧 Visual Studio Code での設定

### 1. MCP 設定ファイルの確認

`.vscode/mcp.json` ファイルが正しく設定されているか確認：

```json
{
  "mcpServers": {
    "lol-mcp-server": {
      "command": "node",
      "args": ["./mcp-servers/lol-mcp-server/dist/index.js"],
      "env": {
        "RIOT_API_KEY": "your-riot-api-key"
      }
    }
  }
}
```

### 2. VS Code での起動

1. **Visual Studio Code** でプロジェクトを開く
2. **Copilot Chat** を開く (サイドバーのアイコンまたは `Ctrl+Shift+I`)
3. **Agent** モードを選択
4. **Tools** セクションで **「lol-mcp-server」** が表示されることを確認

## 💡 使用方法

### 基本的な質問例

```
# サモナー情報の取得
「Hide on bushというサモナーの情報を教えて」

# ランク情報の確認
「Fakerの現在のランクを教えて」

# マッチ履歴の分析
「T1 Fakerの最近の試合を10試合分見せて」

# チャンピオン情報の取得
「ヤスオのチャンピオン情報を日本語で教えて」

# アイテム情報の確認
「無限の大剣というアイテムの情報を教えて」

# パフォーマンス分析
「プレイヤーXのアニビアでのパフォーマンスを分析して」

# ゲーム情報の確認
「現在のLoLのパッチバージョンは何？」
```

### 高度な分析例

```
# 詳細なパフォーマンス分析
「Hide on bushの最近20試合でのヤスオのパフォーマンスを分析して、勝率、平均KDA、平均ダメージを含めて教えて」

# 特定リージョンでの検索
「北アメリカサーバーでTSM Bjergsenの情報を取得して」

# 複数の情報を組み合わせた分析
「現在のパッチバージョンとFakerの最近の試合結果を組み合わせて、どのチャンピオンが強いか分析して」
```

## 🔍 利用可能なツール詳細

### 1. get_summoner_info
- **機能**: サモナーの基本情報を取得
- **パラメータ**: 
  - `summoner_name`: サモナー名
  - `region`: リージョン (kr, na1, euw1等)

### 2. get_rank_info
- **機能**: サモナーのランク情報を取得
- **パラメータ**: 
  - `summoner_id`: サモナーID
  - `region`: リージョン

### 3. get_recent_matches
- **機能**: 最近のマッチ履歴を取得
- **パラメータ**: 
  - `puuid`: プレイヤーUUID
  - `region`: リージョン
  - `count`: 取得する試合数

### 4. get_champion_info
- **機能**: チャンピオン情報を取得
- **パラメータ**: 
  - `champion_name`: チャンピオン名（英語）
  - `language`: 言語 (ja_JP, en_US, ko_KR)

### 5. get_item_info
- **機能**: アイテム情報を取得
- **パラメータ**: 
  - `item_id`: アイテムID
  - `language`: 言語

### 6. get_game_version
- **機能**: 現在のゲームバージョンを取得
- **パラメータ**: なし

### 7. analyze_champion_performance
- **機能**: チャンピオンのパフォーマンス分析
- **パラメータ**: 
  - `puuid`: プレイヤーUUID
  - `champion_name`: チャンピオン名
  - `region`: リージョン
  - `match_count`: 分析する試合数

## 🌍 対応リージョン

| リージョンコード | 地域 |
|---|---|
| kr | 韓国 |
| na1 | 北アメリカ |
| euw1 | ヨーロッパ西部 |
| eun1 | ヨーロッパ北東部 |
| br1 | ブラジル |
| jp1 | 日本 |
| ru | ロシア |
| oc1 | オセアニア |
| tr1 | トルコ |
| la1 | ラテンアメリカ北部 |
| la2 | ラテンアメリカ南部 |

## 🔧 メンテナンス

### サーバーの再起動

```bash
# 開発モードでの起動
make lol-mcp-dev

# 本番モードでの起動
make lol-mcp-start

# ビルドの実行
make lol-mcp-build
```

### 環境のクリーンアップ

```bash
# 完全なクリーンアップ
make lol-mcp-clean

# 再セットアップ
make lol-mcp-setup
```

## 📊 パフォーマンス最適化

### キャッシュ機能

- **キャッシュ時間**: 15分
- **対象**: 全てのAPI呼び出し結果
- **効果**: API制限の回避とレスポンス速度の向上

### レート制限対応

- **Personal API Key**: 100 requests/2 minutes
- **キャッシュ**: 同じリクエストの重複を防ぐ
- **エラーハンドリング**: レート制限エラーの適切な処理

## 🐛 トラブルシューティング

### よくある問題

#### 1. API Key エラー
```
❌ エラー: Unauthorized (401)
```
**解決策**:
- `.env` ファイルのAPI Keyが正しく設定されているか確認
- API Keyの有効期限を確認
- Riot Developer Portalで新しいAPI Keyを生成

#### 2. サモナー名が見つからない
```
❌ エラー: サモナー情報の取得に失敗しました
```
**解決策**:
- サモナー名の正確性を確認
- 適切なリージョンを指定
- 特殊文字や空白の処理を確認

#### 3. MCP サーバーが起動しない
```
❌ エラー: Failed to start server
```
**解決策**:
- Node.js バージョンを確認 (18+)
- 依存関係を再インストール: `npm install`
- TypeScript を再コンパイル: `npm run build`

#### 4. VS Code でツールが表示されない
**解決策**:
- Visual Studio Code を再起動
- Copilot Chat のAgent モードを確認
- `.vscode/mcp.json` の設定を確認

### デバッグモード

```bash
# デバッグモードで起動
DEBUG=true npm run dev
```

## 📈 今後の機能拡張

### 計画中の機能

- **リアルタイム試合情報**: 進行中の試合の詳細
- **メタ分析**: 現在のメタに基づく推奨ビルド
- **チーム分析**: チーム構成の最適化提案
- **統計可視化**: グラフやチャートでの統計表示
- **マルチプレイヤー分析**: 複数プレイヤーの比較分析

### 貢献方法

1. GitHub Issues で機能要求やバグ報告
2. Pull Request での機能追加
3. ドキュメントの改善提案

## 📄 ライセンス

MIT License

## 🤝 サポート

- **GitHub Issues**: バグ報告や機能要求
- **プロジェクトドキュメント**: 詳細な技術情報
- **コミュニティ**: 使用例やベストプラクティスの共有

---

**注意事項**: 
- このツールは非公式であり、Riot Games Inc. とは関係ありません
- Riot Games API の利用規約に従って使用してください
- Personal API Key は24時間で期限切れになります
- 本格的な開発には Production API Key の申請を推奨します 