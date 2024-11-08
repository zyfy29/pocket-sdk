package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultAPI_LiveOne(t *testing.T) {
	api := NewForTest()
	liveId := "1031720783931314176"
	res, err := api.LiveOne(liveId)
	if assert.Nil(t, err) {
		assert.Equal(t, res.LiveId, liveId)
	}
}

// todo: test 500 case
