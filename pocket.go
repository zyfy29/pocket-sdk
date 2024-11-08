package pocket

import "github.com/zyfy29/pocket-sdk/api"

func NewAPI(token string) api.API {
	return api.New(token)
}
