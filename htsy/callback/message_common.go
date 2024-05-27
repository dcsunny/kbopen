package callback

import "encoding/json"

// ParseMsgData 解析消息体
func ParseMsgData(body []byte, msg interface{}) interface{} {
	json.Unmarshal(body, msg)
	return msg
}

// Message 通用消息体
type Message struct {
	Type       int             `json:"type"`
	Handshake  string          `json:"handshake"` //校验的时候有用
	Trace      string          `json:"trace"`
	Data       json.RawMessage `json:"data"`
	BotWxid    string          `json:"botWxid"`
	BotType    int             `json:"botType"`
	Status     string          `json:"status"`
	FailReason string          `json:"failReason"`
}

// Text 文本
type Text struct {
	AtList         []string `json:"at_list"`
	Content        string   `json:"content"`
	ContentType    int      `json:"content_type"`
	ConversationId string   `json:"conversation_id"`
	IsPc           int      `json:"is_pc"`
	LocalId        string   `json:"local_id"`
	QuoteContent   string   `json:"quote_content"`
	Receiver       string   `json:"receiver"`
	SendTime       string   `json:"send_time"`
	Sender         string   `json:"sender"`
	SenderName     string   `json:"sender_name"`
	ServerId       string   `json:"server_id"`
}

// Image 图片
type Image struct {
	Appinfo string `json:"appinfo"`
	Cdn     struct {
		AesKey   string `json:"aes_key"`
		AuthKey  string `json:"auth_key"`
		FileId   string `json:"file_id"`
		FileName string `json:"file_name"`
		Height   int    `json:"height"`
		LdSize   int    `json:"ld_size"`
		LdUrl    string `json:"ld_url"`
		Md5      string `json:"md5"`
		MdSize   int    `json:"md_size"`
		MdUrl    string `json:"md_url"`
		Size     int    `json:"size"`
		Url      string `json:"url"`
		Width    int    `json:"width"`
	} `json:"cdn"`
	CdnType        int    `json:"cdn_type"`
	ContentType    int    `json:"content_type"`
	ConversationId string `json:"conversation_id"`
	IsPc           int    `json:"is_pc"`
	LocalId        string `json:"local_id"`
	Receiver       string `json:"receiver"`
	SendTime       string `json:"send_time"`
	Sender         string `json:"sender"`
	SenderName     string `json:"sender_name"`
	ServerId       string `json:"server_id"`
}

// Video 视频
type Video struct {
	Appinfo string `json:"appinfo"`
	Cdn     struct {
		AesKey         string `json:"aes_key"`
		AuthKey        string `json:"auth_key"`
		FileId         string `json:"file_id"`
		FileName       string `json:"file_name"`
		Md5            string `json:"md5"`
		PreviewImgSize int    `json:"preview_img_size"`
		PreviewImgUrl  string `json:"preview_img_url"`
		Size           int    `json:"size"`
		Url            string `json:"url"`
	} `json:"cdn"`
	CdnType        int    `json:"cdn_type"`
	ContentType    int    `json:"content_type"`
	ConversationId string `json:"conversation_id"`
	Duration       int    `json:"duration"`
	Height         int    `json:"height"`
	IsPc           int    `json:"is_pc"`
	LocalId        string `json:"local_id"`
	Receiver       string `json:"receiver"`
	SendTime       string `json:"send_time"`
	Sender         string `json:"sender"`
	SenderName     string `json:"sender_name"`
	ServerId       string `json:"server_id"`
	Width          int    `json:"width"`
}

// Audio 音频
type Audio struct {
	Appinfo string `json:"appinfo"`
	Cdn     struct {
		AesKey   string `json:"aes_key"`
		FileId   string `json:"file_id"`
		FileName string `json:"file_name"`
		Md5      string `json:"md5"`
		Size     int    `json:"size"`
	} `json:"cdn"`
	CdnType        int    `json:"cdn_type"`
	ContentType    int    `json:"content_type"`
	ConversationId string `json:"conversation_id"`
	Duration       int    `json:"duration"`
	IsPc           int    `json:"is_pc"`
	LocalId        string `json:"local_id"`
	Receiver       string `json:"receiver"`
	SendTime       string `json:"send_time"`
	Sender         string `json:"sender"`
	SenderName     string `json:"sender_name"`
	ServerId       string `json:"server_id"`
	VoiceTime      int    `json:"voice_time"`
}

// File 文件
type File struct {
	Appinfo string `json:"appinfo"`
	Cdn     struct {
		AesKey   string `json:"aes_key"`
		AuthKey  string `json:"auth_key"`
		FileName string `json:"file_name"`
		Md5      string `json:"md5"`
		Size     int    `json:"size"`
		Url      string `json:"url"`
	} `json:"cdn"`
	CdnType        int    `json:"cdn_type"`
	ContentType    int    `json:"content_type"`
	ConversationId string `json:"conversation_id"`
	IsPc           int    `json:"is_pc"`
	LocalId        string `json:"local_id"`
	Receiver       string `json:"receiver"`
	SendTime       string `json:"send_time"`
	Sender         string `json:"sender"`
	SenderName     string `json:"sender_name"`
	ServerId       string `json:"server_id"`
}

// Link 链接
type Link struct {
	Cdn struct {
		AesKey  string `json:"aes_key"`
		AuthKey string `json:"auth_key"`
		Size    int    `json:"size"`
		Url     string `json:"url"`
	} `json:"cdn"`
	CdnType        int    `json:"cdn_type"`
	ContentType    int    `json:"content_type"`
	ConversationId string `json:"conversation_id"`
	Desc           string `json:"desc"`
	ImageUrl       string `json:"image_url"`
	IsPc           int    `json:"is_pc"`
	LocalId        string `json:"local_id"`
	Receiver       string `json:"receiver"`
	SendTime       string `json:"send_time"`
	Sender         string `json:"sender"`
	SenderName     string `json:"sender_name"`
	ServerId       string `json:"server_id"`
	Title          string `json:"title"`
	Url            string `json:"url"`
}

// MiniProgram 小程序
type MiniProgram struct {
	Appicon string `json:"appicon"`
	Appid   string `json:"appid"`
	Appname string `json:"appname"`
	Cdn     struct {
		AesKey   string `json:"aes_key"`
		AuthKey  string `json:"auth_key"`
		FileId   string `json:"file_id"`
		FileName string `json:"file_name"`
		Height   int    `json:"height"`
		LdSize   int    `json:"ld_size"`
		LdUrl    string `json:"ld_url"`
		Md5      string `json:"md5"`
		MdSize   int    `json:"md_size"`
		MdUrl    string `json:"md_url"`
		Size     int    `json:"size"`
		Url      string `json:"url"`
		Width    int    `json:"width"`
	} `json:"cdn"`
	CdnType        int    `json:"cdn_type"`
	ContentType    int    `json:"content_type"`
	ConversationId string `json:"conversation_id"`
	IsPc           int    `json:"is_pc"`
	LocalId        string `json:"local_id"`
	PagePath       string `json:"page_path"`
	Receiver       string `json:"receiver"`
	SendTime       string `json:"send_time"`
	Sender         string `json:"sender"`
	SenderName     string `json:"sender_name"`
	ServerId       string `json:"server_id"`
	ThumbHeight    int    `json:"thumb_height"`
	ThumbWidth     int    `json:"thumb_width"`
	Title          string `json:"title"`
	Username       string `json:"username"`
}

// FreeVideo 视频号
type FreeVideo struct {
	Appinfo        string `json:"appinfo"`
	Avatar         string `json:"avatar"`
	ContentType    int    `json:"content_type"`
	ConversationId string `json:"conversation_id"`
	CoverUrl       string `json:"cover_url"`
	Desc           string `json:"desc"`
	Extras         string `json:"extras"`
	FeedType       int    `json:"feed_type"`
	IsPc           int    `json:"is_pc"`
	Nickname       string `json:"nickname"`
	ObjectId       string `json:"object_id"`
	ObjectNonceId  string `json:"object_nonce_id"`
	Receiver       string `json:"receiver"`
	SendTime       string `json:"send_time"`
	Sender         string `json:"sender"`
	SenderName     string `json:"sender_name"`
	ServerId       string `json:"server_id"`
	ThumbUrl       string `json:"thumb_url"`
	Url            string `json:"url"`
}

// Sticker 表情包
type Sticker struct {
	ContentType    int    `json:"content_type"`
	ConversationId string `json:"conversation_id"`
	Height         int    `json:"height"`
	IsPc           int    `json:"is_pc"`
	LocalId        string `json:"local_id"`
	Name           string `json:"name"`
	Receiver       string `json:"receiver"`
	SendTime       string `json:"send_time"`
	Sender         string `json:"sender"`
	SenderName     string `json:"sender_name"`
	ServerId       string `json:"server_id"`
	SourceType     int    `json:"source_type"`
	Type           int    `json:"type"`
	Url            string `json:"url"`
	Width          int    `json:"width"`
}

// BizCard 名片
type BizCard struct {
	Appinfo        string `json:"appinfo"`
	Avatar         string `json:"avatar"`
	CdnData        string `json:"cdn_data"`
	ContentType    int    `json:"content_type"`
	ConversationId string `json:"conversation_id"`
	CorpId         int    `json:"corp_id"`
	CorpName       string `json:"corp_name"`
	IsPc           int    `json:"is_pc"`
	LocalId        string `json:"local_id"`
	Nickname       string `json:"nickname"`
	Receiver       string `json:"receiver"`
	SendTime       string `json:"send_time"`
	Sender         string `json:"sender"`
	SenderName     string `json:"sender_name"`
	ServerId       string `json:"server_id"`
	Source         string `json:"source"`
	UserId         string `json:"user_id"`
}

// RedPack 红包
type RedPack struct {
	ContentType    int    `json:"content_type"`
	ConversationId string `json:"conversation_id"`
	Desc           string `json:"desc"`
	LocalId        string `json:"local_id"`
	Money          int    `json:"money"`
	PacketId       string `json:"packet_id"`
	Receiver       string `json:"receiver"`
	Remark         string `json:"remark"`
	SendTime       string `json:"send_time"`
	Sender         string `json:"sender"`
	SenderName     string `json:"sender_name"`
	ServerId       string `json:"server_id"`
}

// MsgCollection 聊天合集
type MsgCollection struct {
	Appinfo        string `json:"appinfo"`
	CdnData        string `json:"cdn_data"`
	ContentType    int    `json:"content_type"`
	ConversationId string `json:"conversation_id"`
	IsPc           int    `json:"is_pc"`
	LocalId        string `json:"local_id"`
	Receiver       string `json:"receiver"`
	SendTime       string `json:"send_time"`
	Sender         string `json:"sender"`
	SenderName     string `json:"sender_name"`
	ServerId       string `json:"server_id"`
}
