package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const URL = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("_________HTTP GET REQUEST_________")

	// Fetch data using first method
	data := GetRequest(URL)
	fmt.Println("Response Data (Method 1):", data)

	// Fetch data using second method
	dataAnother := GetRequestAnother(URL)
	fmt.Println("Response Data (Method 2):", dataAnother)
}

// checkNilError checks if the error is not nil and panics
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

// GetRequest performs a simple HTTP GET request and returns the response body as a string
func GetRequest(uri string) string {
	response, err := http.Get(uri)
	checkNilError(err) // Handle error

	defer response.Body.Close() // Ensure response body is closed after function execution

	// Validate response status
	if response.StatusCode != 200 {
		fmt.Println("Status:", response.Status)
		panic("Failed to fetch data")
	}

	buffer, err := io.ReadAll(response.Body)
	checkNilError(err)

	return string(buffer)
}

// GetRequestAnother performs an HTTP GET request and uses strings.Builder to store the response
func GetRequestAnother(uri string) string {
	response, err := http.Get(uri)
	checkNilError(err)

	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("Status:", response.Status)
		panic("Failed to fetch data")
	}

	var responseString strings.Builder
	buffer, err := io.ReadAll(response.Body)
	checkNilError(err)

	byteCount, err := responseString.Write(buffer)
	checkNilError(err)

	fmt.Println("Bytes written:", byteCount)
	return responseString.String()
}

/*
Notes:
---------
1. **HTTP GET Requests**:
   - `http.Get(url)`: Sends a GET request to the specified URL.
   - Always handle errors using proper error checking.
   - The `defer response.Body.Close()` ensures the response body is closed properly.

2. **Handling Response Data**:
   - `io.ReadAll(response.Body)`: Reads the entire response body.
   - `strings.Builder`: More efficient for handling large string manipulations.

3. **Error Handling**:
   - Always check if `response.StatusCode` is 200 (OK).
   - Use a helper function (`checkNilError()`) to handle errors consistently.

4. **Differences Between Methods**:
   - `GetRequest()`: Uses simple string conversion.
   - `GetRequestAnother()`: Uses `strings.Builder` for better performance in large data processing.

Key Takeaways:
-----------------
✔ Always check response status before processing data.
✔ Use `defer` to ensure resources are properly released.
✔ `strings.Builder` can be more efficient for handling large responses.
*/
