package api

import (
	"github.com/go-resty/resty/v2"
	"os"
)

type DefaultAPI struct {
	Client *resty.Client
}

func (d *DefaultAPI) setup(token string) {
	d.Client = resty.New().SetHeaders(map[string]string{
		"token": token,
	})
}

func New(token string) API {
	d := DefaultAPI{}
	d.setup(token)
	return &d
}

func NewForTest() API {
	token := os.Getenv("POCKET_TOKEN")
	if len(token) == 0 {
		panic("no token found. set it in $POCKET_TOKEN")
	}
	return New(token)
}
