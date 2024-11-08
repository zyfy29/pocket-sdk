package api

import (
	"github.com/zyfy29/pocket-sdk/dto"
)

const (
	liveOneUrl = "https://pocketapi.48.cn/live/api/v1/live/getLiveOne"
)

func (d *DefaultAPI) LiveOne(liveId string) (*dto.Live, error) {
	req := struct {
		LiveId string `json:"liveId"`
	}{liveId}
	resp, err := d.Client.R().SetBody(req).SetResult(dto.Resp[dto.Live]{}).Post(liveOneUrl)
	if err != nil {
		return nil, err
	}
	res := resp.Result().(*dto.Resp[dto.Live])
	if !res.Success {
		return nil, res.ErrorFailed()
	}
	return &res.Content, nil
}
