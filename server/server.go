package server

import (
	"context"
	"fmt"
	"main/database"
	"main/models"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	ctx      context.Context
	cancel   context.CancelFunc
	fetchers []*feedCollector
	timeout  time.Duration
	db       *gorm.DB
	models.UnimplementedRssServer
}

func New(ctx context.Context, db *gorm.DB, timeout time.Duration) (*Server, error) {
	urls, err := database.GetUrls(db)
	if err != nil {
		return nil, err
	}
	s := Server{
		ctx:      ctx,
		db:       db,
		timeout:  timeout,
		fetchers: make([]*feedCollector, len(urls)),
	}
	for i, url := range urls {
		s.fetchers[i] = newFetcher(db, url.Url, url.Duration, timeout, url.ParsingRule)
	}
	return &s, nil
}

func (s *Server) Ping(_ context.Context, _ *empty.Empty) (*empty.Empty, error) {
	logrus.Info("Pinged")
	return &empty.Empty{}, nil
}

func (s *Server) Start(_ context.Context, _ *empty.Empty) (*empty.Empty, error) {
	if s.cancel != nil {
		logrus.Warn("server was already started")
		return nil, fmt.Errorf("server already started")
	}
	s.ctx, s.cancel = context.WithCancel(s.ctx)
	for i := range s.fetchers {
		go s.runFetcher(s.ctx, i)
	}
	return &empty.Empty{}, nil
}

func (s *Server) Stop(_ context.Context, _ *empty.Empty) (*empty.Empty, error) {
	if s.cancel == nil {
		logrus.Warnf("server was already stopped")
		return nil, fmt.Errorf("server was already stopped")
	}
	logrus.Tracef("stopping server")
	s.cancel()
	s.cancel = nil
	return &empty.Empty{}, nil
}

func (s *Server) AddRss(ctx context.Context, link *models.RssLink) (*empty.Empty, error) {
	var duration = link.Duration.AsTime().Sub(time.Unix(0, 0))
	if err := database.AddRss(s.db, link.Url, duration, nil); err != nil {
		return nil, err
	}

	s.fetchers = append(s.fetchers, newFetcher(s.db, link.Url, duration, s.timeout, nil))
	go s.runFetcher(s.ctx, len(s.fetchers)-1)
	return &empty.Empty{}, nil
}

func (s *Server) AddUrl(ctx context.Context, link *models.UrlLink) (*empty.Empty, error) {
	var duration = link.Duration.AsTime().Sub(time.Unix(0, 0))
	if err := database.AddRss(s.db, link.Url, duration, &link.Rule); err != nil {
		return nil, err
	}
	s.fetchers = append(s.fetchers, newFetcher(s.db, link.Url, duration, s.timeout, &link.Rule))
	go s.runFetcher(s.ctx, len(s.fetchers)-1)
	return &empty.Empty{}, nil
}

func (s *Server) runFetcher(ctx context.Context, j int) {
	if err := s.fetchers[j].Start(ctx); err != nil {
		logrus.Errorf("starting fetcher failed (cause '%s')", err.Error())
	}
}

func (s *Server) ListNews(_ context.Context, _ *empty.Empty) (*models.News, error) {
	items, err := database.ListNews(s.db)
	if err != nil {
		return nil, err
	}
	var news = models.News{
		Articles: make([]*models.Article, len(items)),
	}
	for i, feed := range items {
		news.Articles[i] = &models.Article{
			Title: feed.Title,
			Url:   feed.Link,
		}
	}
	return &news, nil
}

func (s *Server) GetNews(_ context.Context, request *models.GetRequest) (*models.News, error) {
	items, err := database.GetNews(s.db, request.Request)
	if err != nil {
		return nil, err
	}
	var news = models.News{
		Articles: make([]*models.Article, len(items)),
	}
	for i, feed := range items {
		news.Articles[i] = &models.Article{
			Title: feed.Title,
			Text:  feed.Description,
			Url:   feed.Link,
		}
	}
	return &news, nil
}
