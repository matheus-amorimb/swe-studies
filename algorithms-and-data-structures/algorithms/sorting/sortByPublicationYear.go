package main

import (
	"fmt"
)

type Book struct {
	title         string
	author        string
	pageCount     string
	genre         string
	yearPublished int
}

func main() {
	books := []Book{
		{title: "Dune", author: "Frank Herbert", pageCount: "412", genre: "Science Fiction", yearPublished: 1965},
		{title: "The Catcher in the Rye", author: "J.D. Salinger", pageCount: "234", genre: "Fiction", yearPublished: 1951},
		{title: "The Hobbit", author: "J.R.R. Tolkien", pageCount: "310", genre: "Fantasy", yearPublished: 1937},
		{title: "Foundation", author: "Isaac Asimov", pageCount: "255", genre: "Science Fiction", yearPublished: 1951},
		{title: "1984", author: "George Orwell", pageCount: "328", genre: "Dystopian", yearPublished: 1949},
	}

	BucketSort(&books)
	fmt.Println(books)
}

func BucketSort(books *[]Book) {
	if len(*books) == 0 {
		return
	}

	minYear, maxYear := (*books)[0].yearPublished, (*books)[0].yearPublished
	for _, book := range *books {
		minYear = min(minYear, book.yearPublished)
		maxYear = max(maxYear, book.yearPublished)
	}

	buckets := make([][]Book, maxYear-minYear+1)
	for _, book := range *books {
		buckets[book.yearPublished-minYear] = append(buckets[book.yearPublished-minYear], book)
	}

	sortedArr := make([]Book, 0, len(*books))
	for _, bucket := range buckets {
		sortedArr = append(sortedArr, bucket...)
	}

	*books = sortedArr
}
