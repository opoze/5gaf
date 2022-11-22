package httpserver

func Start() {
	httpRouter := createHTTPServerRouter()
	RegisterRoutes(httpRouter)
	httpServer := createHTTPServer(httpRouter)
	runHTTPServer(httpServer)
}
