package mainz

import (
	"context"
	"errors"
	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	ristretto_store "github.com/eko/gocache/store/ristretto/v4"
)

var (
	ErrCacheNotInited = errors.New("cache not init'ed yet")
)

// initCacheDefaults sets viper defaults - can be called many times
func (w MainWorker) initCacheDefaults() {
	w.Vpr().SetDefault("cache.ristretto.num-counters", 1000)
	w.Vpr().SetDefault("cache.ristretto.max-cost", 100)
	w.Vpr().SetDefault("cache.ristretto.buffer-items", 64)
}

func (w MainWorker) initCache() error {
	c, err := w.CreateCache()
	if err != nil {
		return err
	}
	w.cache = c
	return nil
}

// CreateCache creates a new cache object without setting it
func (w MainWorker) CreateCache() (*cache.Cache[string], error) {
	w.initCacheDefaults()
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: w.Vpr().GetInt64("cache.ristretto.num-counters"),
		MaxCost:     w.Vpr().GetInt64("cache.ristretto.max-cost"),
		BufferItems: w.Vpr().GetInt64("cache.ristretto.buffer-items"),
	})
	if err != nil {
		return nil, err
	}
	ristrettoStore := ristretto_store.NewRistretto(ristrettoCache)

	return cache.New[string](ristrettoStore), nil
}

// GetCache fetches the current cache, creating it if needed when retry==true
func (w MainWorker) GetCache(retry bool) (*cache.Cache[string], error) {
	if w.cache != nil {
		return w.cache, nil
	}
	if !retry {
		w.L().Warn("Cache not init'ed yet - told not to try early init")
		return nil, ErrCacheNotInited
	}
	w.L().Warn("Cache not init'ed yet - trying early init")
	if err := w.initCache(); err != nil {
		return nil, err
	}
	return w.GetCache(false)
}

// CacheGet is a simple string:string cache item fetch
func (w MainWorker) CacheGet(ctx context.Context, key string) (string, error) {
	c, err := w.GetCache(true)
	if err != nil {
		return "", err
	}
	value, err := c.Get(ctx, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

// CacheSet is a simple string:string cache item set
func (w MainWorker) CacheSet(ctx context.Context, key string, value string, options ...store.Option) error {
	c, err := w.GetCache(true)
	if err != nil {
		return err
	}
	if err := c.Set(ctx, key, value, options...); err != nil {
		return err
	}
	return nil
}

// CacheDel is a simple string:string cache item delete
func (w MainWorker) CacheDel(ctx context.Context, key string) error {
	c, err := w.GetCache(true)
	if err != nil {
		return err
	}
	if err := c.Delete(ctx, key); err != nil {
		return err
	}
	return nil
}
