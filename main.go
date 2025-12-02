package main

import (
    "flag"
    "fmt"
)

func main() {
    date := flag.String("date", "Jan-01", "Date to solve for")
    flag.Parse()

    fmt.Println("Solving a-puzzle-a-day for", *date)
}
