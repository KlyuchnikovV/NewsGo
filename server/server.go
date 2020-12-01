package server

import (
	"context"
	"main/database"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	ctx      context.Context
	cancel   context.CancelFunc
	fetchers []*fetcher
}

func New(ctx context.Context, db *gorm.DB, timeout time.Duration) *Server {
	var urls []database.RssUrls
	if result := db.Find(&urls); result.Error != nil {
		panic(result.Error)
	}
	s := Server{
		ctx:      ctx,
		fetchers: make([]*fetcher, len(urls)),
	}
	for i, url := range urls {
		s.fetchers[i] = newFetcher(ctx, db, url.Url, url.Duration, timeout)
	}
	return &s
}

func (s *Server) Start() {
	if s.cancel != nil {
		logrus.Warn("server was already started")
		return
	}
	s.ctx, s.cancel = context.WithCancel(s.ctx)
	for i := range s.fetchers {
		go s.fetchers[i].Start()
	}
}

func (s *Server) Stop() {
	if s.cancel == nil {
		logrus.Warnf("server was already stopped")
		return
	}
	logrus.Tracef("stopping server")
	s.cancel()
	s.cancel = nil
}
