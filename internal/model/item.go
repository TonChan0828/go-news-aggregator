package model

import "time"

type RawItem struct {
	Title       string
	URL         string
	PublishedAt *time.Time
	Source      string
}

type Item struct {
	Title       string
	URL         string
	PublishedAt *time.Time
	Source      string
}
