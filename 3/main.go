package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return true
}

func parseDigit(scanner *bufio.Scanner, endChar string) (int, error) {
	var n string

	scanner.Scan()
	char := scanner.Text()
	if !isInteger(char) {
		return 0, errors.New("")
	}
	n += char

	for i := 0; i < 2; i++ {
		scanner.Scan()
		char = scanner.Text()

		if char == endChar {
			result, err := strconv.Atoi(n)
			if err != nil {
				return 0, err
			}
			return result, nil
		}

		if !isInteger(char) {
			return 0, errors.New("")
		}

		n += char
	}

	scanner.Scan()
	char = scanner.Text()

	if char != endChar {
		return 0, errors.New("")
	}

	result, err := strconv.Atoi(n)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	result1 := 0
	result2 := 0
	enabled := true
	for scanner.Scan() {
		char := scanner.Text()

		if char == "d" {
			scanner.Scan()
			char = scanner.Text()
			if char != "o" {
				continue
			}

			scanner.Scan()
			char = scanner.Text()

			if char == "(" {
				scanner.Scan()
				char = scanner.Text()
				if char != ")" {
					continue
				}
				enabled = true
				continue
			}

			if char != "n" {
				continue
			}

			scanner.Scan()
			char = scanner.Text()
			if char != "'" {
				continue
			}

			scanner.Scan()
			char = scanner.Text()
			if char != "t" {
				continue
			}

			scanner.Scan()
			char = scanner.Text()
			if char != "(" {
				continue
			}

			scanner.Scan()
			char = scanner.Text()
			if char != ")" {
				continue
			}

			enabled = false
		}

		if char == "m" {
			scanner.Scan()
			char = scanner.Text()
			if char != "u" {
				continue
			}

			scanner.Scan()
			char = scanner.Text()
			if char != "l" {
				continue
			}

			scanner.Scan()
			char = scanner.Text()
			if char != "(" {
				continue
			}

			// number parsing (digit 1-3)
			n1, err := parseDigit(scanner, ",")
			if err != nil {
				continue
			}
			n2, err := parseDigit(scanner, ")")
			if err != nil {
				continue
			}

			fmt.Println("n1: ", n1, "n2: ", n2, "enabled: ", enabled)
			result1 += (n1 * n2)
			if enabled {
				result2 += (n1 * n2)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PART 1: result is %d\n", result1)
	fmt.Printf("PART 2: result is %d\n", result2)
}
