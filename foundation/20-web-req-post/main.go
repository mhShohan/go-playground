package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const URL = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("-----------Make a POST request-----------")
	// Start the web request - POST

	data := PostJsonRequest(URL)
	fmt.Println(data)

	PostJsonRequestFromAI(URL)

	formData := PostFormData(URL)
	fmt.Println(formData)

}

// checkNillError checks if the error is not nil and panics
func checkNillError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
-----------------------------------
Post JSON data
-----------------------------------
*/
func PostJsonRequest(uri string) string {
	// fake json data
	reqBody := strings.NewReader(`
		{
			"title": "title 1",
			"content": "this is the content of the post"
		}
	`)

	response, err := http.Post(uri, "application/json", reqBody)
	checkNillError(err)
	defer response.Body.Close()

	buffer, err := io.ReadAll(response.Body)
	checkNillError(err)

	return string(buffer)
}

/*
-----------------------------------
Post JSON data - AI generated code
-----------------------------------
*/
func PostJsonRequestFromAI(uri string) {
	reqBody := strings.NewReader(`
		{
			"title": "ChatGPT Code - copilot",
			"content": "this is the content of the post from AI generated code"
		}
	`)

	// Create a new request
	req, err := http.NewRequest("POST", uri, reqBody)
	checkNillError(err)

	// Add necessary headers
	req.Header.Set("Content-Type", "application/json")

	// Create a new client
	client := &http.Client{}

	// Send the request via a client
	resp, err := client.Do(req)
	checkNillError(err)

	// Close the response body
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	checkNillError(err)

	// Print the response body
	fmt.Printf("%s\n", body)
}

/*
-----------------------------------
Post FormData
-----------------------------------
*/
func PostFormData(uri string) string {
	formData := url.Values{}
	formData.Add("title", "Post FormData")
	formData.Add("content", "This is the content of FormData")

	response, err := http.PostForm(uri, formData)
	checkNillError(err)

	defer response.Body.Close()

	buffer, err := io.ReadAll(response.Body)
	checkNillError(err)

	return string(buffer)
}
