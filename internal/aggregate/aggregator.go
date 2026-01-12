package aggregate

import (
	"sort"

	"go-news-aggregator/internal/model"
)

type Aggregator struct {
	items []model.Item
}

func New(items []model.Item) *Aggregator {
	return &Aggregator{items: items}
}

// ソース別の記事件数を返す
func (a *Aggregator) CountBySource() map[string]int {
	result := make(map[string]int)

	for _, item := range a.items {
		result[item.Source]++
	}

	return result
}

// 公開日時の降順で記事を返す
// published_atがnilの記事は最後にまとめて返す
func (a *Aggregator) SortByPublishedAtDesc() []model.Item {
	sorted := make([]model.Item, len(a.items))
	copy(sorted, a.items)

	sort.Slice(sorted, func(i, j int) bool {
		pi := sorted[i].PublishedAt
		pj := sorted[j].PublishedAt

		if pi == nil && pj == nil {
			return false
		}

		if pi == nil {
			return false
		}
		if pj == nil {
			return true
		}

		return pi.After(*pj)
	})

	return sorted
}
