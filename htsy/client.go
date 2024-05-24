package htsy

import (
	"github.com/dcsunny/kbopen/conf"
	"github.com/dcsunny/kbopen/context"
	"github.com/dcsunny/kbopen/htsy/account"
)

type Client struct {
	Ctx *context.Context
}

func NewClient(cfg *conf.Config, authorizerUserId string) *Client {
	ctx := context.NewContext(cfg)
	ctx.AuthorizerUserId = authorizerUserId
	return &Client{
		Ctx: ctx,
	}
}

func (c *Client) Account() *account.Account {
	return account.NewAccount(c.Ctx)
}
