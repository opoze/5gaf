package main

import (
	httpclient "github.com/opoze/5gaf/http-client"
	httpserver "github.com/opoze/5gaf/http-server"
)

func main() {

	// Post NEF subscription
	httpclient.PostNEFSubscription()

	// Start HTTP server responding to POST on /notify
	httpserver.Start()
}
