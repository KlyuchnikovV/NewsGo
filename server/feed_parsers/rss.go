package feed_parsers

import (
	"context"
	"main/database"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RssParser struct {
	url     string
	timeout time.Duration
	parser  *gofeed.Parser
}

func NewRssParser(url string, timeout time.Duration) *RssParser {
	return &RssParser{
		url:     url,
		timeout: timeout,
		parser:  gofeed.NewParser(),
	}
}

func (r *RssParser) Fetch(ctx context.Context, db *gorm.DB) error {
	logrus.Tracef("Fetching '%s'...", r.url)

	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	parsedFeed, err := r.parser.ParseURLWithContext(r.url, ctx)
	if err != nil {
		logrus.Errorf("can't fetch url '%s' (cause: %s)", r.url, err.Error())
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

	if err := database.CreateFeed(db, feed); err != nil {
		logrus.Errorf("can't save feed '%s' (cause: %s)", r.url, err.Error())
		return err
	}

	return nil
}
