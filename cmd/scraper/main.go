package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"go-news-aggregator/internal/aggregate"
	"go-news-aggregator/internal/orchestrator"
	"go-news-aggregator/internal/scraper"
)

func main() {
	// option
	parallel := flag.Int("parallel", 5, "並行実行数")
	timeout := flag.Duration("timeout", 5*time.Second, "タイムアウト")
	flag.Parse()

	// context準備
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	// scraper構築
	scrapers := []scraper.Scraper{
		&scraper.NHKScraper{},
		&scraper.ITmediaScraper{},
		&scraper.GIGAZINEScraper{},
	}

	// 並行実行
	runner := orchestrator.Runner{
		Scrapers: scrapers,
		Parallel: int64(*parallel),
	}

	items, err := runner.Run(ctx)
	if err != nil {
		log.Fatalf("runner error: %v", err)
	}

	// 集計
	agg := aggregate.New(items)
	counts := agg.CountBySource()
	sorted := agg.SortByPublishedAtDesc()

	// 出力
	fmt.Println("=== 記事件数（ソース別）===")
	for source, count := range counts {
		fmt.Printf("%s: %d\n", source, count)
	}

	fmt.Println("\n=== 最新記事 ===")
	for i, item := range sorted {
		if i >= 100 {
			break
		}

		ts := "unknown"
		if item.PublishedAt != nil {
			ts = item.PublishedAt.Format(time.RFC3339)
		}

		fmt.Printf(
			"-[%s] %s (%s)\n %s\n",
			item.Source,
			item.Title,
			ts,
			item.URL,
		)
	}
}
