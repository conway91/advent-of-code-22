package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Advent of code Day 1. Elf Calorie Counter")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	calorieFilePath := os.Args[1]
	fmt.Println("Calorie file path:", calorieFilePath)

	amountOfElvesToGet, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of top elves to get:", amountOfElvesToGet)

	calorieList := getCalorieCountFromFile(calorieFilePath)

	fmt.Println("PART 1")
	getElfWithHigestCalorieCount(calorieList)

	fmt.Println("PART 2")
	getTopElvesWithHigestCalorieCount(calorieList, amountOfElvesToGet)

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
}

func getCalorieCountFromFile(calorieFilePath string) map[int]int {
	file, err := os.Open(calorieFilePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	elfCalorieList := make(map[int]int)
	currentCalorieCount := 0
	elfNumber := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()

		if value == "" {
			elfCalorieList[elfNumber] = currentCalorieCount
			currentCalorieCount = 0
			elfNumber++

		} else {
			calroies, err := strconv.Atoi(value)

			if err != nil {
				log.Fatal(err)
			}

			currentCalorieCount += calroies
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return elfCalorieList
}

func getElfWithHigestCalorieCount(calorieList map[int]int) {
	fmt.Println("Checking elf with highest calorie count ...")
	maxCalroies := 0
	elfNumber := 0
	for elf, calroies := range calorieList {
		if calroies > maxCalroies {
			maxCalroies = calroies
			elfNumber = elf
		}
	}

	fmt.Println("Elf", elfNumber, "Has the higest calories:", maxCalroies)
}

func getTopElvesWithHigestCalorieCount(calorieList map[int]int, amountOfElvesToGet int) {
	fmt.Println("Checking top", amountOfElvesToGet, "elves with the highest calorie count ...")

	elves := make([]int, 0, len(calorieList))
	for elf := range calorieList {
		elves = append(elves, elf)
	}
	sort.Slice(elves, func(i, j int) bool { return calorieList[elves[i]] > calorieList[elves[j]] })

	totalCalories := 0
	for count, key := range elves {
		fmt.Println("Elf:", key, "Calories:", calorieList[key])
		totalCalories += calorieList[key]

		if count+1 == amountOfElvesToGet {
			fmt.Println("Total combined calories:", totalCalories)
			break
		}
	}
}
