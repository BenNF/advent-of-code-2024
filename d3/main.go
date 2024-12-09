package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(doP2())
}

func doP2() int {
	expr := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)|don't\\(\\)|do\\(\\)")

	input := getInput()

	finds := expr.FindAllStringSubmatch(input, -1)
	total := 0

	isActive := true
	for _, match := range finds {

		fmt.Println(match)
		if strings.Contains(match[0], "mul") && isActive {
			v1, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			v2, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}
			total += v1 * v2
		}

		if strings.TrimSpace(match[0]) == "do()" {
			isActive = true
		}
		if strings.TrimSpace(match[0]) == "don't()" {
			isActive = false
		}
	}
	return total
}

func doP1() int {
	expr := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")

	input := getInput()

	finds := expr.FindAllStringSubmatch(input, -1)
	total := 0

	for _, match := range finds {
		v1, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		v2, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		total += v1 * v2
	}
	return total
}

func getInput() string {
	absPath, err := filepath.Abs("d3/input.txt")

	if err != nil {
		panic(err)

	}

	data, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}
	return string(data)
}
