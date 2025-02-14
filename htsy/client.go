package htsy

import (
	"io"
	"net/http"
	"net/url"

	"github.com/dcsunny/kbopen/conf"
	"github.com/dcsunny/kbopen/context"
	"github.com/dcsunny/kbopen/htsy/account"
	"github.com/dcsunny/kbopen/htsy/bot"
	"github.com/dcsunny/kbopen/htsy/callback"
	"github.com/dcsunny/kbopen/htsy/material"
	"github.com/dcsunny/kbopen/htsy/message"
	"github.com/dcsunny/kbopen/htsy/room"
)

type Client struct {
	Ctx *context.Context
}

func NewClient(cfg *conf.Config) *Client {
	ctx := context.NewContext(cfg)
	return &Client{
		Ctx: ctx,
	}
}

func (c *Client) Account() *account.Account {
	return account.NewAccount(c.Ctx)
}

func (c *Client) Bot() *bot.Bot {
	return bot.NewBot(c.Ctx)
}

func (c *Client) Room() *room.Room {
	return room.NewRoom(c.Ctx)
}

func (c *Client) Msg() *message.Msg {
	return message.NewMsg(c.Ctx)
}

func (c *Client) Material() *material.Material {
	return material.NewMaterial(c.Ctx)
}

func (c *Client) Callback(body []byte, query url.Values) *callback.Callback {
	srv := callback.NewCallback(c.Ctx)
	srv.RequestBody = body
	srv.UrlQuery = query
	return srv
}

func (c *Client) CallbackByHttp(req *http.Request, writer http.ResponseWriter) *callback.Callback {
	srv := callback.NewCallback(c.Ctx)
	srv.UrlQuery = req.URL.Query()
	srv.Writer = writer
	srv.RequestBody, _ = io.ReadAll(req.Body)
	return srv
}
