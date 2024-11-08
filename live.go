package pocket

const (
	liveOneUrl  = "https://pocketapi.48.cn/live/api/v1/live/getLiveOne"
	liveListUrl = "https://pocketapi.48.cn/im/api/v1/chatroom/msg/list/aim/type"
)

func (d *DefaultAPI) LiveOne(liveId string) (*Live, error) {
	req := struct {
		LiveId string `json:"liveId"`
	}{liveId}
	resp, err := d.Client.R().SetBody(req).SetResult(Resp[Live]{}).Post(liveOneUrl)
	if err != nil {
		return nil, err
	}
	res := resp.Result().(*Resp[Live])
	if !res.Success {
		return nil, res.ErrorFailed()
	}
	return &res.Content, nil
}

func (d *DefaultAPI) LiveList(ownerId string, nextTime int64) (*[]LiveItem, int64, error) {
	req := struct {
		ExtMsgType string `json:"extMsgType"` // "USER_LIVE"
		RoomId     string `json:"roomId"`     // ""
		OwnerId    string `json:"ownerId"`
		NextTime   int64  `json:"nextTime"` // 0
	}{
		ExtMsgType: "USER_LIVE",
		RoomId:     "0",
		OwnerId:    ownerId,
		NextTime:   nextTime,
	}
	resp, err := d.Client.R().SetBody(req).SetResult(Resp[liveList]{}).Post(liveListUrl)
	if err != nil {
		return nil, 0, err
	}
	res := resp.Result().(*Resp[liveList])
	if !res.Success {
		return nil, 0, res.ErrorFailed()
	}

	var ret []LiveItem
	for _, it := range res.Content.Message {
		ret = append(ret, it.FormatToLiveItem())
	}
	return &ret, 0, nil
}
