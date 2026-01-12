package scraper

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"go-news-aggregator/internal/model"

	"github.com/PuerkitoBio/goquery"
)

type NHKScraper struct{}

func (n *NHKScraper) Name() string {
	return "NHK"
}

func (n *NHKScraper) Fetch(ctx context.Context) ([]model.RawItem, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www3.nhk.or.jp/news/", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("nhk: status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var items []model.RawItem
	doc.Find("article").Each(func(_ int, s *goquery.Selection) {
		a := s.Find("a").First()
		title := strings.TrimSpace(a.Text())
		link, ok := s.Find("a").Attr("href")

		if title == "" || !ok || link == "" {
			return
		}

		url := link
		if !strings.HasPrefix(link, "http") {
			url = "https://www3.nhk.or.jp/news/" + link
		}

		items = append(items, model.RawItem{
			Title:       strings.TrimSpace(title),
			URL:         "https://www3.nhk.or.jp" + link,
			PublishedAt: nil,
			Source:      "NHK",
		})
	})

	return items, nil
}
