package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post[T any](url string, body T) {

	fmt.Println(url)
	fmt.Println(body)

	req := createHTTPJsonPostRequest(url, body)

	httpClient := &http.Client{}

	resp, err := httpClient.Do(req)

	if err != nil {
		fmt.Printf("http-client: error on POST %s - %s\n", url, err)
	} else {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("http-client: could not read response body: %s\n", err)
		} else {
			fmt.Printf("client: response body: %s\n", respBody)
		}
	}
}

func createHTTPJsonPostRequest[T any](url string, body T) *http.Request {
	var buff bytes.Buffer
	err := json.NewEncoder(&buff).Encode(body)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", url, &buff)

	if err != nil {
		fmt.Println(err)
	}

	return req
}
