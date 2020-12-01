package server

import (
	"context"
	"main/database"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type fetcher struct {
	asyncTask
	parser  *gofeed.Parser
	url     string
	timeout time.Duration
	db      *gorm.DB
}

func newFetcher(ctx context.Context, db *gorm.DB, url string, duration, timeout time.Duration) *fetcher {
	var f = fetcher{
		parser:  gofeed.NewParser(),
		url:     url,
		timeout: timeout,
		db:      db,
	}
	f.asyncTask = *newAsyncTask(ctx, duration, url, f.fetchUrl)
	return &f
}

func (f *fetcher) fetchUrl() error {
	logrus.Tracef("fetching '%s'...", f.url)

	ctx, cancel := context.WithTimeout(f.ctx, f.timeout)
	defer cancel()

	parsedFeed, err := f.parser.ParseURLWithContext(f.url, ctx)
	if err != nil {
		logrus.Errorf("can't fetch url '%s' (cause: %s)", f.url, err.Error())
		return err
	}

	logrus.Infof("link saved '%s'", parsedFeed.Link)
	// TODO: move to a func
	var feed = database.Feed{
		Title:       parsedFeed.Title,
		Description: parsedFeed.Description,
		Link:        parsedFeed.Link,
		FeedLink:    parsedFeed.FeedLink,
		Updated:     parsedFeed.Updated,
		Published:   parsedFeed.Published,
		Author: database.Author{
			Name:  parsedFeed.Author.Name,
			Email: parsedFeed.Author.Email,
		},
		Language:   parsedFeed.Language,
		Categories: make([]database.FeedsCategories, len(parsedFeed.Categories)),
	}

	for i, category := range parsedFeed.Categories {
		feed.Categories[i] = database.FeedsCategories{
			Category: database.Category{Name: strings.Trim(category, " \t")},
		}
	}

	if result := f.db.Create(&feed); result.Error != nil {
		logrus.Errorf("can't save feed '%s' (cause: %s)", f.url, err.Error())
		return result.Error
	}

	return nil
}
