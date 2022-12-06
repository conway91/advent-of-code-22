package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Rock     string = "rock"
	Paper           = "paper"
	Scissors        = "scissors"
	Win             = "win"
	Draw            = "draw"
	Lose            = "lose"
)

var handScore = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var opponentsHands = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
}

var playersHands = map[string]string{
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

var gameResults = map[string]string{
	"X": "lose",
	"Y": "draw",
	"Z": "win",
}

func main() {
	fmt.Println("Advent of code Day 2. Rock Paper Scissors")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	strategyFilePath := os.Args[1]
	fmt.Println("Strategy file path:", strategyFilePath)

	games := getStrategyMapFromFile(strategyFilePath)

	totalScorePartOne := getGameResultFromStrategyPartOne(games)
	fmt.Println("Part 1 Final Score:", totalScorePartOne)

	totalScorePartTwo := getGameResultFromStrategyPartTwo(games)
	fmt.Println("Part 2 Final Score:", totalScorePartTwo)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func getStrategyMapFromFile(strategyFilePath string) []string {
	file, err := os.Open(strategyFilePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	stategyList := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stategyList = append(stategyList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return stategyList
}

func getGameResultFromStrategyPartOne(games []string) int {
	totalScore := 0

	for _, game := range games {
		hands := strings.Split(game, " ")
		opponentsHand := opponentsHands[hands[0]]
		playersHand := playersHands[hands[1]]

		if playersHand == Rock {
			totalScore += rockResults(opponentsHand)
		}

		if playersHand == Paper {
			totalScore += paperResults(opponentsHand)
		}

		if playersHand == Scissors {
			totalScore += scissorsResults(opponentsHand)
		}
	}

	return totalScore
}

func getGameResultFromStrategyPartTwo(games []string) int {
	totalScore := 0

	for _, game := range games {
		hands := strings.Split(game, " ")
		opponentsHand := opponentsHands[hands[0]]
		gameResult := gameResults[hands[1]]

		if opponentsHand == Rock {
			totalScore += opponentRockHand(gameResult)
		}

		if opponentsHand == Paper {
			totalScore += opponentPaperHand(gameResult)
		}

		if opponentsHand == Scissors {
			totalScore += opponentScissorHand(gameResult)
		}
	}

	return totalScore
}

func rockResults(opponentsHand string) int {
	if opponentsHand == Rock {
		return 3 + handScore[Rock]
	}

	if opponentsHand == Scissors {
		return 6 + handScore[Rock]
	}

	return handScore[Rock]
}

func paperResults(opponentsHand string) int {
	if opponentsHand == Paper {
		return 3 + handScore[Paper]
	}

	if opponentsHand == Rock {
		return 6 + handScore[Paper]
	}

	return handScore[Paper]
}

func scissorsResults(opponentsHand string) int {
	if opponentsHand == Scissors {
		return 3 + handScore[Scissors]
	}

	if opponentsHand == Paper {
		return 6 + handScore[Scissors]
	}

	return handScore[Scissors]
}

func opponentRockHand(gameResult string) int {
	if gameResult == Draw {
		return 3 + handScore[Rock]
	}

	if gameResult == Win {
		return 6 + handScore[Paper]
	}

	return handScore[Scissors]
}

func opponentPaperHand(gameResult string) int {
	if gameResult == Draw {
		return 3 + handScore[Paper]
	}

	if gameResult == Win {
		return 6 + handScore[Scissors]
	}

	return handScore[Rock]
}

func opponentScissorHand(gameResult string) int {
	if gameResult == Draw {
		return 3 + handScore[Scissors]
	}

	if gameResult == Win {
		return 6 + handScore[Rock]
	}

	return handScore[Paper]
}
