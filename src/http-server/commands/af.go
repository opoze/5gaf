package commands

import "fmt"

type NotifyData struct {
	ID uint64 `json:"id"`
}

// Use the data
func AF(data *NotifyData) {
	fmt.Printf("Data received in AF: %d \n", data.ID)
}
