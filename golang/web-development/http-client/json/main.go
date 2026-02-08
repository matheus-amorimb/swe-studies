package main

import "fmt"

func main() {
	//Decoder
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/issues"
	issues, err := GetIssues(url)
	if err != nil {
		fmt.Println("error getting issues: %w", err)
	}

	//Marshal JSON
	issuesInBytes, err := MarshAll(issues)
	if err != nil {
		fmt.Println("error getting issues: %w", err)
	}

	for _, issueInByte := range issuesInBytes {
		fmt.Println(string(issueInByte))
	}

	//map[string]interface{}
	resources, err := GetResourcesAsMapStringAny(url)
	if err != nil {
		fmt.Println("error unmarshing to map[string]any: %w", err)
	}
	LogResources(resources)
}
