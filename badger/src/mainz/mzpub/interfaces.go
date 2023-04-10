package mzpub

import (
	"context"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Main interface {
	OnLateInit() error
	OnRun() error
	OnExit() error
	OnCycle() error
	Vpr() *viper.Viper
	FailIt(msg string, err error)
	L() *logrus.Logger
	CreateCache(retry bool) (*cache.Cache[string], error)
	GetCache(retry bool) (*cache.Cache[string], error)
	CacheGet(ctx context.Context, key string) (string, error)
	CacheSet(ctx context.Context, key string, value string, options ...store.Option) error
	CacheDel(ctx context.Context, key string) error
}
