package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var left []int
	var right []int
	var res int
	file, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		leftVal, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting left value:", err)
			continue
		}
		rightVal, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting right value:", err)
			continue
		}

		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	for _, value := range left {
		count := 0
		for _, innerValue := range right {
			if value == innerValue {
				count++
			}
		}

		res = res + (count * value)
	}

	print(res)

}
