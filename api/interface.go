package api

import "github.com/zyfy29/pocket-sdk/dto"

type API interface {
	LiveAPI
}

type LiveAPI interface {
	LiveOne(liveId string) (*dto.Live, error)
}
