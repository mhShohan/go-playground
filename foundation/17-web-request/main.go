package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/postsa/1"

func main() {

	// Fetch data from the url
	data := fetchData(url)
	fmt.Println(data)

}

// checkNillError checks if the error is not nil and panics
func checkNillError(err error) {
	if err != nil {
		panic(err)
	}
}

func fetchData(url string) string {
	response, err := http.Get(url)

	checkNillError(err) // Check if there is an error

	defer response.Body.Close()

	// Check if the status code is not 200
	if response.StatusCode != 200 {
		fmt.Println("Status:", response.Status)
		panic("Failed to fetch data")
	}

	buffer, err := io.ReadAll(response.Body)

	checkNillError(err) // Check if there is an error

	return string(buffer)
}
