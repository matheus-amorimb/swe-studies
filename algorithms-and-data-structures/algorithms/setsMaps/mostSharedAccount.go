package main

func main() {
	conns := []Connection{
		{ip: "203.0.113.10", username: "mike"},
		{ip: "298.51.100.25", username: "bob"},
		{ip: "292.0.2.5", username: "mike"},
		{ip: "203.0.113.15", username: "bob"},
		{ip: "203.0.113.11", username: "bob"},
	}
	mostSharedAccount(conns)
	// fmt.Println("user:", u)
}

type Connection struct {
	ip       string
	username string
}

func mostSharedAccount(connections []Connection) string {
	userConns := map[string]int{}
	for _, conn := range connections {
		user := conn.username
		if _, ok := userConns[user]; ok {
			userConns[user]++
		} else {
			userConns[user] = 1
		}
	}

	mostSharedUser := ""
	for user, nConns := range userConns {
		if mostSharedUser == "" || nConns > userConns[mostSharedUser] {
			mostSharedUser = user
		}
	}

	return mostSharedUser
}
