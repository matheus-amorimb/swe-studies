package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	users := []Conn{
		{"mike", []string{"203.03.10", "208.51.0.5", "52.0.2.5"}},
		{"bob", []string{"111.0.0.10", "222.0.0.5", "222.0.0.8"}},
		{"bob2", []string{"222.0.0.5", "222.0.0.8", "111.0.0.10"}},
	}
	cheated := multiAccountCheating(users)
	fmt.Println("cheaters:", cheated)

}

type Conn struct {
	username string
	ips      []string
}

func multiAccountCheating(conns []Conn) bool {
	uniqueLists := make(map[string]struct{}, len(conns))
	for _, conn := range conns {
		slices.Sort(conn.ips)
		ipsAsStr := strings.Join(conn.ips, "|")
		fmt.Println(ipsAsStr)
		if _, ok := uniqueLists[ipsAsStr]; ok {
			return true
		}
		uniqueLists[ipsAsStr] = struct{}{}
	}

	return false
}
