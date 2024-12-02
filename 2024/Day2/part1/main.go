package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function to check if a report is safe
func isSafeReport(parts []string) bool {
	// Convert the string parts to integers
	levels := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Error converting to integer at part %d: %s\n", i, part)
			return false
		}
		levels[i] = num
	}

	// Check monotonicity and difference bounds
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		// Difference must be strictly between -3 and 3 (inclusive) but not 0
		if diff < -3 || diff > 3 || diff == 0 {
			fmt.Printf("Difference out of bounds at index %d: %d -> %d (diff: %d)\n", i-1, levels[i-1], levels[i], diff)
			return false
		}

		// Check monotonicity
		if diff > 0 {
			isDecreasing = false
		}
		if diff < 0 {
			isIncreasing = false
		}
	}

	// Debug: Print the monotonicity result
	fmt.Printf("Levels: %v | Increasing: %t, Decreasing: %t\n", levels, isIncreasing, isDecreasing)

	// Return true if the report is either all increasing or all decreasing
	return isIncreasing || isDecreasing
}

func main() {
	// Open the input file
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeReports := 0

	// Read each report line
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		// Skip empty lines or single-number reports
		if len(parts) <= 1 {
			continue
		}

		// Debug: Print the current line
		fmt.Printf("Processing line: %s\n", line)

		if isSafeReport(parts) {
			safeReports++
			fmt.Printf("Line is safe: %s\n", line)
		} else {
			fmt.Printf("Line is unsafe: %s\n", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Print the number of safe reports
	fmt.Printf("Number of safe reports: %d\n", safeReports)
}
