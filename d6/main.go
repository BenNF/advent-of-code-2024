package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	up    = 0
	right = 1
	down  = 2
	left  = 3
)

type position struct {
	i int
	j int
}

type posDir struct {
	position
	dir int
}

func main() {
	fmt.Println(doP1())
	fmt.Println(doP2())
}

func doP1() int {
	field := parse()
	path, _ := tracePath(field, findStart(field))
	return len(uniquePositions(path))
}

func doP2() int {
	field := parse()
	startLocation := findStart(field)
	initPath, _ := tracePath(field, startLocation)
	var wg sync.WaitGroup
	var total atomic.Int64

	for _, pos := range uniquePositions(initPath) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if doesBlockCauseLoop(field, pos, startLocation) {
				total.Add(1)
			}
		}()
	}
	wg.Wait()
	return int(total.Load())
}

func uniquePositions(positions map[posDir]bool) []position {
	unique := make(map[position]bool)
	for v := range positions {
		unique[v.position] = true
	}
	keys := make([]position, len(unique))
	i := 0
	for k := range unique {
		keys[i] = k
		i++
	}
	return keys
}

func doesBlockCauseLoop(field [][]string, pos position, startPos position) bool {
	if field[pos.i][pos.j] == "^" {
		return false
	}

	//deep copy
	blockField := make([][]string, len(field))
	for i, _ := range field {
		blockField[i] = make([]string, len(field[i]))
		copy(blockField[i], field[i])
	}

	blockField[pos.i][pos.j] = "#"
	_, err := tracePath(blockField, startPos)
	if err != nil {
		//we loop
		return true
	}
	return false
}

func tracePath(field [][]string, startPos position) (map[posDir]bool, error) {
	visitedPos := map[posDir]bool{}
	pos := startPos
	dir := up
	cur := posDir{pos, dir}
	visitedPos[cur] = true
	for {
		nextPos := moveDirection(pos, dir)
		if nextPos.i < 0 || nextPos.i >= len(field) || nextPos.j < 0 || nextPos.j >= len(field[nextPos.i]) {
			//out of bounds, we exited
			break
		}
		if field[nextPos.i][nextPos.j] == "#" {
			dir = turn90(dir)
		} else {
			pos = nextPos
		}
		cur = posDir{pos, dir}
		if visitedPos[cur] {
			return nil, errors.New("loop detected")
		} else {
			visitedPos[cur] = true
		}
	}

	return visitedPos, nil
}

func findStart(field [][]string) position {
	for i, row := range field {
		for j, v := range row {
			if v == "^" {
				return position{i, j}
			}
		}
	}
	panic("No start found")
}

func moveDirection(pos position, dir int) position {
	switch dir {
	case up:
		return position{pos.i - 1, pos.j}
	case right:
		return position{pos.i, pos.j + 1}
	case down:
		return position{pos.i + 1, pos.j}

	case left:
		return position{pos.i, pos.j - 1}

	}
	panic("unknown direction")

}

func turn90(dir int) int {
	return (dir + 1) % 4
}

func parse() [][]string {
	absPath, err := filepath.Abs("d6/input.txt")
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
