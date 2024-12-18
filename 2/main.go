package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// 1 5 7 9 11
// 1 2 5 4 7
// 1 2 3 10
// 10 1 2 3 4 5

func isSafe2(row []int) bool {

	// middle safety
	isIncreasing := row[1] < row[2]

	for i := 1; i < len(row)-2; i++ {
		if isIncreasing {
			if row[i] >= row[i+1] {
				return false
			}

			if row[i+1]-row[i] > 3 {
				return false
			}
		} else {
			if row[i] <= row[i+1] {
				return false
			}

			if row[i]-row[i+1] > 3 {
				return false
			}
		}
	}
	return true
}

func isSafe(row []int) bool {
	isIncreasing := row[0] < row[1]

	for i := 0; i < len(row)-1; i++ {
		if isIncreasing {
			if row[i] >= row[i+1] {
				return false
			}

			if row[i+1]-row[i] > 3 {
				return false
			}
		} else {
			if row[i] <= row[i+1] {
				return false
			}

			if row[i]-row[i+1] > 3 {
				return false
			}
		}
	}
	return true
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	var reports [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")

		var row []int
		for _, s := range split {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, n)
		}

		reports = append(reports, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0
	count2 := 0
	for _, row := range reports {
		if isSafe(row) {
			count += 1
		}

		if isSafe2(row) {
			count2 += 1
		}
	}

	log.Println("Part1: total safe rows count is ", count)
	log.Println("Part2: total safe rows count is ", count)
}
