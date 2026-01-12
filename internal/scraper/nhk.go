package scraper

import "context"

type NHKScraper struct{}

func (n *NHKScraper) Name() string {
	return "NHK"
}

func (n *NHKScraper) Fetch(ctx context.Context) ([]model.RawItem, error) {
	panic("not implemented")
}
