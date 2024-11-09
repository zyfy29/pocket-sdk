package pocket

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
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
			name: "success",
			args: args{
				ownerId:  "63566",
				nextTime: 0,
			},
			wantErr: assert.NoError,
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
			assert.Greater(t, len(got), 0)
		})
	}
}

func TestDefaultAPI_Voice(t *testing.T) {
	type args struct {
		serverId  string
		channelId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				"1213755",
				"1360923",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := newForTest()
			got, err := d.Voice(tt.args.serverId, tt.args.channelId)
			if !tt.wantErr(t, err, fmt.Sprintf("Voice(%v, %v)", tt.args.serverId, tt.args.channelId)) {
				return
			}
			t.Log(got)
		})
	}
}

func TestDefaultAPI_Message(t *testing.T) {
	type args struct {
		serverId  string
		channelId string
		nextTime  int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				serverId:  "2197175",
				channelId: "2720929",
				nextTime:  0,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := newForTest()
			got, got1, err := d.Message(tt.args.serverId, tt.args.channelId, tt.args.nextTime)
			if !tt.wantErr(t, err, fmt.Sprintf("Message(%v, %v, %v)", tt.args.serverId, tt.args.channelId, tt.args.nextTime)) {
				return
			}
			t.Log(got)
			t.Log(got1)
			assert.Greater(t, len(got), 0)
			for i, mi := range got {
				customType := mi.GetCustomType()
				t.Log(i, customType, strings.Repeat("-", 15))
				switch customType {
				case CustomTypeText:
					customMsg := mi.FormatToTextType()
					t.Log(customMsg.TextContent)
				case CustomTypeReply:
					customMsg := mi.FormarToReplyType()
					t.Log(customMsg.TextContent)
					t.Log(customMsg.ReplyTo)
					t.Log(customMsg.ReferenceContent)
				case CustomTypeMedea:
					customMsg := mi.FormatToMediaType()
					t.Log(customMsg.TextContent)
					t.Log(customMsg.MediaUrl)
					t.Log(customMsg.MediaExt)
				}
			}
		})
	}
}
