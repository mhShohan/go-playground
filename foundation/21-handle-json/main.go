package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"title"`          // Change the field name in JSON output
	Price    int      `json:"price"`          // Change the field name in JSON output
	Platform string   `json:"website"`        // Change the field name in JSON output
	Password string   `json:"-"`              // Exclude the field from JSON output
	Tags     []string `json:"tags,omitempty"` // Change the field name in JSON output and exclude the field if it is empty
}

func main() {
	fmt.Println("____________Handle JSON____________")

	// Encode JSON
	EncodeJson()

	// Decode JSON
	DecodeJson()
}

func checkNil(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
-----------------------------------
Encoding JSON
-----------------------------------
*/
func EncodeJson() {
	courses := []Course{
		{Name: "Docker Deep Dive", Price: 299, Platform: "Udemy", Password: "password", Tags: []string{"docker", "devops", "container"}},
		{Name: "Docker Clustering", Price: 199, Platform: "Udemy", Password: "password", Tags: []string{"docker", "devops", "cluster"}},
		{Name: "Docker and Kubernetes", Price: 199, Platform: "Udemy", Password: "password", Tags: []string{"docker", "kubernetes", "devops"}},
		{Name: "Docker Mastery", Price: 299, Platform: "Youtube", Password: "password", Tags: nil},
	}

	// Encode JSON using MarshalIndent for readability
	jsonOutput, err := json.MarshalIndent(courses, "", "   ")
	checkNil(err)

	fmt.Println("Encoded JSON Output:")
	fmt.Println(string(jsonOutput))
}

/*
-----------------------------------
Decoding JSON
-----------------------------------
*/
func DecodeJson() {
	jsonData := []byte(`
		[
			{
				"title": "Docker Deep Dive",
				"price": 299,
				"website": "Udemy",
				"tags": [ "docker", "devops", "container" ]
			},
			{
				"title": "Docker Clustering",
				"price": 199,
				"website": "Udemy",
				"tags": [ "docker", "devops", "cluster" ]
			},
			{
				"title": "Docker and Kubernetes",
				"price": 199,
				"website": "Udemy",
				"tags": [ "docker", "kubernetes", "devops" ]
			},
			{
				"title": "Docker Mastery",
				"price": 299,
				"website": "Youtube"
			}
		]
	`)

	// Validate JSON before unmarshalling
	var courses []Course
	isValidJson := json.Valid(jsonData)

	if isValidJson {
		err := json.Unmarshal(jsonData, &courses)
		checkNil(err)

		fmt.Println("Decoded JSON into Struct:")
		fmt.Printf("%+v\n", courses) // Print field names along with values
	} else {
		fmt.Println("Invalid JSON")
	}

	fmt.Println("__________Another way to decode JSON_____________________")
	// Decode JSON into a generic map
	var coursesMap []map[string]interface{}
	err := json.Unmarshal(jsonData, &coursesMap)
	checkNil(err)
	fmt.Println("Decoded JSON into Map:")
	fmt.Printf("%#v\n", coursesMap)
}

/*
Notes:
---------
1. **JSON Encoding in Go**:
   - `json.Marshal()` encodes Go structs into JSON.
   - `json.MarshalIndent()` provides formatted JSON for readability.
   - Struct tags (`json:"name"`) allow field renaming and formatting options.

2. **JSON Decoding in Go**:
   - `json.Unmarshal()` decodes JSON into Go structs or maps.
   - JSON validation using `json.Valid()` helps ensure the correctness of data before parsing.

3. **Handling JSON Data**:
   - Use `omitempty` in struct tags to exclude empty fields from JSON output.
   - Use `json:"-"` to ignore fields in JSON encoding.
   - Decoding into a `map[string]interface{}` allows working with dynamic JSON structures.

Key Takeaways:
-----------------
✔ Use `json.MarshalIndent()` for readable JSON output.
✔ Validate JSON before decoding to prevent errors.
✔ `map[string]interface{}` is useful for handling unknown JSON structures.
*/
