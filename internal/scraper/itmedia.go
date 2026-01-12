package scraper

import (
	"context"

	"go-news-aggregator/internal/model"
)

type ITmediaScraper struct{}

func (i *ITmediaScraper) Name() string {
	return "ITmedia"
}

func (i *ITmediaScraper) Fetch(ctx context.Context) ([]model.RawItem, error) {
	panic("not implemented")
}
