package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts/1" // Corrected URL

func main() {
	fmt.Println("----- HTTP Request Example -----")

	// Fetch data from the URL
	data, err := fetchData(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response Data:")
	fmt.Println(data)
}

// fetchData sends an HTTP GET request and returns the response body
func fetchData(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch data: %v", err)
	}
	defer response.Body.Close()

	// Check if the status code is not 200
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d %s", response.StatusCode, response.Status)
	}

	buffer, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(buffer), nil
}

/*
Notes:
=========
- **Making HTTP Requests in Go**
  - `http.Get(url)`: Sends an HTTP GET request.
  - `response.Body.Close()`: Ensures the response body is closed after reading.
  - `io.ReadAll(response.Body)`: Reads the response body.

- **Error Handling in HTTP Requests**
  - Always check for network errors when making requests.
  - Validate the HTTP status code (`response.StatusCode`).
  - Use `defer response.Body.Close()` to avoid resource leaks.
  - Instead of panicking, return errors and handle them properly.

Key Takeaways:
- Always handle HTTP errors and unexpected status codes gracefully.
- Use `defer response.Body.Close()` to prevent resource leaks.
- Avoid `panic` for expected errors; return an `error` and handle it.
*/
