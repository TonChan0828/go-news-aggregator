package aggregate

import (
	"testing"
	"time"

	"go-news-aggregator/internal/model"
)

func TestCountBySource(t *testing.T) {
	items := []model.Item{
		{Source: "NHK"},
		{Source: "NHK"},
		{Source: "ITmedia"},
	}

	agg := New(items)
	result := agg.CountBySource()

	if result["NHK"] != 2 {
		t.Errorf("expected NHK=2, got %d", result["NHK"])
	}
	if result["ITmedia"] != 1 {
		t.Errorf("expected ITmedia=1, got %d", result["ITmedia"])
	}
}

func TestSortByPublishedAtDesc(t *testing.T) {
	t1 := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 1, 2, 10, 0, 0, 0, time.UTC)

	items := []model.Item{
		{Title: "old", PublishedAt: &t1},
		{Title: "nil", PublishedAt: nil},
		{Title: "new", PublishedAt: &t2},
	}

	agg := New(items)
	sorted := agg.SortByPublishedAtDesc()

	if sorted[0].Title != "new" {
		t.Errorf("expected newest first, got %s", sorted[0].Title)
	}
	if sorted[len(sorted)-1].Title != "nil" {
		t.Errorf("expected nil published_at last")
	}
}
