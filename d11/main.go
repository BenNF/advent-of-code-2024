package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(doP1(), "in: ", time.Since(now))
	now = time.Now()
	fmt.Println(doP2(), "in: ", time.Since(now))

}

func doP1() int {
	stones := parse()

	for i := 0; i < 25; i++ {
		stones = cycleStones(stones)
	}
	return len(stones)
}

func doP2() int {
	stonesList := parse()
	stones := make(map[int]int)
	splitCache := make(map[int][]int)
	lenCache := make(map[int]int)
	mulCache := make(map[int]int)
	for _, v := range stonesList {
		stones[v] += 1
	}
	for i := 0; i < 75; i++ {
		stones = cycleStonesMap(stones, &splitCache, &lenCache, &mulCache)
	}

	total := 0
	for k := range stones {
		total += stones[k]
	}
	return total

}

func cycleStonesMap(stones map[int]int, splitCache *map[int][]int, lenCache, mulCache *map[int]int) map[int]int {
	update := make(map[int]int, len(stones))

	for k := range stones {
		if k == 0 {
			update[1] += stones[k]
			continue
		}
		if _, ok := (*lenCache)[k]; ok {
			split := (*splitCache)[k]
			update[split[0]] += stones[k]
			update[split[1]] += stones[k]
		} else if l := lenInt(k); l%2 == 0 {
			front, back := splitInt(k, l)
			update[front] += stones[k]
			update[back] += stones[k]
			(*lenCache)[k] = l
			(*splitCache)[k] = []int{front, back}
		} else {
			if val, ok := (*mulCache)[k]; ok {
				update[val] += stones[k]
			} else {
				val := k * 2024
				update[val] += stones[k]
				(*mulCache)[k] = val
			}
		}
	}
	return update
}

func cycleStones(stones []int) []int {
	update := make([]int, 0)

	for _, v := range stones {
		if v == 0 {
			update = append(update, 1)
		} else if l := lenInt(v); l%2 == 0 {
			front, back := splitInt(v, l)
			update = append(update, front, back)
		} else {
			update = append(update, v*2024)
		}
	}
	return update
}
func splitInt(n, l int) (int, int) {
	return n % int(math.Pow10(l/2)), n / int(math.Pow10(l/2))
}

func lenInt(i int) int {
	return int(math.Log10(float64(i)) + 1)
}

func parse() []int {
	absPath, err := filepath.Abs("d11/input.txt")
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	chars := strings.Fields(string(data))
	output := make([]int, len(chars))

	for i, c := range chars {
		output[i], _ = strconv.Atoi(c)
	}
	return output
}
