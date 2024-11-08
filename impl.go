package pocket

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

func newWithToken(token string) API {
	d := DefaultAPI{}
	d.setup(token)
	return &d
}

func newForTest() API {
	token := os.Getenv("POCKET_TOKEN")
	if len(token) == 0 {
		panic("no token found. set it in $POCKET_TOKEN")
	}
	return newWithToken(token)
}
