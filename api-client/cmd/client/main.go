package main

import (
	"api-client/api"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	client := api.NewClient(api.BaseURL, &http.Client{
		Timeout: 10 * time.Second,
	})

	response, err := client.GetMonster()

	if err != nil {
		log.Fatal(err)
	}

	if response.Status == 200 {
		for _, monster := range response.Data {
			fmt.Printf("%+v\n", monster)
		}
	}
}
