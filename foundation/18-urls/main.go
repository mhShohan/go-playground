package main

import (
	"fmt"
	"net/url"
)

const uri = "https://localhost:5000/posts/1?name=shohan&age=20&city=NewYork"

func main() {
	fmt.Println("____________URLs______________")

	// Parse the URL
	result, err := url.Parse(uri)
	checkNilError(err)

	// Display URL components
	fmt.Println("Scheme:", result.Scheme)     // Protocol (e.g., http, https)
	fmt.Println("Host:", result.Host)         // Host (e.g., domain name or IP with port)
	fmt.Println("Port:", result.Port())       // Extract port from the host
	fmt.Println("Path:", result.Path)         // Path after domain
	fmt.Println("RawQuery:", result.RawQuery) // Full query string
	fmt.Println("Fragment:", result.Fragment) // URL fragment (if any)
	fmt.Println("String:", result.String())   // Reconstruct URL from parsed parts
	fmt.Println("Query:", result.Query())     // Parse query parameters

	// Handling Query Parameters
	queryParams := result.Query()

	fmt.Println("____________Query Params________________")
	fmt.Println("Name:", queryParams.Get("name")) // Get a specific query parameter

	// Iterate over all query parameters
	for key, values := range queryParams {
		fmt.Println(key, ":", values)
	}

	// Adding and Modifying Query Parameters
	queryParams.Set("country", "USA")
	fmt.Println("Updated Query String:", queryParams.Encode())

	// Constructing a New URL
	fmt.Println("____________Construct URL_____________")
	constructedURL := &url.URL{
		Scheme:   "https",
		Host:     "localhost:5000",
		Path:     "/posts/2",
		RawQuery: queryParams.Encode(),
	}
	fmt.Println("Constructed URL:", constructedURL.String())
}

// checkNilError checks if the error is not nil and panics
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

/*
Notes:
---------
1. **URL Components**:
   - `Scheme`: Protocol (e.g., `http`, `https`).
   - `Host`: Domain name and port.
   - `Path`: Specific resource path.
   - `RawQuery`: Query parameters in key-value pairs.
   - `Fragment`: Optional section after `#` (not used in this example).

2. **Query Parameters**:
   - Use `Query().Get("param")` to fetch a single value.
   - `Query()` returns a `map[string][]string` to handle multiple values.
   - Use `Set()` to modify/add query parameters.
   - `Encode()` converts query parameters to a URL-encoded string.

3. **Constructing a URL**:
   - Use `url.URL{}` to manually construct a URL.
   - Ensure `RawQuery` is properly encoded.

4. **Error Handling**:
   - Always handle errors when parsing URLs (`url.Parse()` can return errors).

Key Takeaways:
-----------------
✔ Understanding URL components is essential for web scraping, API calls, and request handling.
✔ Query parameters can be easily modified and accessed.
✔ Constructing URLs manually is useful when generating dynamic links.
*/
