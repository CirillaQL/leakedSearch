package cache

import (
	"github.com/allegro/bigcache/v3"
	"time"
)

var Cache *bigcache.BigCache

func init() {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	Cache = cache
}
