package material

import (
	"github.com/dcsunny/kbopen/context"
	"github.com/dcsunny/kbopen/http"
)

const (
	cdnDownloadUrl = "/api/htsy/material/cdnDownload"
)

type Material struct {
	ctx *context.Context
}

func NewMaterial(ctx *context.Context) *Material {
	return &Material{ctx: ctx}
}

type CdnDownloadRequest struct {
	http.AuthorizerUser
	CDNType int    `json:"cdnType"`
	QWCdn   QWCdn  `json:"qwCdn"`
	MsgType string `json:"msgType"`
	Wxid    string `json:"wxid"`
	BotType int    `json:"botType"`
}

// QWCdn 定义 qwCdn 字段的结构体
type QWCdn struct {
	AESKey   string `json:"aesKey,omitempty"`
	FileID   string `json:"fileId,omitempty"`
	Size     int    `json:"size,omitempty"`
	URL      string `json:"url,omitempty"`
	AuthKey  string `json:"authKey,omitempty"`
	FileName string `json:"fileName,omitempty"`
	MD5      string `json:"md5,omitempty"`
}

type CndDownloadResult struct {
	Data string `json:"data"`
}

func (m *Material) CdnDownload(req *CdnDownloadRequest) (string, error) {
	_, result, err := m.ctx.HttpClient.HttpPostJsonWithAuthorizer(cdnDownloadUrl, req, &CndDownloadResult{}, m.ctx.GetAccessToken)
	if err != nil {
		return "", err
	}
	return result.(*CndDownloadResult).Data, nil
}
