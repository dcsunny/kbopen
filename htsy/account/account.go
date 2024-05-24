package account

import (
	"github.com/dcsunny/kbopen/context"
	"github.com/dcsunny/kbopen/http"
)

const (
	infoUrl = "/api/htsy/account/info"
)

type Account struct {
	ctx *context.Context
}

func NewAccount(ctx *context.Context) *Account {
	return &Account{
		ctx: ctx,
	}
}

type InfoData struct {
	AccountId int    `json:"accountId"`
	Nickname  string `json:"nickname"`
	Username  string `json:"username"`
	Config    struct {
		CmdCallbackLink string `json:"cmdCallbackLink"`
		MsgCallbackLink string `json:"msgCallbackLink"`
	} `json:"config"`
}

func (a *Account) Info() (*InfoData, error) {
	req := http.Request{AuthorizerUserId: a.ctx.AuthorizerUserId}
	var infoData InfoData
	_, err := a.ctx.HttpClient.HttpPostJson(infoUrl, req, &infoData, a.ctx.GetAccessToken)
	if err != nil {
		return nil, err
	}
	return &infoData, nil
}
