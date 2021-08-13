package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	req, _ := http.NewRequest(http.MethodGet, "https://data.alpaca.markets/v1/last/stocks/MSFT", nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", "*****")
	req.Header.Add("APCA-API-SECRET-KEY", "*****")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error sending HTTP request: %v", err)
	}
	responseBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading HTTP response body: %v", err)
	}

	fmt.Println("We got the response:", string(responseBytes))

}
