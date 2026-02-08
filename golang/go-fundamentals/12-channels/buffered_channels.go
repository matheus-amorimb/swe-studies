package main

func bufferedChannel(emails []string) chan string {
	ch := make(chan string, len(emails))
	for _, email := range emails {
		ch <- email
	}

	return ch
}
