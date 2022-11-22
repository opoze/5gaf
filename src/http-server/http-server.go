package httpserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Handler[T any] func(data T)

func createHTTPServer(mux *http.ServeMux) *http.Server {
	fmt.Printf("HTTP server listen on %d \n", ServerPort)
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", ServerPort),
		Handler: mux,
	}
}
func RegisterHTTPRoute(uri string, mux *http.ServeMux, handler http.HandlerFunc) {
	mux.HandleFunc(uri, handler)
}

func createHTTPServerRouter() *http.ServeMux {
	return http.NewServeMux()
}

func runHTTPServer(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Error running http server: %s\n", err)
		}
	}
}

func writeBadRequestResponse(w *http.ResponseWriter) {
	http.Error(*w, "Wrong Data Type", http.StatusBadRequest)
}

func HttpJsonPostHandlerFactory[T any](handler Handler[*T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			fmt.Printf("HTTP server %s request received at %s \n", r.Method, r.RequestURI)

			var parsedJsonData T
			parsedJsonErr := json.NewDecoder(r.Body).Decode(&parsedJsonData)

			if parsedJsonErr != nil {
				writeBadRequestResponse(&w)
			}

			handler(&parsedJsonData)
		}

	}
}
