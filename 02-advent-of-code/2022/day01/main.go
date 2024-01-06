package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, ferr := os.Open("C:/Users/chho/dev/repos/learning-golang/02-advent-of-code/2022/day01/input.txt")
	defer file.Close()
	if ferr != nil {
		panic(ferr)
	}

	var score *int = new(int)
	var topThreeElfs []int = []int{0, 0, 0}
	var sum *int = new(int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		var line string = scanner.Text()
		if line != "" {
			var num, _ = strconv.Atoi(line)
			*sum = num + *sum
		} else {
			if *sum > *score {
				*score = *sum
			}

			checkTopElfs(*score, topThreeElfs)
			*sum = 0
		}
	}

	if *sum != 0 && *sum > *score {
		*score = *sum
		checkTopElfs(*score, topThreeElfs)
	}

	partTwoScore := 0
	for _, value := range topThreeElfs {
		partTwoScore = partTwoScore + value
	}

	fmt.Println("\nHighest score: ", *score)
	fmt.Printf("\nTop three elfs: %d, with sum: %d", topThreeElfs, partTwoScore)
}

func checkTopElfs(calories int, topThreeElfs []int) {
	var index int = 0
	for i, value := range topThreeElfs {
		if value < topThreeElfs[index] {
			index = i
		}
	}
	topThreeElfs[index] = calories
}
