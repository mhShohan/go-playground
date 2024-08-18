package main

import (
	"fmt"
	"net/url"
)

const uri = "https://localhost:5000/posts/1?name=shohan&age=20"

func main() {
	fmt.Println("____________URLs______________")

	result, err := url.Parse(uri)
	checkNillError(err)

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Port:", result.Port())
	fmt.Println("Path:", result.Path)
	fmt.Println("RawQuery:", result.RawQuery)
	fmt.Println("Fragment:", result.Fragment)
	fmt.Println("String:", result.String())
	fmt.Println("Query:", result.Query())

	queryParams := result.Query()

	fmt.Println("____________Query Params________________")
	fmt.Println("Name:", queryParams.Get("name"))

	for key, value := range queryParams {
		fmt.Println(key, ":", value)
	}

	fmt.Println("____________Construct URL_____________")

	constructedURL := &url.URL{
		Scheme:   "https",
		Host:     "localhost:5000",
		Path:     "/posts/1",
		RawQuery: "name=shohan&age=20",
	}

	fmt.Println("Constructed URL:", constructedURL.String())
}

// checkNillError checks if the error is not nil and panics
func checkNillError(err error) {
	if err != nil {
		panic(err)
	}
}
