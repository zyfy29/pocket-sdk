package pocket

import (
	"fmt"
	"time"
)

type Resp[T any] struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Content T      `json:"content"`
}

func (r Resp[T]) ErrorFailed() error {
	return fmt.Errorf("pocket returns %d, due to: %s", r.Status, r.Message)
}

// 结构体命名规则：
// XxBase表示借口返回的原始内容，XxDetail表示将所有扩展信息反序列化后得到的内容，XxItem表示将一部分扩展信息反序列化后得到的内容

// Live 直播详情
// deprecated，即将重命名为LiveDetail
type Live struct {
	LiveId         string `json:"liveId"`
	RoomId         string `json:"roomId"`
	OnlineNum      int    `json:"onlineNum"`
	Type           int    `json:"type"`
	LiveType       int    `json:"liveType"`
	Review         bool   `json:"review"`
	NeedForward    bool   `json:"needForward"`
	SystemMsg      string `json:"systemMsg"`
	MsgFilePath    string `json:"msgFilePath"`
	PlayStreamPath string `json:"playStreamPath"`
	User           struct {
		UserId     string `json:"userId"`
		UserName   string `json:"userName"`
		UserAvatar string `json:"userAvatar"`
		Level      int    `json:"level"`
	} `json:"user"`
	TopUser            []interface{} `json:"topUser"`
	Mute               bool          `json:"mute"`
	LiveMode           int           `json:"liveMode"`
	PictureOrientation int           `json:"pictureOrientation"`
	IsCollection       int           `json:"isCollection"`
	MergeStreamUrl     string        `json:"mergeStreamUrl"`
	Crm                string        `json:"crm"`
	CoverPath          string        `json:"coverPath"`
	Title              string        `json:"title"`
	Ctime              string        `json:"ctime"`
	Announcement       string        `json:"announcement"`
	SpecialBadge       []interface{} `json:"specialBadge"`
}

type LiveDetail = Live

type liveList struct {
	Message  []liveListBase `json:"message"`
	NextTime int64          `json:"nextTime"`
}

type liveListBase struct {
	MsgidClient string `json:"msgidClient"`
	MsgTime     int64  `json:"msgTime"`
	MsgType     string `json:"msgType"`
	Bodys       string `json:"bodys"`
	ExtInfo     string `json:"extInfo"`
	Privacy     bool   `json:"privacy"`
}

type liveExtInfo struct {
	ID             int64    `json:"id"`
	CoverUrl       string   `json:"coverUrl"`
	CoverUrlList   []string `json:"coverUrlList"`
	Title          string   `json:"title"`
	Content        string   `json:"content"`
	Url            string   `json:"url"`
	JumpType       string   `json:"jumpType"`
	JumpPath       string   `json:"jumpPath"`
	ThirdAppName   string   `json:"thirdAppName"`
	ThirdAPPImgUrl string   `json:"thirdAPPImgUrl"`
	MessageType    string   `json:"messageType"`
	User           struct {
		UserId   int    `json:"userId"`
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
	} `json:"user"`
}

// LiveItem 直播列表显示
type LiveItem struct {
	liveListBase
	liveExtInfo
	Time time.Time
}

type voiceOperate struct {
	VoiceUserList []voiceUser `json:"voiceUserList"`
	StreamUrl     string      `json:"streamUrl"`
}

type voiceUser struct {
	UserId      int    `json:"userId"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	PfUrl       string `json:"pfUrl"`
	VoiceStatus bool   `json:"voiceStatus"`
}

// VoiceStatus 上麦状态
type VoiceStatus struct {
	voiceUser
	StreamUrl string
}

const (
	// text only
	MessageTypeText = "TEXT"

	// text reply
	MessageTypeReply     = "REPLY"
	MessageTypeGiftReply = "GIFTREPLY"
	MessageTypeFlipCard  = "FLIPCARD" // 文字翻牌（公开）

	// media
	MessageTypeLivePush     = "LIVEPUSH"
	MessageTypeImage        = "IMAGE"
	MessageTypeExpressImage = "EXPRESSIMAGE"
	MessageTypeAudio        = "AUDIO"
	MessageTypeVideo        = "VIDEO"
	// todo 语音翻牌，视频翻牌

	// custom types
	CustomTypeText  = "POCKET_TEXT"
	CustomTypeReply = "POCKET_REPLY"
	CustomTypeMedea = "POCKET_MEDIA"
)

type RoomMessageContent struct {
	Messages []messageBase `json:"message"`
	NextTime int64         `json:"nextTime"`
}

type messageBase struct {
	MsgIDServer string `json:"msgIdServer"`
	MsgIDClient string `json:"msgIdClient"`
	MsgTime     int64  `json:"msgTime"`
	MsgType     string `json:"msgType"`
	Bodys       string `json:"bodys"`
	ExtInfo     string `json:"extInfo"`
}

// for messageBase.ExtInfo
type messageExtInfo struct {
	Module      string `json:"module"`
	ChannelRole string `json:"channelRole"`
	User        struct {
		UserId   int    `json:"userId"`
		NickName string `json:"nickName"`
		TeamLogo string `json:"teamLogo"`
		Avatar   string `json:"avatar"`
		Level    int    `json:"level"`
		RoleId   int    `json:"roleId"`
		Vip      bool   `json:"vip"`
		PfUrl    string `json:"pfUrl"`
	} `json:"user"`
	BubbleId string `json:"bubbleId"`
}

// MessageItem 只翻译了ext部分
type MessageItem struct {
	messageBase
	messageExtInfo
}

func (mi MessageItem) GetCustomType() string {
	switch mi.MsgType {
	case MessageTypeText:
		return CustomTypeText
	case MessageTypeReply, MessageTypeGiftReply, MessageTypeFlipCard:
		return CustomTypeReply
	case MessageTypeLivePush, MessageTypeImage, MessageTypeExpressImage, MessageTypeAudio, MessageTypeVideo:
		return CustomTypeMedea
	}
	return ""
}

// for messageBase.Bodys

type replyBody struct {
	ReplyInfo struct {
		ReplyText      string `json:"replyText"`
		ReplyName      string `json:"replyName"`
		ReplyMessageId string `json:"replyMessageId"`
		Text           string `json:"text"`
	} `json:"replyInfo"`
	MessageType string `json:"messageType"`
}

type giftReplyBody struct {
	GiftReplyInfo struct {
		ReplyText      string `json:"replyText"`
		ReplyName      string `json:"replyName"`
		ReplyMessageId string `json:"replyMessageId"`
		Text           string `json:"text"`
	} `json:"giftReplyInfo"`
	MessageType string `json:"messageType"`
}

type livePushBody struct {
	LivePushInfo struct {
		LiveCover string `json:"liveCover"`
		LiveTitle string `json:"liveTitle"`
		LiveId    string `json:"liveId"`
		ShortPath string `json:"shortPath"`
	} `json:"livePushInfo"`
	MessageType string `json:"messageType"`
}

type imageBody struct {
	Size int    `json:"size"`
	Ext  string `json:"ext"`
	W    int    `json:"w"`
	Url  string `json:"url"`
	Md5  string `json:"md5"`
	H    int    `json:"h"`
}

type expImageBody struct {
	ExpressImgInfo struct {
		EmotionRemote string `json:"emotionRemote"`
		Width         int    `json:"width"`
		Height        int    `json:"height"`
	} `json:"expressImgInfo"`
	MessageType string `json:"messageType"`
}

type audioBody struct {
	Size int    `json:"size"`
	Ext  string `json:"ext"`
	Dur  int    `json:"dur"`
	Url  string `json:"url"`
	Md5  string `json:"md5"`
}

type videoBody struct {
	Url  string `json:"url"`
	Md5  string `json:"md5"`
	Ext  string `json:"ext"`
	H    int    `json:"h"`
	Size int    `json:"size"`
	W    int    `json:"w"`
	Dur  int    `json:"dur"`
}

type filpCardBody struct {
	FilpCardInfo struct {
		Question   string `json:"question"`
		Answer     string `json:"answer"`
		QuestionId string `json:"questionId"`
		AnswerId   string `json:"answerId"`
		AnswerType string `json:"answerType"`
	} `json:"filpCardInfo"`
	MessageType string `json:"messageType"`
}

// custom types

type TextMessageDetail struct {
	MessageItem
	TextContent string
}

type ReplyMessageDetail struct {
	MessageItem
	ReplyTo          string
	ReferenceContent string
}

type MediaMessageDetail struct {
	MessageItem
	MediaUrl string
}
