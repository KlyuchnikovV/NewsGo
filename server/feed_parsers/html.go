package feed_parsers

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"main/database"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type HtmlParser struct {
	url     string
	timeout time.Duration
	client  *resty.Client
	rule    string
}

func NewHtmlParser(url string, timeout time.Duration, rule string) *HtmlParser {
	return &HtmlParser{
		url:     url,
		timeout: timeout,
		client: resty.New().SetTransport(&http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}),
		rule: rule,
	}
}

func (h *HtmlParser) Fetch(ctx context.Context, db *gorm.DB) error {
	logrus.Tracef("Fetching '%s'...", h.url)

	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	response, err := h.client.R().
		SetContext(ctx).
		Get(h.url)
	if err != nil {
		return err
	}

	if response.StatusCode() != http.StatusOK {
		return fmt.Errorf("status code error: %s %s", response.Status(), response.Body())
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(response.Body()))
	if err != nil {
		return err
	}

	doc.Find(h.rule).Each(func(i int, s *goquery.Selection) {
		link, ok := s.Find("a").Attr("href")
		if !ok || len(link) == 0 {
			logrus.Warn("link wasn't found")
		} else if link[0] == '/' {
			link = h.url + link
		}
		title := s.Find("a,\nspan,\nH3,\nH2,\nH1\nstrong").Text()

		if len(link) == 0 || len(title) == 0 {
			return
		}

		if err := database.CreateItem(db, database.Item{
			Title:       title,
			Link:        link,
			Description: title,
		}); err != nil {
			logrus.Errorf("can't save item in DB (cause: %s)", err.Error())
		}
	})
	return nil
}
