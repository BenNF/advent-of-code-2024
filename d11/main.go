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
	for _, v := range stonesList {
		stones[v] += 1
	}

	for i := 0; i < 75; i++ {
		stones = cycleStonesMap(stones)
	}

	total := 0
	for k := range stones {
		total += stones[k]
	}
	return total

}

func cycleStonesMap(stones map[int]int) map[int]int {
	update := make(map[int]int)

	for k := range stones {
		if k == 0 {
			update[1] += stones[k]
		} else if l := lenLoop(k); l%2 == 0 {
			front, back := splitInt(k, l)
			update[front] += stones[k]
			update[back] += stones[k]
		} else {
			update[k*2024] += stones[k]
		}
	}
	return update
}

func cycleStones(stones []int) []int {
	update := make([]int, 0)

	for _, v := range stones {
		if v == 0 {
			update = append(update, 1)
		} else if l := lenLoop(v); l%2 == 0 {
			front, back := splitInt(v, l)
			update = append(update, front, back)
		} else {
			update = append(update, v*2024)
		}
	}
	return update
}
func splitInt(n, l int) (int, int) {
	slc := make([]int, l)
	i := 0
	for n > 0 {
		slc[i] = n % 10
		n /= 10
		i++
	}
	f, b := 0, 0
	for i, v := range slc[:len(slc)/2] {
		b += int(math.Pow10(i)) * v
	}
	for i, v := range slc[len(slc)/2:] {
		f += int(math.Pow10(i)) * v
	}

	return f, b
}

func lenLoop(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
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
