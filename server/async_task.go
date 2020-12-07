package server

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type asyncTask struct {
	ctx    context.Context
	db     *gorm.DB
	cancel context.CancelFunc
	timer  *time.Timer
	name   string
	// onTick func(context.Context, *gorm.DB) error
	Fetcher
}

func newAsyncTask(duration time.Duration, name string, db *gorm.DB, fetcher Fetcher) *asyncTask {
	return &asyncTask{
		timer:  time.NewTimer(duration),
		name:   name,
		db:     db,
		Fetcher: fetcher,
	}
}

func (a *asyncTask) Start(ctx context.Context) error {
	if a.cancel != nil {
		logrus.Warnf("'%s' task was already started", a.name)
		return fmt.Errorf("'%s' task was already started", a.name)
	}
	if a.Fetch == nil {
		logrus.Panicf("'%s' task onTick is nil, can't start", a.name)
	}
	a.ctx, a.cancel = context.WithCancel(ctx)
	for {
		select {
		case <-a.timer.C:
			if err := a.Fetch(a.ctx, a.db); err != nil {
				logrus.Errorf("'%s' task thrown error (cause: '%s')", a.name, err.Error())
			}
		case <-a.ctx.Done():
			logrus.Tracef("'%s' task stopped", a.name)
			return nil
		}
	}
}

func (a *asyncTask) Stop() {
	if a.cancel == nil {
		logrus.Warnf("'%s' task was already stopped", a.name)
		return
	}
	logrus.Tracef("stopping '%s' task", a.name)
	a.cancel()
	a.cancel = nil
}
