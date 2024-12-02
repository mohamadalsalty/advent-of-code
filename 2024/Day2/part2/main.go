package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isSafe checks if a report is safe based on Part One rules:
// 1. Levels are either all increasing or all decreasing.
// 2. Adjacent differences are between 1 and 3 inclusive.
func isSafe(levels []int) bool {
	if len(levels) == 0 {
		return false // Empty report is considered unsafe
	}

	var increasing *bool = nil

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		if diff == 0 {
			return false // No change is neither increasing nor decreasing
		}

		// Determine the trend (increasing or decreasing)
		if increasing == nil {
			temp := diff > 0
			increasing = &temp
		} else {
			currentTrend := diff > 0
			if *increasing != currentTrend {
				return false // Not consistently increasing/decreasing
			}
		}

		// Check if difference is within allowed range
		absDiff := diff
		if absDiff < 0 {
			absDiff = -absDiff
		}
		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}

	return true
}

// countSafeReports counts the number of safe reports considering the Problem Dampener.
func countSafeReports(reports [][]int) int {
	safeCount := 0

	for _, report := range reports {
		if isSafe(report) {
			safeCount++
			continue // No need to apply dampener
		}

		// Attempt to remove each level one by one and check safety
		wasSafe := false
		for i := 0; i < len(report); i++ {
			modified := append([]int{}, report[:i]...)
			modified = append(modified, report[i+1:]...)
			if isSafe(modified) {
				safeCount++
				wasSafe = true
				break // No need to try removing other levels
			}
		}

		if !wasSafe {
			// Report remains unsafe even after trying to remove each level
			// Do not increment safeCount
		}
	}

	return safeCount
}

// parseInput reads the input file and parses it into a slice of integer slices.
func parseInput(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}
		parts := strings.Fields(line)
		var levels []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid number '%s' in input", part)
			}
			levels = append(levels, num)
		}
		reports = append(reports, levels)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}

func main() {
	// Check command-line arguments for input file
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	filename := os.Args[1]
	reports, err := parseInput(filename)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	safeReports := countSafeReports(reports)
	fmt.Printf("Number of safe reports (Part Two): %d\n", safeReports)
}
