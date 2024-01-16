package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	defer input.Close()

	if err != nil {
		fmt.Println("Error opening file")
	}
	score := 0
	secondScore := 0
	sc := bufio.NewScanner(input)

	for sc.Scan() {
		var elf1 string = strings.Split(sc.Text(), ",")[0]
		var elf2 string = strings.Split(sc.Text(), ",")[1]
		elf1Start, elf1End, elf2Start, elf2End := splitEfls(elf1, elf2)
		checkForCaseOne(elf1Start, elf1End, elf2Start, elf2End, &score)
		checkForCaseTwo(elf1Start, elf1End, elf2Start, elf2End, &secondScore)
	}
	fmt.Println("Score: ", score)
	fmt.Println("SecondScore: ", secondScore)
}

func checkForCaseOne(elf1Start int, elf1End int, elf2Start int, elf2End int, score *int) {
	if elf1Start >= elf2Start && elf1End <= elf2End {
		fmt.Printf("elf: 1 is within elf 2's range\n")
		*score += 1
	} else if elf2Start >= elf1Start && elf2End <= elf1End {
		fmt.Printf("elf: 2 is within elf 1's range\n")
		*score += 1
	}
}

func checkForCaseTwo(elf1Start int, elf1End int, elf2Start int, elf2End int, score *int) {

	if elf1Start >= elf2Start && elf1Start <= elf2End {
		fmt.Printf("elf: 1 is within elf 2's range\n")
		*score += 1
	} else if elf2Start >= elf1Start && elf2Start <= elf1End {
		fmt.Printf("elf: 2 is within elf 1's range\n")
		*score += 1
	}

}

func splitEfls(elf1 string, elf2 string) (int, int, int, int) {
	elf1Start, _ := strconv.Atoi(strings.Split(elf1, "-")[0])
	elf2Start, _ := strconv.Atoi(strings.Split(elf2, "-")[0])
	elf1End, _ := strconv.Atoi(strings.Split(elf1, "-")[1])
	elf2End, _ := strconv.Atoi(strings.Split(elf2, "-")[1])
	return elf1Start, elf1End, elf2Start, elf2End
}
