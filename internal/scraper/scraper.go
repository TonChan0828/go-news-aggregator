package scraper

import "context"

type Scraper interface {
	Name() string
	Fetch(ctx context.Context) ([]model.RawItem, error)
}
