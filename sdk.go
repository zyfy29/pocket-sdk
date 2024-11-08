package pocket

func NewAPI(token string) API {
	return newWithToken(token)
}
