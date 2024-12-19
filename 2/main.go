package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func notSafe(a1 int, a2 int) bool {
	diff := abs(a1 - a2)
	return diff == 0 || diff > 3
}

func canRecoverBruteForce(row []int) bool {
	if isRowSafe(row) {
		return true
	}

	for i := 0; i < len(row)-1; i++ {
		if isRowSafeIgnoreIdx(row, i) {
			return true
		}
	}

	if isRowSafeIgnoreIdx(row, len(row)-1) {
		return true
	}

	return false
}

// 8 6 4 4 1
func isRowSafe(row []int) bool {
	increasing, decreasing := true, true

	for i := 0; i < len(row)-1; i++ {
		if row[i] < row[i+1] {
			decreasing = false
		}
		if row[i] > row[i+1] {
			increasing = false
		}

		if !decreasing && !increasing {
			return false
		}

		if notSafe(row[i], row[i+1]) {
			return false
		}
	}

	return true
}

func isRowSafeIgnoreIdx(row []int, ignoreIdx int) bool {
	increasing, decreasing := true, true

	for i := 0; i < len(row)-1; i++ {
		current, next := i, i+1
		if i+1 == ignoreIdx {
			if i+2 == len(row) {
				continue
			}
			next = i + 2
		}

		if i == ignoreIdx {
			if i == 0 {
				continue
			}
			current = i - 1
		}

		if row[current] < row[next] {
			decreasing = false
		}
		if row[current] > row[next] {
			increasing = false
		}

		if !decreasing && !increasing {
			return false
		}

		if notSafe(row[current], row[next]) {
			return false
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
		if isRowSafe(row) {
			count += 1
		}

		if canRecoverBruteForce(row) {
			fmt.Printf("safe row %+v\n", row)
			count2 += 1
		} else {
			fmt.Printf("unsafe row %+v\n", row)
		}
	}

	fmt.Printf("Part1: total safe rows count is %d out of %d rows\n", count, len(reports))
	fmt.Printf("Part2: total safe rows count is %d out of %d rows\n", count2, len(reports))
}
