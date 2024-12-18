package main

import (
	"aoc-2024/d18/astar"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const width = 71

type World map[int]map[int]*tile

type point struct {
	x, y int
}

type posPoint struct {
	p point
	i int
}

type tile struct {
	x, y     int
	blockers map[point]bool
	w        World
}

func (t *tile) PathNeighbors() []astar.Pather {
	neighbors := make([]astar.Pather, 0)

	up := point{t.x, t.y + 1}
	down := point{t.x, t.y - 1}
	left := point{t.x - 1, t.y}
	right := point{t.x + 1, t.y}

	if _, ok := t.blockers[up]; !ok && isInBounds(up) {
		neighbors = append(neighbors, t.w[up.x][up.y])
	}
	if _, ok := t.blockers[down]; !ok && isInBounds(down) {
		neighbors = append(neighbors, t.w[down.x][down.y])
	}
	if _, ok := t.blockers[left]; !ok && isInBounds(left) {
		neighbors = append(neighbors, t.w[left.x][left.y])
	}
	if _, ok := t.blockers[right]; !ok && isInBounds(right) {
		neighbors = append(neighbors, t.w[right.x][right.y])
	}
	return neighbors
}
func (t *tile) PathNeighborCost(_ astar.Pather) float64 {
	return 1
}

func (t *tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*tile)
	absX := toT.x - t.x
	if absX < 0 {
		absX = -absX
	}
	absY := toT.y - t.y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

func isInBounds(p point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < width && p.y < width
}

func main() {
	now := time.Now()
	fmt.Println("p1: ", doP1(), "in: ", time.Since(now))
	now = time.Now()
	fmt.Println("p2: ", doP2(), "in: ", time.Since(now))
}

func doP1() int {
	blocked, field := buildBlockerWorld(parse()[:1024])

	start := tile{0, 0, blocked, field}
	end := tile{width - 1, width - 1, blocked, field}

	_, dist, _ := astar.Path(&start, &end)
	return int(dist)
}

func doP2() point {
	blockers := parse()
	output := make(chan posPoint, len(blockers)-1024)
	wg := sync.WaitGroup{}

	for i := 1024; i < len(blockers); i++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			blocked, field := buildBlockerWorld(parse()[:p])
			start := tile{0, 0, blocked, field}
			end := tile{width - 1, width - 1, blocked, field}
			_, _, found := astar.Path(&start, &end)
			if !found {
				output <- posPoint{blockers[p-1], i}
			}
		}(i)
	}
	wg.Wait()
	close(output)
	minp := posPoint{}
	for v := range output {
		if minp.i == 0 {
			minp = v
			continue
		}
		if v.i < minp.i {
			minp = v
		}
	}
	return minp.p
}

func buildBlockerWorld(blockers []point) (map[point]bool, World) {
	blockMap := make(map[point]bool, len(blockers))

	for _, k := range blockers {
		blockMap[k] = true
	}
	field := World{}
	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			if field[x] == nil {
				field[x] = make(map[int]*tile)
			}
			field[x][y] = &tile{x, y, blockMap, field}
		}
	}
	return blockMap, field
}

func parse() []point {
	absPath, err := filepath.Abs("d18/input.txt")
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	blockers := make([]point, 0)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		p := point{x, y}
		blockers = append(blockers, p)
	}

	return blockers
}
