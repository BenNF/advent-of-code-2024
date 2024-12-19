package main

import (
	"bufio"
	"fmt"
	"github.com/orcaman/concurrent-map/v2"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	now := time.Now()
	p1, p2 := doP1P2()
	fmt.Println("p1: ", p1, "in: ", time.Since(now))
	fmt.Println("p2: ", p2, "in: ", time.Since(now))
}

func doP1P2() (uint64, uint64) {
	towels, patterns := parse()
	cache := cmap.New[int]()
	var totalPassing atomic.Uint64
	var total atomic.Uint64
	wg := sync.WaitGroup{}

	for _, p := range patterns {
		wg.Add(1)
		func() {
			defer wg.Done()
			count := findCombo(&towels, &cache, p)
			if count > 0 {
				total.Add(1)
			}
			totalPassing.Add(uint64(count))
		}()
	}
	wg.Wait()
	return total.Load(), totalPassing.Load()
}

func findCombo(towels *map[string]bool, cache *cmap.ConcurrentMap[string, int], pattern string) int {
	if v, ok := cache.Get(pattern); ok {
		return v
	}
	total := 0
	defer func() {
		cache.Set(pattern, total)
	}()

	if _, ok := (*towels)[pattern]; ok {
		total++
	}
	for i := range len(pattern) {
		_, ok := (*towels)[pattern[:i]] //front
		if ok {
			total += findCombo(towels, cache, pattern[i:]) //back
		}
	}
	return total
}

func parse() (map[string]bool, []string) {
	absPath, err := filepath.Abs("d19/input.txt")
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	scanner.Scan()
	towels := strings.Split(scanner.Text(), ", ")
	towelMap := make(map[string]bool, len(towels))
	for _, t := range towels {
		towelMap[t] = true
	}
	patterns := make([]string, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		patterns = append(patterns, scanner.Text())
	}
	return towelMap, patterns
}
