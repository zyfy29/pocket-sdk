package pocket

type API interface {
	LiveAPI
}

type LiveAPI interface {
	LiveOne(liveId string) (*Live, error)
	LiveList(ownerId string, nextTime int64) (*[]LiveItem, int64, error)
}
