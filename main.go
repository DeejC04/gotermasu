package main

import (
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	// "flag"
)

func main() {
	// will be used to make requests with headers (doesn't seem like
	// default http.Get() command can use headers
	client := &http.Client{}

	// creating the request itself, afaik it doesn't actually send it yet
	// since I'm literally just creating a new request
	req, err := http.NewRequest("GET", "https://eadvs-cscc-catalog-api.apps.asu.edu/catalog-microservices/api/v1/search/subjects?sl=Y&term=2247", nil)
	if err != nil {
		log.Fatal(err)
	}

	// adding the header necessary to send the request
	req.Header.Add("Authorization", "Bearer null")

	// finally, sending the actual request via the client with the headers included
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// by default, response is a byte slice (research it), so it must be converted
	// to a string
	fmt.Println(string(body))

}