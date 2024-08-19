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
	// Handle JSON
	// JSON is a popular format for data exchange between services. Go has excellent support for encoding and decoding JSON data. The encoding/json package provides functions to encode and decode JSON data.
	// EncodeJson()
	DecodeJson()
}

func checkNil(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func EncodeJson() {
	courses := []Course{
		{Name: "Docker Deep Dive", Price: 299, Platform: "Udemy", Password: "password", Tags: []string{"docker", "devops", "container"}},
		{Name: "Docker Clustering", Price: 199, Platform: "Udemy", Password: "password", Tags: []string{"docker", "devops", "cluster"}},
		{Name: "Docker and Kubernetes", Price: 199, Platform: "Udemy", Password: "password", Tags: []string{"docker", "kubernetes", "devops"}},
		{"Docker Mastery", 299, "Youtube", "password", nil},
	}

	// Encode JSON
	// The json.Marshal function converts a Go value to its JSON representation. It returns a byte slice containing the JSON data.
	// The json.MarshalIndent function formats the JSON data with indentation for better readability.
	// The json.NewEncoder function creates an encoder that writes JSON data to an io.Writer.
	jsonOutput, err := json.MarshalIndent(courses, "", "   ")
	checkNil(err)

	fmt.Println(string(jsonOutput))
	// fmt.Printf("%s\n", jsonOutput)
}

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

	// Decode JSON
	// The json.Unmarshal function converts JSON data to a Go value. It takes a byte slice containing the JSON data and a pointer to a value that will hold the decoded data.
	// The json.NewDecoder function creates a decoder that reads JSON data from an io.Reader.
	var courses []Course
	isValidJson := json.Valid(jsonData)

	if isValidJson {
		err := json.Unmarshal(jsonData, &courses)
		checkNil(err)

		fmt.Printf("%+v\n", courses) // %+v prints the field names along with the values
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("%#v\n", courses) // %#v prints the Go syntax representation of the value
		// fmt.Println(courses)
	} else {
		fmt.Println("Invalid JSON")
	}

	fmt.Println("__________another way_____________________")
	// Another way
	var courses2 map[string]interface{}
	err := json.Unmarshal(jsonData, &courses2)
	checkNil(err)
	fmt.Printf("%#v\n", courses)
}
