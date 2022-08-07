package main

import "fmt"

func outerFunc(name string) func() string {
	return func() string {
		return "Hello " + name
	}
}

func main() {
	innerFunc := outerFunc("shohan")

	fmt.Println(innerFunc())
}
