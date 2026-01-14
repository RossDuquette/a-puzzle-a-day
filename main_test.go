package main

import (
	"testing"
)

func TestDateValidity(t *testing.T) {
	tests := []struct {
		input string
		valid bool
	}{
		{"Jan-04", true},
		{"Apr-30", true},
		{"Jul-4", false},
		{"AUG-11", false},
	}
	for _, test := range tests {
		testName := test.input
		t.Run(testName, func(t *testing.T) {
			valid := isValidDate(test.input)
			if valid != test.valid {
				t.Errorf("'%s' was marked as %t", test.input, valid)
			}
		})
	}
}

func TestExtractMonthDay(t *testing.T) {
	month, day := extractMonthDay("May-31")
	if month != "May" {
		t.Errorf("Failed to extract May")
	} else if day != "31" {
		t.Errorf("Failed to extract 31")
	}

	month, day = extractMonthDay("Jun-01")
	if month != "Jun" {
		t.Errorf("Failed to extract Jun")
	} else if day != "01" {
		t.Errorf("Failed to extract 01")
	}
}
