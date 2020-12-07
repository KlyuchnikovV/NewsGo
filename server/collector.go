package server

import (
	"context"
	"main/server/feed_parsers"
	"time"
	"gorm.io/gorm"
)

type Fetcher interface {
	Fetch(context.Context, *gorm.DB) error
}

type feedCollector struct {
	asyncTask
	url     string
	timeout time.Duration
	db      *gorm.DB
}

func newFetcher(db *gorm.DB, url string, duration, timeout time.Duration, parseRule *string) *feedCollector {
	var f = feedCollector{
		url:     url,
		timeout: timeout,
		db:      db,
	}

	var fetcher Fetcher
	if parseRule == nil {
		fetcher = feed_parsers.NewRssParser(url, timeout)
	} else {
		fetcher = feed_parsers.NewHtmlParser(url, timeout, *parseRule)
	}
	f.asyncTask = *newAsyncTask(duration, url, db, fetcher)

	return &f
}
