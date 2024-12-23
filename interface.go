package pocket

type API interface {
	LiveOne(liveId string) (Live, error)
	LiveList(ownerId string, nextTime int64) ([]LiveItem, int64, error)
	Voice(serverId string, channelId string) (VoiceStatus, error)
	Message(serverId, channelId string, nextTime int64) ([]MessageItem, int64, error)
}
