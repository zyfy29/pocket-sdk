package dto

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
