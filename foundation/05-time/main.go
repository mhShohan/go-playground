package main

import (
	"fmt"
	"time"
)

func main() {

	const format = "2006-01-02 15:04:05 Monday"
	const format2 = "2006-January-02 15:04:05 Monday"

	presentTime := time.Now()

	fmt.Println("Current time:", presentTime.Format(format))
	fmt.Println("Current time:", presentTime.Format(format2))

	date := time.Date(1996, time.April, 17, 0, 0, 0, 0, time.UTC)

	fmt.Println("My birthday:", date.Format(format2))
}

// go env
// go build
