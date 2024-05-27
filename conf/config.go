package conf

import "github.com/dcsunny/kbopen/cache"

type Config struct {
	Endpoint  string //节点 默认值 https://open-api.cpshelp.cn
	Appid     string //应用appid
	AppSecret string //应用secret

	AuthorizerUserId string //授权应用的id

	HttpTimeout int64 //http请求超时的设置 单位秒

	RedisOpts *cache.RedisOpts
}

func Default(cfg *Config) {
	if cfg == nil {
		cfg = new(Config)
	}
	if cfg.HttpTimeout <= 0 || cfg.HttpTimeout > 1000 {
		cfg.HttpTimeout = 60
	}
	if cfg.Endpoint == "" {
		cfg.Endpoint = "https://open-api.cpshelp.cn"
	}
	return
}
