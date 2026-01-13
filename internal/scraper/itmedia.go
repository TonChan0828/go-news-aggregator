package scraper

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"go-news-aggregator/internal/model"

	"github.com/PuerkitoBio/goquery"
)

type ITmediaScraper struct{}

func (i *ITmediaScraper) Name() string {
	return "ITmedia"
}

func (i *ITmediaScraper) Fetch(ctx context.Context) ([]model.RawItem, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.itmedia.co.jp/news/", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("itmedia: status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var items []model.RawItem
	doc.Find(".colBoxTitle").Each(func(_ int, s *goquery.Selection) {
		a := s.Find("a").First()
		title := strings.TrimSpace(a.Text())
		link, ok := a.Attr("href")

		if title == "" || !ok || link == "" {
			return
		}

		url := link
		if !strings.HasPrefix(link, "http") {
			url = "https://www.itmedia.co.jp/" + link
		}

		items = append(items, model.RawItem{
			Title:       title,
			URL:         url,
			PublishedAt: nil,
			Source:      "ITmedia",
		})
	})

	return items, nil
}
