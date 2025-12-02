package main

import (
    "flag"
    "fmt"
    "regexp"
)

func is_valid_date(date string) bool {
	pattern := "(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)-(0[1-9]|[12][0-9]|3[01])"
	re := regexp.MustCompile(pattern)
    match := re.FindString(date)
    return match != ""
}

func main() {
    date := flag.String("date", "Jan-01", "Date to solve for")
    flag.Parse()

    if is_valid_date(*date) {
        fmt.Println("Solving a-puzzle-a-day for", *date)
    } else {
        fmt.Println("Invalid date:", *date)
    }
}
