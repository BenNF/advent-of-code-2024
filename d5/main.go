package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(doP2())
}

func doP1() int {
	beforeOrderRules, afterOrderRules, pageUpdates := parse()
	total := 0

	fmt.Println("before rules", beforeOrderRules)
	fmt.Println("after rules", afterOrderRules)

	for _, pages := range pageUpdates {
		if isValid(pages, beforeOrderRules, afterOrderRules) {
			total += pages[(len(pages)-1)/2]
		}
	}

	return total
}

func doP2() int {
	beforeOrderRules, afterOrderRules, pageUpdates := parse()
	total := 0

	for _, pages := range pageUpdates {
		if !isValid(pages, beforeOrderRules, afterOrderRules) {
			total += sortPages(pages, beforeOrderRules)[(len(pages)-1)/2]
		}
	}

	return total
}

func sortPages(pages []int, beforeRules map[int][]int) []int {
	slices.SortFunc(pages, func(a int, b int) int {
		if a == b {
			return 0
		}
		for _, v := range beforeRules[a] {
			if v == b {
				return -1
			}
		}
		return 1
	})
	return pages
}

func isValid(pages []int, beforeRules map[int][]int, afterRules map[int][]int) bool {
	for i, p := range pages {
		beforePages := pages[:i]
		afterPages := pages[i+1:]
		requiredBeforeRules := beforeRules[p]
		requiredAfterRules := afterRules[p]

		for _, v := range beforePages {
			for _, r := range requiredBeforeRules {
				if v == r {
					return false
				}
			}
		}
		for _, v := range afterPages {
			for _, r := range requiredAfterRules {
				if v == r {
					return false
				}
			}
		}
	}
	return true
}

func parse() (map[int][]int, map[int][]int, [][]int) {
	absPath, err := filepath.Abs("d5/input.txt")

	if err != nil {
		panic(err)

	}

	data, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	orderingRules := make([][]int, 0)
	beforeOrderRules := make(map[int][]int)
	afterOrderRules := make(map[int][]int)
	pageUpdates := make([][]int, 0)
	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		chars := strings.Split(scanner.Text(), "|")
		intParse := make([]int, len(chars))
		for i, v := range chars {
			intParse[i], _ = strconv.Atoi(v)
		}
		orderingRules = append(orderingRules, intParse)
	}

	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), ",")
		intParse := make([]int, len(chars))
		for i, v := range chars {
			intParse[i], _ = strconv.Atoi(v)
		}
		pageUpdates = append(pageUpdates, intParse)
	}

	for _, v := range orderingRules {
		beforeOrderRules[v[0]] = append(beforeOrderRules[v[0]], v[1])
		afterOrderRules[v[1]] = append(afterOrderRules[v[1]], v[0])
	}
	return beforeOrderRules, afterOrderRules, pageUpdates
}
