package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var score int
	var currentNum int

	maxCalories1 := 0
	maxCalories2 := 0
	maxCalories3 := 0

	for sc.Scan() {
		if sc.Text() != "" {
			num, _ := strconv.Atoi(sc.Text())
			currentNum += num
		} else {
			if currentNum > score {
				score = currentNum
			}
			if currentNum > maxCalories3 {
				maxCalories3 = currentNum
			}
			if maxCalories3 > maxCalories2 {
				maxCalories3, maxCalories2 = maxCalories2, maxCalories3
			}
			if maxCalories2 > maxCalories1 {
				maxCalories2, maxCalories1 = maxCalories1, maxCalories2
			}

			currentNum = 0
		}
	}
	fmt.Println(score)
	fmt.Println(maxCalories1 + maxCalories2 + maxCalories3)
}
