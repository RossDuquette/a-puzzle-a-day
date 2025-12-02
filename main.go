package main

import (
    "flag"
    "fmt"
    "regexp"
)

import (
    "a-puzzle-a-day/internal/solver"
)

func is_valid_date(date string) bool {
	pattern := "(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)-(0[1-9]|[12][0-9]|3[01])"
	re := regexp.MustCompile(pattern)
    match := re.FindString(date)
    return match != ""
}

func extract_month_day(date string) (month string, day string) {
    return date[:3], date[4:]
}

func main() {
    date := flag.String("date", "Jan-01", "Date to solve for")
    flag.Parse()

    if is_valid_date(*date) {
        month, day := extract_month_day(*date)
        fmt.Println("Solving a-puzzle-a-day for", month, day)
        solver.Solve(month, day)
    } else {
        fmt.Println("Invalid date:", *date)
    }
}
