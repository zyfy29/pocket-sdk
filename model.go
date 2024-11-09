package pocket

import (
	"encoding/json"
	"fmt"
	"strconv"
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

type liveListItem struct {
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

func (l liveListItem) FormatToLiveItem() LiveItem {
	var extInfo liveExtInfo
	_ = json.Unmarshal([]byte(l.ExtInfo), &extInfo)
	return LiveItem{
		ID:    strconv.FormatInt(extInfo.ID, 10),
		Title: extInfo.Title,
		Time:  time.UnixMilli(l.MsgTime),
	}
}

type LiveItem struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	Time  time.Time `json:"liveTime"`
}

type liveList struct {
	Message  []liveListItem `json:"message"`
	NextTime int64          `json:"nextTime"`
}

type VoiceOperateContent struct {
	VoiceUserList []VoiceUser `json:"voiceUserList"`
	StreamUrl     string      `json:"streamUrl"`
}

type VoiceUser struct {
	UserId      int    `json:"userId"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	PfUrl       string `json:"pfUrl"`
	VoiceStatus bool   `json:"voiceStatus"`
}

type VoiceStatus struct {
	VoiceUser
	StreamUrl string
}
