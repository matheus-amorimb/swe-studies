package main

import (
	"fmt"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://api.example.com/users", nil)
	if err != nil {
		fmt.Println("error creating request: ", err)
		return
	}

	// setting a header on the new request
	req.Header.Set("x-api-key", "123456789")

	// making the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error making request: ", err)
		return
	}
	defer res.Body.Close()

	// reading a header from the response
	header := res.Header.Get("last-modified")
	fmt.Println("last modified: ", header)

	// deleting a header from the response
	res.Header.Del("last-modified")
}
