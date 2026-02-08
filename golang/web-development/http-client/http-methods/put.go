package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func httpPut(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := getFullUrl(baseURL, id)

	userData, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(userData))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	httpClient := &http.Client{}

	res, err := httpClient.Do(req)
	if err != nil {
		return User{}, err
	}

	var user User
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&user); err != nil {
		return User{}, err

	}

	return user, nil
}

func getFullUrl(baseURL, id string) string {
	fullURL := baseURL + "/" + id

	return fullURL
}

type User struct {
	Role       string `json:"role"`
	ID         string `json:"id"`
	Experience int    `json:"experience"`
	Remote     bool   `json:"remote"`
	User       struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Age      int    `json:"age"`
	} `json:"user"`
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := getFullUrl(baseURL, id)

	//new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}

	//create client
	httpClient := &http.Client{}

	//make request
	res, err := httpClient.Do(req)
	if err != nil {
		return User{}, err
	}

	var user User
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&user); err != nil {
		return User{}, err
	}

	return user, nil

}
