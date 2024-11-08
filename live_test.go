package pocket

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultAPI_LiveOne(t *testing.T) {
	api := newForTest()
	liveId := "1031720783931314176"
	res, err := api.LiveOne(liveId)
	if assert.Nil(t, err) {
		assert.Equal(t, res.LiveId, liveId)
	}

	// todo: test 500 case
}

func TestDefaultAPI_LiveList(t *testing.T) {
	type args struct {
		ownerId  string
		nextTime int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"success",
			args{
				ownerId:  "63566",
				nextTime: 0,
			},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := newForTest()
			got, got1, err := d.LiveList(tt.args.ownerId, tt.args.nextTime)
			if !tt.wantErr(t, err, fmt.Sprintf("LiveList(%v, %v)", tt.args.ownerId, tt.args.nextTime)) {
				return
			}
			t.Log(got)
			t.Log(got1)
			assert.Greater(t, len(*got), 0)
		})
	}
}
