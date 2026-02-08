package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Comment struct {
	Id      string `json:"id"`
	UserId  string `json:"user_id"`
	Comment string `json:"comment"`
}

func httpPost() {
	url := "any-url.com.br"

	commentStruct := Comment{
		Id:      "123",
		UserId:  "user123",
		Comment: "hello world",
	}

	jsonData, err := json.Marshal(commentStruct)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "my-authorization-key")

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	fmt.Println(res)
}
