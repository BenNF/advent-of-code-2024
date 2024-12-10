package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type point struct {
	i, j int
}

func main() {
	fmt.Println(doP1())
	fmt.Println(doP2())

}

func doP1() int {
	field := parse()

	total := 0
	for i, r := range field {
		for j, v := range r {
			if v == 0 {
				t := make(map[point]bool)
				countUniqueDestinations(point{i, j}, &field, v-1, &t)
				total += len(t)
			}
		}
	}

	return total
}

func doP2() int {
	field := parse()

	total := 0
	for i, r := range field {
		for j, v := range r {
			if v == 0 {
				total += countPaths(point{i, j}, &field, v-1)
			}
		}
	}

	return total

}

func countPaths(p point, field *[][]int, prev int) int {
	if !isInBounds(p, field) {
		return 0
	}
	v := (*field)[p.i][p.j]
	if v != prev+1 {
		return 0
	}
	if v == 9 {
		return 1
	}

	return countPaths(point{p.i - 1, p.j}, field, v) +
		countPaths(point{p.i, p.j - 1}, field, v) +
		countPaths(point{p.i + 1, p.j}, field, v) +
		countPaths(point{p.i, p.j + 1}, field, v)
}

func countUniqueDestinations(p point, field *[][]int, prev int, tracker *map[point]bool) {
	if !isInBounds(p, field) {
		return
	}

	if (*tracker)[p] == true {
		return
	}

	v := (*field)[p.i][p.j]

	if v != prev+1 {
		return
	}
	if v == 9 {
		(*tracker)[p] = true
	}

	countUniqueDestinations(point{p.i - 1, p.j}, field, v, tracker)
	countUniqueDestinations(point{p.i, p.j - 1}, field, v, tracker)
	countUniqueDestinations(point{p.i + 1, p.j}, field, v, tracker)
	countUniqueDestinations(point{p.i, p.j + 1}, field, v, tracker)
}

func isInBounds(p point, field *[][]int) bool {
	return p.i >= 0 && p.i < len(*field) && p.j >= 0 && p.j < len((*field)[p.i])
}
func parse() [][]int {
	absPath, err := filepath.Abs("d10/input.txt")
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	output := make([][]int, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		chars := strings.Split(scanner.Text(), "")
		parsed := make([]int, len(chars))
		for i, v := range chars {
			parsed[i], _ = strconv.Atoi(v)
		}
		output = append(output, parsed)
	}
	return output
}
