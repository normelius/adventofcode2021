package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day1() {
	content, _ := os.Open("solutions/day1.txt")
	defer content.Close()
	scanner := bufio.NewScanner(content)

	var contents []int

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		contents = append(contents, value)
	}

	increases := 0
	for idx := 1; idx < len(contents); idx++ {
		diff := contents[idx] - contents[idx-1]
		if diff > 0 {
			increases++
		}
	}

	fmt.Println("Day1: answer 1:", increases)

	increases = 0
	for idx := 3; idx < len(contents); idx++ {
		prev := contents[idx-1] + contents[idx-2] + contents[idx-3]
		current := contents[idx] + contents[idx-1] + contents[idx-2]
		if current > prev {
			increases++
		}
	}

	fmt.Println("Day1: answer 2:", increases)

}
