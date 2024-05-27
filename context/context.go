package context

import (
	"context"

	"github.com/dcsunny/kbopen/cache"
	"github.com/dcsunny/kbopen/conf"
	"github.com/dcsunny/kbopen/credential"
	"github.com/dcsunny/kbopen/http"
)

type Context struct {
	context.Context
	Cfg *conf.Config
	credential.AccessTokenHandle

	HttpClient *http.HttpClient
}

func NewContext(cfg *conf.Config) *Context {
	conf.Default(cfg)

	var _cache cache.Cache
	if cfg.RedisOpts == nil {
		_cache = cache.NewMemoryCache()
	} else {
		_cache = cache.NewRedis(cfg.RedisOpts)
	}
	httpClient := http.NewHttpClient(cfg)
	ctx := &Context{
		Context:           context.Background(),
		Cfg:               cfg,
		HttpClient:        httpClient,
		AccessTokenHandle: credential.NewDefaultAccessToken(cfg.Appid, cfg.AppSecret, _cache, httpClient),
	}

	return ctx
}
