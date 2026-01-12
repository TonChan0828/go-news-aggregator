package scraper

import "context"

type ITmediaScraper struct{}

func (i *ITmediaScraper) Name() string {
	return "ITmedia"
}

func (i *ITmediaScraper) Fetch(ctx context.Context) ([]model.RawItem, error) {
	panic("not implemented")
}
