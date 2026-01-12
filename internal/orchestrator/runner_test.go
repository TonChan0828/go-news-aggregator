package orchestrator

import (
	"context"
	"testing"
	"time"

	"go-news-aggregator/internal/model"
	"go-news-aggregator/internal/scraper"
)

// テスト専用 Fake Scraper
type fakeScraper struct {
	name  string
	delay time.Duration
}

func (f *fakeScraper) Name() string {
	return f.name
}

func (f *fakeScraper) Fetch(ctx context.Context) ([]model.RawItem, error) {
	select {
	case <-time.After(f.delay):
		return []model.RawItem{
			{
				Title:  f.name,
				Source: f.name,
			},
		}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func TestRunner_Run(t *testing.T) {
	scrapers := []scraper.Scraper{
		&fakeScraper{name: "A", delay: 10 * time.Millisecond},
		&fakeScraper{name: "B", delay: 10 * time.Millisecond},
	}

	runner := Runner{
		Scrapers: scrapers,
		Parallel: 1,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	items, err := runner.Run(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}
}
