package credential

import (
	"fmt"
	"sync"
	"time"

	"github.com/dcsunny/kbopen/cache"
	"github.com/dcsunny/kbopen/http"
)

type AccessTokenHandle interface {
	GetAccessToken() (accessToken string, err error)
}

const (
	getAccessTokenUrl = "/api/get_access_token"
)

type DefaultAccessToken struct {
	appID      string
	appSecret  string
	cache      cache.Cache
	httpClient *http.HttpClient
	mu         sync.Mutex
}

func NewDefaultAccessToken(appID string,
	appSecret string,
	_cache cache.Cache, httpClient *http.HttpClient) *DefaultAccessToken {
	return &DefaultAccessToken{
		appID:      appID,
		appSecret:  appSecret,
		cache:      _cache,
		httpClient: httpClient,
	}
}

type GetServerTokenReq struct {
	Appid     string `json:"appid"`
	AppSecret string `json:"appSecret"`
}

type GetServerTokenResp struct {
	AccessToken string `json:"accessToken"`
	ExpireTime  int64  `json:"expireTime"`
}

func (a *DefaultAccessToken) GetRemoteServerToken() (*GetServerTokenResp, error) {
	req := GetServerTokenReq{
		Appid:     a.appID,
		AppSecret: a.appSecret,
	}
	_, result, err := a.httpClient.HttpPostJson(getAccessTokenUrl, req, &GetServerTokenResp{}, nil)
	if err != nil {
		return nil, err
	}
	return result.(*GetServerTokenResp), nil
}

func (a *DefaultAccessToken) GetAccessToken() (string, error) {
	accessTokenCacheKey := fmt.Sprintf("kbopen_access_token_%s", a.appID)
	val := a.cache.Get(accessTokenCacheKey)
	if val != nil {
		return val.(string), nil
	}

	a.mu.Lock()
	defer a.mu.Unlock()

	val = a.cache.Get(accessTokenCacheKey)
	if val != nil {
		return val.(string), nil
	}

	// cache 失效，从远端服务刷新
	r, err := a.GetRemoteServerToken()
	if err != nil {
		return "", err
	}

	cacheTTL := tokenCacheTTL(r.ExpireTime, time.Now())
	if cacheTTL <= 0 {
		return r.AccessToken, nil
	}

	err = a.cache.Set(accessTokenCacheKey, r.AccessToken, cacheTTL)
	if err != nil {
		return "", nil
	}
	return r.AccessToken, nil
}

func tokenCacheTTL(expireTime int64, now time.Time) time.Duration {
	expireAt := normalizeExpireTime(expireTime)
	remaining := expireAt.Sub(now)
	if remaining <= 0 {
		return 0
	}

	const safetyBuffer = 10 * time.Minute
	if remaining > safetyBuffer {
		return remaining - safetyBuffer
	}

	// 短有效期 token 仍然缓存一段时间，避免并发请求反复刷新。
	if remaining > time.Second {
		return remaining / 2
	}

	return remaining
}

func normalizeExpireTime(expireTime int64) time.Time {
	if expireTime > 1e11 {
		return time.UnixMilli(expireTime)
	}
	return time.Unix(expireTime, 0)
}
