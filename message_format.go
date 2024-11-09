package pocket

import (
	"encoding/json"
	"fmt"
)

func (mi MessageItem) GetCustomType() string {
	switch mi.MsgType {
	case MessageTypeText:
		return CustomTypeText
	case MessageTypeReply, MessageTypeGiftReply, MessageTypeFlipCard:
		return CustomTypeReply
	case MessageTypeLivePush, MessageTypeImage, MessageTypeExpressImage, MessageTypeAudio, MessageTypeVideo:
		return CustomTypeMedea
	}
	return ""
}

func (mi MessageItem) FormatToTextType() TextMessageDetail {
	return TextMessageDetail{
		MessageItem: mi,
		TextContent: mi.Bodys,
	}
}

func (mi MessageItem) FormarToReplyType() ReplyMessageDetail {
	var ret ReplyMessageDetail
	ret.MessageItem = mi
	switch ret.MsgType {
	case MessageTypeReply:
		var body replyBody
		_ = json.Unmarshal([]byte(mi.Bodys), &body)
		ret.ReplyTo = body.ReplyInfo.ReplyName
		ret.ReferenceContent = body.ReplyInfo.ReplyText
		ret.TextContent = body.ReplyInfo.Text
	case MessageTypeGiftReply:
		var body giftReplyBody
		_ = json.Unmarshal([]byte(mi.Bodys), &body)
		ret.ReplyTo = body.GiftReplyInfo.ReplyName
		ret.ReferenceContent = body.GiftReplyInfo.ReplyText
		ret.TextContent = body.GiftReplyInfo.Text
	case MessageTypeFlipCard:
		var body filpCardBody
		_ = json.Unmarshal([]byte(mi.Bodys), &body)
		ret.ReplyTo = "（文字翻牌）"
		ret.ReferenceContent = body.FilpCardInfo.Question
		ret.TextContent = body.FilpCardInfo.Answer
	default:
		// unimplemented
		ret.TextContent = mi.Bodys
	}
	return ret
}

func (mi MessageItem) FormatToMediaType() MediaMessageDetail {
	var ret MediaMessageDetail
	ret.MessageItem = mi
	switch ret.MsgType {
	case MessageTypeLivePush:
		var body livePushBody
		_ = json.Unmarshal([]byte(mi.Bodys), &body)
		ret.TextContent = fmt.Sprintf("%s (%s)", body.LivePushInfo.LiveTitle, body.LivePushInfo.LiveId)
		ret.MediaUrl = "https://source.48.cn" + body.LivePushInfo.LiveCover
	case MessageTypeImage:
		var body imageBody
		_ = json.Unmarshal([]byte(mi.Bodys), &body)
		ret.MediaUrl = body.Url
		ret.MediaExt = body.Ext
	case MessageTypeExpressImage:
		var body expImageBody
		_ = json.Unmarshal([]byte(mi.Bodys), &body)
		ret.MediaUrl = body.ExpressImgInfo.EmotionRemote
	case MessageTypeAudio:
		var body audioBody
		_ = json.Unmarshal([]byte(mi.Bodys), &body)
		ret.MediaUrl = body.Url
		ret.MediaExt = body.Ext
	case MessageTypeVideo:
		var body videoBody
		_ = json.Unmarshal([]byte(mi.Bodys), &body)
		ret.MediaUrl = body.Url
		ret.MediaExt = body.Ext
	}
	return ret
}
