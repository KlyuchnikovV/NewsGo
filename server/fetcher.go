package server

import (
	"context"
	"main/database"
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

func newFetcher(db *gorm.DB, url string, duration, timeout time.Duration) *fetcher {
	var f = fetcher{
		parser:  gofeed.NewParser(),
		url:     url,
		timeout: timeout,
		db:      db,
	}
	f.asyncTask = *newAsyncTask(duration, url, f.fetchUrl)
	return &f
}

func (f *fetcher) fetchUrl() error {
	logrus.Tracef("Fetching '%s'...", f.url)

	ctx, cancel := context.WithTimeout(f.ctx, f.timeout)
	defer cancel()

	parsedFeed, err := f.parser.ParseURLWithContext(f.url, ctx)
	if err != nil {
		logrus.Errorf("can't fetch url '%s' (cause: %s)", f.url, err.Error())
		return err
	}

	logrus.Tracef("Feed saved '%s'", parsedFeed.Link)
	// TODO: move to a func
	var feed = database.Feed{
		Title:       parsedFeed.Title,
		Description: parsedFeed.Description,
		Link:        parsedFeed.Link,
		FeedLink:    parsedFeed.FeedLink,
		Items:       make([]database.FeedItems, len(parsedFeed.Items)),
	}

	for i, item := range parsedFeed.Items {
		feed.Items[i].Item = database.Item{
			Title:       item.Title,
			Description: item.Description,
			Content:     item.Content,
			Link:        item.Link,
		}

		if item.Author != nil {
			feed.Items[i].Item.Author = database.Author{
				Name:  parsedFeed.Author.Name,
				Email: parsedFeed.Author.Email,
			}
		}
	}

	if err := database.CreateFeed(f.db, feed); err != nil {
		logrus.Errorf("can't save feed '%s' (cause: %s)", f.url, err.Error())
		return err
	}

	return nil
}
