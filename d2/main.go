package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(doP2())
}

func doP1() int {
	in := parse()

	total := 0
	for _, v := range in {
		if isValid(computeDifs(v)) {
			total++
		}
	}
	return total
}

func allRightSize(input []int) bool {
	for _, v := range input {
		if v < 0 {
			v = v * -1
		}
		if v > 3 {
			return false
		}
		if v == 0 {
			return false
		}
	}
	return true
}

func allPos(input []int) bool {
	for _, v := range input {
		if v < 0 {
			return false
		}
	}
	return true
}

func allNeg(input []int) bool {
	for _, v := range input {
		if v > 0 {
			return false
		}
	}
	return true
}

func doP2() int {
	in := parse()

	total := 0
	for _, v := range in {
		difs := computeDifs(v)
		if isValid(difs) {
			total++
		} else if checkOneRemove(v) {
			total++
		}
	}
	return total
}

func computeDifs(input []int) []int {
	difs := make([]int, len(input)-1)
	for i := 1; i < len(input); i++ {
		difs[i-1] = input[i] - input[i-1]
	}
	return difs
}

func isValid(input []int) bool {
	return (allPos(input) || allNeg(input)) && allRightSize(input)
}

func checkOneRemove(input []int) bool {
	for i := 0; i < len(input); i++ {
		inputCopy := make([]int, len(input))
		copy(inputCopy, input)
		minusOne := remove(inputCopy, i)
		fmt.Println(i)
		fmt.Println(input, minusOne)
		if isValid(computeDifs(minusOne)) {
			fmt.Println(input, i, minusOne)
			return true
		}
	}

	return false
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
func parse() [][]int {
	absPath, err := filepath.Abs("d2/input.txt")

	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	output := make([][]int, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		lineList := make([]int, len(split))
		for i, v := range split {
			lineList[i], err = strconv.Atoi(v)
		}
		output = append(output, lineList)
	}
	return output
}
