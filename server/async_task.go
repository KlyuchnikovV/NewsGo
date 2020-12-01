package server

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)


type asyncTask struct {
	ctx    context.Context
	cancel context.CancelFunc
	timer  *time.Timer
	name   string
	onTick func() error
}

func newAsyncTask(ctx context.Context, duration time.Duration, name string, onTick func() error) *asyncTask {
	return &asyncTask{
		ctx:    ctx,
		timer:  time.NewTimer(duration),
		name:   name,
		onTick: onTick,
	}
}

func (a *asyncTask) Start() {
	if a.cancel != nil {
		logrus.Warnf("'%s' task was already started", a.name)
		return
	}
	if a.onTick == nil {
		logrus.Warnf("'%s' task onTick is nil, can't start", a.name)
		return
	}
	a.ctx, a.cancel = context.WithCancel(a.ctx)
	for {
		select {
		case <-a.timer.C:
			if err := a.onTick(); err != nil {
				logrus.Errorf("'%s' task thrown error (cause: '%s')", a.name, err.Error())
			}
		case <-a.ctx.Done():
			logrus.Tracef("'%s' task stopped", a.name)
			return
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
