package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	var list1 []int
	var list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "   ")

		l1, err := strconv.Atoi(text[0])
		if err != nil {
			log.Fatal(err)
		}
		l2, err := strconv.Atoi(text[1])
		if err != nil {
			log.Fatal(err)
		}

		list1 = append(list1, l1)
		list2 = append(list2, l2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0
	for i := 0; i < len(list1); i++ {
		d := abs(list1[i] - list2[i])
		sum += d
	}

	log.Println("Part 1 --")
	log.Println("The total distance is: ", sum)

	log.Println("Part 2 --")

	//var calculated []int
	sumPart2 := 0

	for i := 0; i < len(list1); i++ {
		n := list1[i]

		count := 0
		for j := 0; j < len(list2); j++ {
			if n == list2[j] {
				count += 1
			}
		}

		sumPart2 += count * n
	}

	log.Println("The total sum is: ", sumPart2)
}
