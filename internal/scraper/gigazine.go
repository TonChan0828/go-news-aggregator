package scraper

import (
	"context"
	"fmt"
	"go-news-aggregator/internal/model"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type GIGAZINEScraper struct{}

func (g *GIGAZINEScraper) Name() string {
	return "GIGAZINE"
}

func (g *GIGAZINEScraper) Fetch(ctx context.Context) ([]model.RawItem, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://gigazine.net/", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gigazine: status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var items []model.RawItem

	// GIGAZINE の記事一覧は article > h2 > a 構造が安定
	doc.Find(".content h2 a").Each(func(_ int, a *goquery.Selection) {
		title := strings.TrimSpace(a.Text())
		link, ok := a.Attr("href")

		if title == "" || !ok || link == "" {
			return
		}

		// GIGAZINE は href が絶対URL
		items = append(items, model.RawItem{
			Title:       title,
			URL:         link,
			PublishedAt: nil, // 日時は後回しでOK
			Source:      "GIGAZINE",
		})
	})

	return items, nil
}
