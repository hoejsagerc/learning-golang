package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const ROCK int = 1
const PAPER int = 2
const SICSSORS int = 3
const WIN int = 6
const DRAW int = 3
const LOSS int = 0

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	score := 0
	secondScore := 0

	for sc.Scan() {
		moves := strings.Split(sc.Text(), "")
		calcScoreForGame(&score, moves)
		advanceCalcScoreForGame(&secondScore, moves)
	}
	fmt.Println("Score: ", score)
	fmt.Println("Second Score: ", secondScore)
}

func advanceCalcScoreForGame(score *int, moves []string) {
	opponent := strings.TrimSpace(moves[0])
	tactic := strings.TrimSpace(moves[2])

	switch {
	// when opponent chooses rock
	case opponent == "A" && tactic == "X":
		*score += (SICSSORS + LOSS)
	case opponent == "A" && tactic == "Y":
		*score += (ROCK + DRAW)
	case opponent == "A" && tactic == "Z":
		// when opponent chooses paper
		*score += (PAPER + WIN)
	case opponent == "B" && tactic == "X":
		*score += (ROCK + LOSS)
	case opponent == "B" && tactic == "Y":
		*score += (PAPER + DRAW)
	case opponent == "B" && tactic == "Z":
		// when opponent chooses sicssors
		*score += (SICSSORS + WIN)
	case opponent == "C" && tactic == "X":
		*score += (PAPER + LOSS)
	case opponent == "C" && tactic == "Y":
		*score += (SICSSORS + DRAW)
	case opponent == "C" && tactic == "Z":
		*score += (ROCK + WIN)
	}
}

func calcScoreForGame(score *int, moves []string) {
	opponent := strings.TrimSpace(moves[0])
	you := strings.TrimSpace(moves[2])

	switch {
	// when opponent chooses rock
	case opponent == "A" && you == "X":
		*score += (ROCK + DRAW)
	case opponent == "A" && you == "Y":
		*score += (PAPER + WIN)
	case opponent == "A" && you == "Z":
		*score += (SICSSORS + LOSS)
		// when opponent chooses paper
	case opponent == "B" && you == "X":
		*score += (ROCK + LOSS)
	case opponent == "B" && you == "Y":
		*score += (PAPER + DRAW)
	case opponent == "B" && you == "Z":
		*score += (SICSSORS + WIN)
		// when opponent chooses sicssors
	case opponent == "C" && you == "X":
		*score += (ROCK + WIN)
	case opponent == "C" && you == "Y":
		*score += (PAPER + LOSS)
	case opponent == "C" && you == "Z":
		*score += (SICSSORS + DRAW)
	}
}
