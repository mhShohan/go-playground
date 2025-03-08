package main

import (
	"fmt"
	"time"
)

func main() {

	// Defining date-time formats
	const format = "2006-01-02 15:04:05 Monday"
	const format2 = "2006-January-02 15:04:05 Monday"

	// Getting current time
	presentTime := time.Now()

	// Formatting current time in different formats
	fmt.Println("Current time:", presentTime.Format(format))
	fmt.Println("Current time:", presentTime.Format(format2))

	// Creating a specific date
	date := time.Date(1996, time.April, 17, 0, 0, 0, 0, time.UTC)

	// Formatting and printing a specific date
	fmt.Println("My birthday:", date.Format(format2))
}

/*
Notes:
1. **Date-Time Formatting in Go:**
   - The reference time for formatting in Go is: `Mon Jan 2 15:04:05 MST 2006`.
   - This specific time must be used as a pattern when defining date-time formats.
   - Example format: `"2006-01-02 15:04:05 Monday"` (YYYY-MM-DD HH:MM:SS Day).

2. **Getting Current Time:**
   - `time.Now()` retrieves the current date and time.
   - `Format()` is used to convert it into a human-readable string.

3. **Creating a Specific Date:**
   - `time.Date(year, month, day, hour, min, sec, sec, location)` is used to define a custom date.
   - Example: `time.Date(1996, time.April, 17, 0, 0, 0, 0, time.UTC)` creates April 17, 1996.

4. **Time Zones:**
   - `time.UTC` sets the time in Coordinated Universal Time (UTC).
   - Other time zones can be used with `time.LoadLocation("America/New_York")`.

5. **Useful Commands:**
   - `go env` → Displays Go environment variables.
   - `go build` → Compiles the Go program into an executable file.
   - `go run filename.go` → Runs the Go program directly.

6. **Reference:**
   - Go `time` package documentation: https://pkg.go.dev/time
*/
