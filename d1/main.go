package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(doP2())
}

func doP2() int {
	left, right := parsLists()
	rightCount := make(map[int]int)

	for _, i := range right {
		rightCount[i]++
	}
	fmt.Println(rightCount)

	total := 0
	for _, v := range left {
		if rightCount[v] > 0 {
			total += v * rightCount[v]
		}
	}

	return total
}

func doP1() int {
	left, right := parsLists()

	sort.Ints(left)
	sort.Ints(right)

	fmt.Println(left)
	fmt.Println(right)
	var total int
	for i := 0; i < len(left); i++ {
		fmt.Println(left[i], right[i])
		distance := left[i] - right[i]
		if distance < 0 {
			distance = distance * -1
		}
		total += distance
		fmt.Println(distance)
		fmt.Println(total)
	}

	return total
}

func parsLists() ([]int, []int) {
	absPath, err := filepath.Abs("d1/p1-input.txt")

	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(absPath)

	if err != nil {
		panic(err)
	}

	var left []int
	var right []int

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		split := strings.Fields(scanner.Text())
		fmt.Println(split)
		i, err := strconv.Atoi(strings.TrimSpace(split[0]))
		if err != nil {
			panic(err)
		}
		j, err := strconv.Atoi(strings.TrimSpace(split[1]))
		if err != nil {
			panic(err)
		}

		left = append(left, i)
		right = append(right, j)
	}
	return left, right
}
