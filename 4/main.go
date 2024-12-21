package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// diag-dr i+1, j+1
func diagDr(matrix []string, i int, j int) bool {
	if i+3 >= len(matrix) || j+3 >= len(matrix[i]) {
		return false
	}

	if matrix[i+1][j+1] == 'M' {
		if matrix[i+2][j+2] == 'A' {
			if matrix[i+3][j+3] == 'S' {
				return true
			}
		}
	}

	return false
}

// diag-dl i+1, j-1
func diagDl(matrix []string, i int, j int) bool {
	if i+3 >= len(matrix) || j-3 < 0 {
		return false
	}

	if matrix[i+1][j-1] == 'M' {
		if matrix[i+2][j-2] == 'A' {
			if matrix[i+3][j-3] == 'S' {
				return true
			}
		}
	}

	return false
}

// diag-ur i-1, j+1
func diagUr(matrix []string, i int, j int) bool {
	if i-3 < 0 || j+3 >= len(matrix[i]) {
		return false
	}

	if matrix[i-1][j+1] == 'M' {
		if matrix[i-2][j+2] == 'A' {
			if matrix[i-3][j+3] == 'S' {
				return true
			}
		}
	}

	return false
}

// diag-ul i-1, j-1
func diagUl(matrix []string, i int, j int) bool {
	if i-3 < 0 || j-3 < 0 {
		return false
	}

	if matrix[i-1][j-1] == 'M' {
		if matrix[i-2][j-2] == 'A' {
			if matrix[i-3][j-3] == 'S' {
				return true
			}
		}
	}

	return false
}

func right(matrix []string, i int, j int) bool {
	if j+3 >= len(matrix[i]) {
		return false
	}

	if matrix[i][j+1] == 'M' {
		if matrix[i][j+2] == 'A' {
			if matrix[i][j+3] == 'S' {
				return true
			}
		}
	}

	return false
}

func left(matrix []string, i int, j int) bool {
	if j-3 < 0 {
		return false
	}

	if matrix[i][j-1] == 'M' {
		if matrix[i][j-2] == 'A' {
			if matrix[i][j-3] == 'S' {
				return true
			}
		}
	}

	return false
}

func down(matrix []string, i int, j int) bool {
	if i+3 >= len(matrix) {
		return false
	}

	if matrix[i+1][j] == 'M' {
		if matrix[i+2][j] == 'A' {
			if matrix[i+3][j] == 'S' {
				return true
			}
		}
	}

	return false
}

func up(matrix []string, i int, j int) bool {
	if i-3 < 0 {
		return false
	}

	if matrix[i-1][j] == 'M' {
		if matrix[i-2][j] == 'A' {
			if matrix[i-3][j] == 'S' {
				return true
			}
		}
	}

	return false
}

func part1(matrix []string) int {
	xmasCount := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			char := matrix[i][j]
			if char == 'X' {
				// up i-1, j
				if up(matrix, i, j) {
					xmasCount += 1
				}

				// down i+1, j
				if down(matrix, i, j) {
					xmasCount += 1
				}

				// left i, j-1
				if left(matrix, i, j) {
					xmasCount += 1
				}

				// right i, j+1
				if right(matrix, i, j) {
					xmasCount += 1
				}

				// diag-ul i-1, j-1
				if diagUl(matrix, i, j) {
					xmasCount += 1
				}

				// diag-ur i-1, j+1
				if diagUr(matrix, i, j) {
					xmasCount += 1
				}

				// diag-dl i+1, j-1
				if diagDl(matrix, i, j) {
					xmasCount += 1
				}

				// diag-dr i+1, j+1
				if diagDr(matrix, i, j) {
					xmasCount += 1
				}
			}
		}
	}
	return xmasCount
}

func part2(matrix []string) int {
	result := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			char := matrix[i][j]
			if char == 'A' {
				// diag1 ul dr
				if i-1 < 0 || j-1 < 0 {
					continue
				}
				if i+1 >= len(matrix) || j+1 >= len(matrix[i]) {
					continue
				}

				if matrix[i-1][j-1] != 'M' && matrix[i-1][j-1] != 'S' {
					continue
				}

				if matrix[i-1][j-1] == 'M' {
					if matrix[i+1][j+1] != 'S' {
						continue
					}
				}

				if matrix[i-1][j-1] == 'S' {
					if matrix[i+1][j+1] != 'M' {
						continue
					}
				}

				// diag1 OK

				// diag2 ur dl
				if i-1 < 0 || j+1 >= len(matrix[i]) {
					continue
				}
				if i+1 >= len(matrix) || j-1 < 0 {
					continue
				}

				if matrix[i-1][j+1] != 'M' && matrix[i-1][j+1] != 'S' {
					continue
				}

				if matrix[i-1][j+1] == 'M' {
					if matrix[i+1][j-1] != 'S' {
						continue
					}
				}

				if matrix[i-1][j+1] == 'S' {
					if matrix[i+1][j-1] != 'M' {
						continue
					}
				}

				// ok diag2
				result += 1
			}
		}
	}

	return result
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var matrix []string

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := part1(matrix)
	part2 := part2(matrix)

	fmt.Printf("PART 1: result is %d\n", part1)
	fmt.Printf("PART 2: result is %d\n", part2)
}
