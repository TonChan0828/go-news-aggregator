package orchestrator

import ("context"
	"go-news-aggregator/internal/model"
	"go-news-aggregator/internal/scraper"
)

type Runner struct {
	Scrapers []scraper.Scraper
	Parallel int64
}

func (r *Runner) Run(ctx context.Context) ([]model.Item, error) {
	panic("not implemented")
}
