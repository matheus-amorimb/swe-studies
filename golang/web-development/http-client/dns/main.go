package main

import "fmt"

func main() {
	domain := "google.com"
	ipAddress, err := GetIPAddress(domain)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ipAddress)
}
