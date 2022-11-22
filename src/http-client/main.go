package httpclient

type NEFPostRequest struct {
	Id   string `json:"id"`
	Addr string `json:"addr"`
}

func PostNEFSubscription() {
	notifyUrl := CoreUrl + "/subscriptions"
	body := NEFPostRequest{NEFid, notifyUrl}

	Post(notifyUrl, body)

}
