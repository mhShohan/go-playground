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

	// Send a POST request with JSON data
	data := PostJsonRequest(URL)
	fmt.Println("Response from PostJsonRequest:", data)

	// Send a POST request with JSON using a custom HTTP client
	PostJsonRequestFromAI(URL)

	// Send a POST request with form data
	formData := PostFormData(URL)
	fmt.Println("Response from PostFormData:", formData)
}

// checkNilError checks if the error is not nil and panics
func checkNilError(err error) {
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
	// Fake JSON data
	reqBody := strings.NewReader(`{
		"title": "title 1",
		"content": "this is the content of the post"
	}`)

	response, err := http.Post(uri, "application/json", reqBody)
	checkNilError(err)
	defer response.Body.Close()

	buffer, err := io.ReadAll(response.Body)
	checkNilError(err)

	return string(buffer)
}

/*
-----------------------------------
Post JSON data - AI generated code
-----------------------------------
*/
func PostJsonRequestFromAI(uri string) {
	reqBody := strings.NewReader(`{
		"title": "ChatGPT Code - Copilot",
		"content": "This is the content of the post from AI-generated code"
	}`)

	// Create a new HTTP POST request
	req, err := http.NewRequest("POST", uri, reqBody)
	checkNilError(err)

	// Add necessary headers
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client
	client := &http.Client{}

	// Send the request via the client
	resp, err := client.Do(req)
	checkNilError(err)

	// Close the response body
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	checkNilError(err)

	// Print the response body
	fmt.Println("Response from PostJsonRequestFromAI:", string(body))
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
	checkNilError(err)
	defer response.Body.Close()

	buffer, err := io.ReadAll(response.Body)
	checkNilError(err)

	return string(buffer)
}

/*
Notes:
---------
1. **Making HTTP POST Requests**:
   - `http.Post(url, contentType, body)`: A simple way to send a JSON POST request.
   - `http.NewRequest("POST", url, body)`: More flexible method to send a POST request (allows headers and custom client).
   - `http.PostForm(url, data)`: Used for sending form-encoded data.

2. **Headers and Clients**:
   - Headers (like `Content-Type`) should be explicitly set when using `http.NewRequest`.
   - Using `http.Client{}` allows for more control over the request.

3. **Error Handling**:
   - Always check for `response.StatusCode` when handling responses.
   - Ensure `defer response.Body.Close()` is used to release resources.

4. **Choosing Between Methods**:
   - `http.Post()`: Simple and quick for JSON requests.
   - `http.NewRequest() + http.Client{}`: Required when setting headers or customizing the request.
   - `http.PostForm()`: Best for traditional HTML form submissions.

Key Takeaways:
-----------------
✔ Always close response bodies using `defer`.
✔ Use `http.NewRequest()` for more flexibility with headers.
✔ Choose the right method based on the type of data being sent.
*/
