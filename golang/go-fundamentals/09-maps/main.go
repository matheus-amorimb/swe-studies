package main

import "errors"

func main() {

}

// Maps
func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := map[string]user{}
	isSizeValid := len(names) == len(phoneNumbers)
	if !isSizeValid {
		return userMap, errors.New("invalid sizes")
	}

	for i, name := range names {
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[i],
		}
	}

	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
