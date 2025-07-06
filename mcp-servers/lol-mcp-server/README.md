# League of Legends MCP Server (Python版)

🎮 League of Legends専用のModel Context Protocol (MCP) サーバー

## 概要

このMCPサーバーは、League of Legendsの様々な情報を自然言語で取得・分析できるAIアシスタント用のツールです。Pythonの`fast-mcp`とRiot Games APIを使用します。

## 🚀 セットアップ

### 1. 依存パッケージのインストール
```bash
cd mcp-servers/lol-mcp-server
pip install -r requirements.txt
```

### 2. 環境変数の設定
```bash
cp env.example .env
# .envファイルを編集してRIOT_API_KEYを設定
```

### 3. サーバーの起動
```bash
python main.py
```

## 🔑 Riot Games API Key の取得
- [Riot Developer Portal](https://developer.riotgames.com/) で取得
- `.env`ファイルに `RIOT_API_KEY=...` を記載

## 💡 使用例

```
「Hide on bushサモナーの情報を教えて」
```

## 注意
- このサーバーはPythonで動作します。Node.js/TypeScript版ではありません。
- Riot Games API Keyは24時間で期限切れになります。

## 🚀 機能

### 利用可能なツール

1. **get_summoner_info** - サモナーの基本情報を取得
2. **get_rank_info** - サモナーのランク情報を取得
3. **get_recent_matches** - 最近のマッチ履歴を取得
4. **get_champion_info** - チャンピオン情報を取得
5. **get_item_info** - アイテム情報を取得
6. **get_game_version** - 現在のゲームバージョンを取得
7. **analyze_champion_performance** - チャンピオンのパフォーマンス分析

### 主な特徴

- 🔄 **キャッシュ機能**: API呼び出し結果を15分間キャッシュ
- 🌏 **マルチリージョン対応**: 世界各地域のデータに対応
- 📊 **統計分析**: 詳細なパフォーマンス分析
- 🎯 **日本語対応**: 日本語でのデータ取得

## 🎯 対応リージョン

- **kr** - 韓国
- **na1** - 北アメリカ
- **euw1** - ヨーロッパ西部
- **eun1** - ヨーロッパ北東部
- **br1** - ブラジル
- **jp1** - 日本
- **ru** - ロシア
- **oc1** - オセアニア
- **tr1** - トルコ
- **la1** - ラテンアメリカ北部
- **la2** - ラテンアメリカ南部

## 📊 パフォーマンス分析機能

### 提供される統計情報

- **勝率**: 指定チャンピオンでの勝率
- **平均KDA**: キル・デス・アシスト平均値
- **平均ダメージ**: 試合あたりの平均ダメージ量
- **平均ゴールド**: 試合あたりの平均ゴールド獲得量
- **パフォーマンス評価**: Excellent / Good / Average / Needs Improvement

## 🔄 キャッシュシステム

- **TTL**: 15分間
- **対象**: 全てのAPI呼び出し結果
- **利点**: API制限の回避とレスポンス速度の向上

## 🛠️ 開発

### 開発環境での実行

```bash
npm run dev
```

### ビルド

```bash
npm run build
```

### デバッグ

```bash
# 環境変数でデバッグモードを有効化
DEBUG=true npm run dev
```

## 📋 トラブルシューティング

### よくある問題

1. **API Key エラー**
   - Riot Games API Keyが正しく設定されているか確認
   - API Keyの有効期限を確認

2. **サモナー名が見つからない**
   - 正確なサモナー名を使用
   - 適切なリージョンを指定

3. **レート制限エラー**
   - 短時間に多くのリクエストを送信しすぎている
   - キャッシュの活用を推奨

## 🤝 コントリビューション

プルリクエストやイシューの報告を歓迎します。

## 📄 ライセンス

MIT License

## 🔗 関連リンク

- [Riot Games API Documentation](https://developer.riotgames.com/docs/lol)
- [Data Dragon API](https://developer.riotgames.com/docs/lol#data-dragon)
- [Model Context Protocol](https://modelcontextprotocol.io/)
- [Fast MCP](https://github.com/joshmu/fast-mcp)

---

**注意**: このツールは非公式のものであり、Riot Games Inc.とは関係ありません。使用にはRiot Games APIの利用規約に従う必要があります。
 