package httpclient

type NEFPostRequest struct {
	Id   string `json:"id"`
	Addr string `json:"addr"`
}

func PostNEFSubscription() {
	const notifyUrl = "http://localhost:3333/notify"
	const id = "54321"
	const NEFSubscriptionURI = "/subscriptions"

	body := NEFPostRequest{id, notifyUrl}

	Post(CoreUrl+NEFSubscriptionURI, body)

}
