# go-news-aggregator
複数のニュースサイトを並列にスクレイピングして集計・表示するCLIツールです。
Goの並行処理（goroutine/context）や責務分離の設計を学習する目的で作られています。

## 特徴
- 取得対象: NHK / ITmedia / GIGAZINE
- 並行実行数とタイムアウトをCLIオプションで制御
- ソース別件数の集計と最新記事の時刻順表示

## 必要環境
- Go 1.25 以上

## 使い方
```bash
go run ./cmd/scraper -parallel=5 -timeout=5s
```

### オプション
- `-parallel`: 並行実行数（デフォルト: 5）
- `-timeout`: 全体のタイムアウト（デフォルト: 5s）

## 出力例
```
=== 記事件数（ソース別）===
NHK: 20
ITmedia: 30
GIGAZINE: 15

=== 最新記事 ===
-[NHK] 記事タイトル (2025-01-01T12:34:56Z)
 https://example.com/...
```

## ディレクトリ構成
```
cmd/scraper/           CLIエントリーポイント
docs/                  目的・設計メモ
internal/aggregate/    集計・並び替え
internal/model/        ドメインモデル
internal/orchestrator/ 並行実行の制御
internal/scraper/      各サイトのスクレイパー実装
```

## 設計メモ
- 目的や設計方針は `docs/requirements.md` と `docs/design.md` にまとめています。
