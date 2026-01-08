# 設計メモ

## ファイル構造

```
.
├── cmd/
│   └── scraper/
│       └── main.go
├── docs/
│   ├── design.md
│   └── requirements.md
└── internal/
    ├── aggregate/
    │   └── aggregator.go
    ├── model/
    │   └── item.go
    ├── orchestrator/
    │   └── runner.go
    └── scraper/
        ├── itmedia.go
        ├── nhk.go
        └── scraper.go
```

## パッケージ責務

### cmd/scraper
CLIのエントリーポイント。引数や実行オプションを読み取り、内部パッケージを組み合わせて処理を開始する。

### internal/model
取得した記事のドメインモデルを定義する。記事タイトル、URL、公開日時、ソース名などを扱う。

### internal/scraper
スクレイピングのインターフェースと実装を提供する `scraper.go` に並行処理を担
* `nhk.go` と `itmedia.go` に「1サイトから記事を取得する責務」に限定

### internal/orchestrator
* scraper の実行順序・並行数を制御する
* context によるキャンセル・タイムアウトを一元管理する
* scraper の実装詳細や集計ロジックには関知しない

### internal/aggregate
取得結果の集計と整形を担当する。ソース別件数の集計や、公開日時順のソートなど表示に必要なデータ加工を行う。
