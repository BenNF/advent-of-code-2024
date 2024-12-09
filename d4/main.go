package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(doP2())
}

func doP1() int {
	total := 0

	input := parse()
	for i, row := range input {
		for j, char := range row {
			if char != "X" {
				continue
			}
			fmt.Println(i, j)
			//up path
			if i > 2 && input[i-1][j] == "M" && input[i-2][j] == "A" && input[i-3][j] == "S" {
				total += 1
				fmt.Println("up")
			}
			//left
			if j > 2 && input[i][j-1] == "M" && input[i][j-2] == "A" && input[i][j-3] == "S" {
				total += 1
				fmt.Println("left")
			}
			//right
			if j < len(input[i])-3 && input[i][j+1] == "M" && input[i][j+2] == "A" && input[i][j+3] == "S" {
				total += 1
				fmt.Println("right")
			}
			//down
			if i < len(input)-3 && input[i+1][j] == "M" && input[i+2][j] == "A" && input[i+3][j] == "S" {
				total += 1
				fmt.Println("down")
			}
			//up left
			if i > 2 && j > 2 && input[i-1][j-1] == "M" && input[i-2][j-2] == "A" && input[i-3][j-3] == "S" {
				total += 1
				fmt.Println(" up left")
			}
			//up right
			if i > 2 && j < len(input[i])-3 && input[i-1][j+1] == "M" && input[i-2][j+2] == "A" && input[i-3][j+3] == "S" {
				total += 1
				fmt.Println("up right")
			}
			//down left
			if j > 2 && i < len(input)-3 && input[i+1][j-1] == "M" && input[i+2][j-2] == "A" && input[i+3][j-3] == "S" {
				total += 1
				fmt.Println(" down left")
			}
			//down right
			if i < len(input)-3 && j < len(input[i])-3 && input[i+1][j+1] == "M" && input[i+2][j+2] == "A" && input[i+3][j+3] == "S" {
				total += 1
				fmt.Println(" down right")
			}

		}
	}
	return total
}

func doP2() int {

	centers := make(map[string]bool)

	input := parse()
	for i, row := range input {
		for j, char := range row {
			if char != "M" {
				continue
			}
			//up left
			if i > 1 && j > 1 && input[i-1][j-1] == "A" && input[i-2][j-2] == "S" {
				crossUpRight := &input[i-2][j]
				crossDownLeft := &input[i][j-2]

				if (*crossDownLeft == "M" && *crossDownLeft == "S") || (*crossDownLeft == "S" && *crossUpRight == "M") {
					centers[fmt.Sprintf("%d,%d", i-1, j-1)] = true
				}
			}
			//up right
			if i > 1 && j < len(input[i])-2 && input[i-1][j+1] == "A" && input[i-2][j+2] == "S" {
				crossUpLeft := &input[i-2][j]
				crossDownRight := &input[i][j+2]

				if (*crossUpLeft == "M" && *crossDownRight == "S") || (*crossUpLeft == "S" && *crossDownRight == "M") {
					centers[fmt.Sprintf("%d,%d", i-1, j+1)] = true
				}
			}
			//down left
			if j > 1 && i < len(input)-2 && input[i+1][j-1] == "A" && input[i+2][j-2] == "S" {
				crossUpLeft := &input[i][j-2]
				crossDownRight := &input[i+2][j]

				if (*crossUpLeft == "M" && *crossDownRight == "S") || (*crossUpLeft == "S" && *crossDownRight == "M") {
					centers[fmt.Sprintf("%d,%d", i+1, j-1)] = true
				}
			}
			//down right
			if i < len(input)-2 && j < len(input[i])-2 && input[i+1][j+1] == "A" && input[i+2][j+2] == "S" {
				crossUpRight := &input[i][j+2]
				crossDownLeft := &input[i+2][j]

				if (*crossUpRight == "M" && *crossDownLeft == "S") || (*crossUpRight == "S" && *crossDownLeft == "M") {
					centers[fmt.Sprintf("%d,%d", i+1, j+1)] = true
				}
			}

		}
	}

	return len(centers)
}

func parse() [][]string {
	absPath, err := filepath.Abs("d4/input.txt")

	if err != nil {
		panic(err)

	}

	data, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	output := make([][]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		output = append(output, chars)
	}
	return output
}
