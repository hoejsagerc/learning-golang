package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var items string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {

	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	score := 0
	secondScore := 0
	lineCount := 0
	var elfs []string

	for sc.Scan() {
		retrieveSimilarItems(sc.Text(), &score)
		fmt.Println("Looking at line: ", sc.Text())
		if lineCount < 2 {
			fmt.Printf("Adding elf: %v\n", lineCount)
			elfs = append(elfs, sc.Text())
			lineCount++
		} else if lineCount == 2 {
			elfs = append(elfs, sc.Text())
			fmt.Printf("Checking line count: %v\n", lineCount)
			fmt.Printf("Three elfs has been found, runnning checks on them and starting over\n")
			findElfBadges(elfs, &secondScore)
			lineCount = 0
			elfs = []string{}
			lineCount = 0
		}
	}

	fmt.Printf("Score: %v\n", score)
	fmt.Printf("SecondScore: %v\n", secondScore)
}

func findElfBadges(elfs []string, score *int) {
	var foundItems []rune
	for _, char := range elfs[0] {
		// checking if the char is in any of the other elfs
		if runeInSlice(char, []rune(elfs[1])) && runeInSlice(char, []rune(elfs[2])) {
			// if char is found then add to foundItems
			if runeInSlice(char, foundItems) == false {
				fmt.Printf("Found item: %c\n", char)
				foundItems = append(foundItems, char)
			}
		}
	}

	for _, char := range foundItems {
		index := strings.IndexRune(items, char)
		*score += index + 1
	}
}

func retrieveSimilarItems(text string, score *int) {
	var foundItems []rune
	var textArray []rune = []rune(text)
	var firstHalf []rune = textArray[:len(textArray)/2]
	var secondHalf []rune = textArray[len(textArray)/2:]
	// fmt.Printf("First half: %v - Second half: %v\n", string(firstHalf), string(secondHalf))

	for _, char := range firstHalf {
		if runeInSlice(char, secondHalf) {
			index := strings.IndexRune(items, char)
			if index != -1 {
				if runeInSlice(char, foundItems) == false {
					foundItems = append(foundItems, char)
					*score += index + 1
					// fmt.Printf("Found new item: %c int items with index: %v\n", char, index)
				}
			} else {
				// fmt.Printf("item: %c was not found in in items with index: %v\n", char, index)
			}
		}
	}
}

func runeInSlice(char rune, list []rune) bool {
	for _, c := range list {
		if c == char {
			return true
		}
	}
	return false
}
