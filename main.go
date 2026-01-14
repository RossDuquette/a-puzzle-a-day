package main

import (
	"flag"
	"fmt"
	"regexp"
)

import (
	"a-puzzle-a-day/internal/solver"
	"a-puzzle-a-day/internal/verifier"
)

func isValidDate(date string) bool {
	pattern := "(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec)-(0[1-9]|[12][0-9]|3[01])"
	re := regexp.MustCompile(pattern)
	match := re.FindString(date)
	return match != ""
}

func extractMonthDay(date string) (month string, day string) {
	return date[:3], date[4:]
}

func main() {
	date := flag.String("date", "Jan-01", "Date to solve for")
	saveToFiles := flag.Bool("to-file", false, "Save solutions to files")
	verify := flag.Bool("verify", false, "Verify that solution files are unique")
	flag.Parse()

	if isValidDate(*date) {
		month, day := extractMonthDay(*date)
		if *verify {
			verifier.CheckUniqueness(month, day)
		} else {
			fmt.Println("Solving a-puzzle-a-day for", month, day)
			solver.Solve(month, day, *saveToFiles)
		}
	} else {
		fmt.Println("Invalid date:", *date)
	}
}
