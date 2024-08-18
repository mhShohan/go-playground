package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const URL = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	// Start the web request - GET
	data := GetRequest(URL)
	fmt.Println(data)

	// Another way to write the GetRequest function
	dataAnother := GetRequestAnother(URL)
	fmt.Println(dataAnother)
}

// checkNillError checks if the error is not nil and panics
func checkNillError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetRequest(uri string) string {
	response, err := http.Get(uri)
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

// Another way to write the GetRequest function
func GetRequestAnother(uri string) string {
	response, err := http.Get(uri)
	checkNillError(err) // Check if there is an error

	defer response.Body.Close()

	// Check if the status code is not 200
	if response.StatusCode != 200 {
		fmt.Println("Status:", response.Status)
		panic("Failed to fetch data")
	}

	var responseString strings.Builder

	buffer, err := io.ReadAll(response.Body)
	checkNillError(err) // Check if there is an error

	byteCount, err := responseString.Write(buffer)
	checkNillError(err) // Check if there is an error

	fmt.Println("Bytes written:", byteCount)
	return responseString.String()
}
