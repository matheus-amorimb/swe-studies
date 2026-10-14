package main

import "fmt"

func main() {
	conns := []Connection{
		{ip: "203.0.113.10", username: "mike"},
		{ip: "298.51.100.25", username: "bob"},
		{ip: "292.0.2.5", username: "mike"},
		{ip: "203.0.113.15", username: "bob"},
		{ip: "203.0.113.11", username: "mike"},
	}
	ip := accountSharing(conns)
	fmt.Println("ip:", ip)
}

//username is being shared if it appears in two connections
//if a username is being shared, return an IP of any of them. otherwhise, return an empty string

type Connection struct {
	ip       string
	username string
}

func accountSharing(connections []Connection) string {
	userConns := map[string]string{}
	for _, conn := range connections {
		user := conn.username
		ip := conn.ip
		if _, ok := userConns[user]; ok {
			return ip
		}
		userConns[user] = ip
	}

	return ""
}
