package orchestrator

import (
	"context"
	"go-news-aggregator/internal/model"
	"go-news-aggregator/internal/scraper"
	"sync"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

type Runner struct {
	Scrapers []scraper.Scraper
	Parallel int64
}

func (r *Runner) Run(ctx context.Context) ([]model.Item, error) {
	g, ctx := errgroup.WithContext(ctx)
	sem := semaphore.NewWeighted(r.Parallel)

	var (
		mu    sync.Mutex
		items []model.Item
	)

	for _, s := range r.Scrapers {
		s := s

		g.Go(func() error {

			if err := sem.Acquire(ctx, 1); err != nil {
				return err
			}
			defer sem.Release(1)

			rawItems, err := s.Fetch(ctx)
			if err != nil {
				// 今回は学習用なのでエラーしても続行
				return nil
			}

			// RawItem を Item に変換して追加
			var normalized []model.Item
			for _, r := range rawItems {
				normalized = append(normalized, model.Item{
					Title:       r.Title,
					URL:         r.URL,
					PublishedAt: r.PublishedAt,
					Source:      r.Source,
				})
			}

			mu.Lock()
			items = append(items, normalized...)
			mu.Unlock()

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	return items, nil
}
