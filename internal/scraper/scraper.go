package scraper

import (
	"context"

	"go-news-aggregator/internal/model"
)

type Scraper interface {
	Name() string
	Fetch(ctx context.Context) ([]model.RawItem, error)
}
