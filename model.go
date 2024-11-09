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
