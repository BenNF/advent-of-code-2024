package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
	lookup := buildLookup(field)
	nodes := make(map[point]bool)

	loopKeys(lookup, func(p point, t point, vec point) {
		p1 := plusPoint(p, vec)
		p2 := minusPoint(t, vec)
		if isInBounds(p1, field) {
			nodes[p1] = true
		}
		if isInBounds(p2, field) {
			nodes[p2] = true
		}
	})
	return len(nodes)
}

func doP2() int {
	field := parse()
	lookup := buildLookup(field)
	nodes := make(map[point]bool)

	loopKeys(lookup, func(p point, t point, vec point) {
		nodes[t] = true
		nodes[p] = true
		tempP := plusPoint(p, vec)
		for isInBounds(tempP, field) {
			nodes[tempP] = true
			tempP = plusPoint(tempP, vec)
		}
		tempP = minusPoint(t, vec)
		for isInBounds(tempP, field) {
			nodes[tempP] = true
			tempP = minusPoint(tempP, vec)
		}
	})
	return len(nodes)
}

func loopKeys(lookup map[string][]point, f func(p point, t point, vec point)) {
	for k := range lookup {
		for i, p := range lookup[k] {
			rest := append(lookup[k][:i], lookup[k][i+1:]...)
			for _, t := range rest {
				vec := point{p.i - t.i, p.j - t.j}
				if vec.j == 0 && vec.i == 0 {
					continue
				}
				f(p, t, vec)
			}
		}
	}
}
func minusPoint(p point, v point) point {
	return point{p.i - v.i, p.j - v.j}
}

func plusPoint(p point, v point) point {
	return point{p.i + v.i, p.j + v.j}
}

func buildLookup(field [][]string) map[string][]point {
	lookup := make(map[string][]point)
	for i, r := range field {
		for j, c := range r {
			if c == "." {
				continue
			}
			lookup[c] = append(lookup[c], point{i, j})
		}
	}
	return lookup
}

func isInBounds(p point, field [][]string) bool {
	return p.i >= 0 && p.i < len(field) && p.j >= 0 && p.j < len(field[p.i])
}

func parse() [][]string {
	absPath, err := filepath.Abs("d8/input.txt")
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	output := make([][]string, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		chars := strings.Split(scanner.Text(), "")
		output = append(output, chars)
	}
	return output
}
