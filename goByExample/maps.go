package main

import "fmt"

func main() {
	myMaps := make(map[string]int)

	myMaps["age"] = 24
	myMaps["count"] = 24

	fmt.Println(len(myMaps))
}
