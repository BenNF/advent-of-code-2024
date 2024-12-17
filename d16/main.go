package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	up    = 0
	right = 1
	down  = 2
	left  = 3
)

type point struct {
	i, j int
}

type move struct {
	p   point
	dir int
}

type frame struct {
	m          move
	score      int
	visitedMap map[point]bool
}

func main() {
	now := time.Now()
	minScore, bestPaths := findPaths()
	fmt.Println("p1: ", minScore, "in: ", time.Since(now))
	fmt.Println("p2: ", len(bestPaths[minScore])+1, "in: ", time.Since(now))
}
func findPaths() (int, map[int]map[point]bool) {
	field := parse()
	minScore := math.MaxInt
	s, e := findStartEnd(&field)

	minPositions := make(map[move]int)
	bestPaths := make(map[int]map[point]bool)

	queue := make([]frame, 0)
	queue = append(queue, frame{move{s, right}, 0, map[point]bool{}})
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		if c.m.p == e && c.score <= minScore {
			minScore = c.score
			if bestPaths[c.score] == nil {
				bestPaths[c.score] = make(map[point]bool, len(c.visitedMap))
			}
			for k := range c.visitedMap {
				bestPaths[c.score][k] = true
			}
			continue
		}

		if c.score > minScore {
			continue
		}

		if _, ok := c.visitedMap[c.m.p]; ok {
			continue
		}

		mScoreMin, ok := minPositions[c.m]
		if !ok {
			minPositions[c.m] = c.score
		} else if c.score > mScoreMin {
			continue
		} else {
			minPositions[c.m] = c.score
		}
		c.visitedMap[c.m.p] = true

		straight := move{moveDir(c.m.dir, c.m.p), c.m.dir}

		cwDir := turnClockWise(c.m.dir)
		cw := move{moveDir(cwDir, c.m.p), cwDir}

		ccwDir := turnCounterClockWise(c.m.dir)
		ccw := move{moveDir(ccwDir, c.m.p), ccwDir}

		if !isWall(straight.p, &field) {
			queue = append(queue, frame{straight, c.score + 1, cpMap(c.visitedMap)})
		}
		if !isWall(cw.p, &field) {
			queue = append(queue, frame{cw, c.score + 1001, cpMap(c.visitedMap)})
		}
		if !isWall(ccw.p, &field) {
			queue = append(queue, frame{ccw, c.score + 1001, cpMap(c.visitedMap)})
		}
	}
	return minScore, bestPaths
}

func isWall(p point, field *[][]string) bool {
	return (*field)[p.i][p.j] == "#"
}
func cpMap(m map[point]bool) map[point]bool {
	out := make(map[point]bool, len(m))
	for k := range m {
		out[k] = m[k]
	}
	return out
}

func moveDir(dir int, p point) point {
	switch dir {
	case up:
		return point{p.i - 1, p.j}
	case down:
		return point{p.i + 1, p.j}
	case left:
		return point{p.i, p.j - 1}
	case right:
		return point{p.i, p.j + 1}
	}
	panic("")
}

func turnClockWise(dir int) int {
	return (dir + 1) % 4
}

func turnCounterClockWise(dir int) int {
	move := dir - 1
	if move < 0 {
		move = left
	}
	return move
}

func findStartEnd(field *[][]string) (point, point) {
	var start, end point
	for i, r := range *field {
		for j, c := range r {
			if c == "S" {
				start = point{i, j}
			}
			if c == "E" {
				end = point{i, j}
			}

			if start != (point{}) && end != (point{}) {
				return start, end
			}
		}
	}
	panic("Cannot find start and end")
}
func parse() [][]string {
	absPath, err := filepath.Abs("d16/input.txt")
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
