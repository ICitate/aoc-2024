package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	rules := make(map[string][]string)
	var updates [][]string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			for scanner.Scan() {
				// parse updates
				update := strings.Split(scanner.Text(), ",")
				updates = append(updates, update)
			}
			break
		}

		// parse rules
		rule := strings.Split(line, "|")
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for k, v := range rules {
		fmt.Printf("%s goes before %+v\n", k, v)
	}

	result1 := 0
	result2 := 0
	for _, update := range updates {
		valid := true
		for i := 1; i < len(update); i++ {
			rule := rules[update[i]]
			if slices.Contains(rule, update[i-1]) {
				valid = false
			}
		}

		if valid {
			fmt.Printf("%+v\n", update)
			middle := update[len(update)/2]
			fmt.Printf("middle %+v\n", middle)
			n, err := strconv.Atoi(middle)
			if err != nil {
				panic(err)
			}
			result1 += n
		} else {
			fmt.Println("####")
			fmt.Printf("fixing: %+v\n", update)
			for k := 0; k < len(update); k++ {
				for i := 1; i < len(update); i++ {
					rule := rules[update[i]]
					if slices.Contains(rule, update[i-1]) {
						update[i-1], update[i] = update[i], update[i-1]
					}
				}
			}
			fmt.Printf("fixed: %+v\n", update)
			fmt.Println("####")

			middle := update[len(update)/2]
			fmt.Printf("middle %+v\n", middle)
			n, err := strconv.Atoi(middle)
			if err != nil {
				panic(err)
			}
			result2 += n
		}
	}

	fmt.Printf("PART 1: result is %d\n", result1)
	fmt.Printf("PART 2: result is %d\n", result2)
}
