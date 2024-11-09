package pocket

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

// custom types

type TextMessageDetail struct {
	MessageItem
	TextContent string
}

type ReplyMessageDetail struct {
	TextMessageDetail
	ReplyTo          string
	ReferenceContent string
}

type MediaMessageDetail struct {
	TextMessageDetail
	MediaUrl string
	MediaExt string
}
