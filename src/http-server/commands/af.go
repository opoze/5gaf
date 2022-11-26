package commands

import "fmt"

type NotifyData struct {
	Data string `json:"data"`
}

// Use the data
func AF(data *NotifyData) {
	fmt.Printf("Data received in AF: %s \n", data.Data)
}
