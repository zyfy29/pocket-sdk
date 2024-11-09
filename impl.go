package pocket

const (
	liveOneUrl  = "https://pocketapi.48.cn/live/api/v1/live/getLiveOne"
	liveListUrl = "https://pocketapi.48.cn/im/api/v1/chatroom/msg/list/aim/type"
	voiceUrl    = "https://pocketapi.48.cn/im/api/v1/team/voice/operate"
)

func (d *DefaultAPI) LiveOne(liveId string) (Live, error) {
	req := struct {
		LiveId string `json:"liveId"`
	}{liveId}
	resp, err := d.Client.R().SetBody(req).SetResult(Resp[Live]{}).Post(liveOneUrl)
	if err != nil {
		return Live{}, err
	}
	res := resp.Result().(*Resp[Live])
	if !res.Success {
		return Live{}, res.ErrorFailed()
	}
	return res.Content, nil
}

func (d *DefaultAPI) LiveList(ownerId string, nextTime int64) ([]LiveItem, int64, error) {
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
	return ret, res.Content.NextTime, nil
}

func (d *DefaultAPI) Voice(serverId string, channelId string) (VoiceStatus, error) {
	req := struct {
		ChannelId   string `json:"channelId"`
		ServerId    string `json:"serverId"`
		OperateCode int    `json:"operateCode"` // 2
	}{
		ChannelId:   channelId,
		ServerId:    serverId,
		OperateCode: 2,
	}
	resp, err := d.Client.R().SetBody(req).SetResult(Resp[voiceOperate]{}).Post(voiceUrl)
	if err != nil {
		return VoiceStatus{}, err
	}
	res := resp.Result().(*Resp[voiceOperate])
	if !res.Success {
		return VoiceStatus{}, res.ErrorFailed()
	}

	if len(res.Content.VoiceUserList) == 1 {
		return VoiceStatus{res.Content.VoiceUserList[0], res.Content.StreamUrl}, nil
	}
	return VoiceStatus{voiceUser{}, res.Content.StreamUrl}, nil
}
