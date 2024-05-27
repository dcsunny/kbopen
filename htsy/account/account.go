package account

import (
	"github.com/dcsunny/kbopen/context"
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

type InfoReq struct {
	AuthorizerUserId string `json:"authorizerUserId"`
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
	_, infoData, err := a.ctx.HttpClient.HttpPostJsonWithAuthorizer(infoUrl, nil, &InfoData{}, a.ctx.GetAccessToken)
	if err != nil {
		return nil, err
	}
	return infoData.(*InfoData), nil
}
