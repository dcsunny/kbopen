package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/dcsunny/kbopen/conf"
)

type HttpClient struct {
	cfg            *conf.Config
	HttpLastResult HttpLastResult
}

type HttpLastResult struct {
	Body []byte
	Err  error
}

func NewHttpClient(cfg *conf.Config) *HttpClient {
	return &HttpClient{
		cfg: cfg,
	}
}

func (c *HttpClient) HttpPostJson(link string, req interface{}, resp interface{}, tokenFunc func() (string, error)) (CommonResult, error) {
	var err error
	var result CommonResultWithData
	link, err = url.JoinPath(c.cfg.Endpoint, link)
	if err != nil {
		return result.CommonResult, err
	}

	client := http.Client{
		Timeout: time.Second * time.Duration(c.cfg.HttpTimeout),
	}
	j, _ := json.Marshal(req)
	httpReq, err := http.NewRequest(http.MethodPost, link, bytes.NewReader(j))
	if err != nil {
		return result.CommonResult, err
	}
	httpReq.Header.Set("content-type", "application/json")
	if tokenFunc != nil {
		var token string
		token, err = tokenFunc()
		if err != nil {
			return CommonResult{}, err
		}
		httpReq.Header.Set("Authorization", "Bearer "+token)
	}
	var r *http.Response
	r, err = client.Do(httpReq)
	defer func() {
		c.HttpLastResult.Err = err
		if r.Body != nil {
			r.Body.Close()
		}
	}()
	if err != nil {
		return result.CommonResult, err
	}
	var body []byte
	body, err = io.ReadAll(r.Body)
	if err != nil {
		return result.CommonResult, err
	}
	c.HttpLastResult.Body = body
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result.CommonResult, err
	}
	if result.CommonResult.Code != 0 {
		return result.CommonResult, errors.New(result.CommonResult.Message)
	}
	if resp == nil {
		return result.CommonResult, nil
	}
	err = json.Unmarshal(result.Data, resp)
	if err != nil {
		return result.CommonResult, err
	}
	return result.CommonResult, nil
}
