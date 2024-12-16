package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

const up = "^"
const down = "v"
const left = "<"
const right = ">"

const push = 0
const move = 1
const stop = -1

type point struct {
	i, j int
}

func main() {
	now := time.Now()
	fmt.Println("P1 in: ", time.Since(now), doP1())
	now = time.Now()
	fmt.Println("P2 in: ", time.Since(now), doP2())

}

func doP1() int {
	field, instructions := parse()
	rPos := findStart(&field)

	for _, i := range instructions {
		//drawField(&field)
		rPos = moveRobot(&field, rPos, i, pushBox)
	}
	return scoreField(&field, "O")
}

func doP2() int {
	field, instructions := parse()
	field = scaleMap(field)
	rPos := findStart(&field)

	for _, i := range instructions {
		rPos = moveRobot(&field, rPos, i, pushWideBox)
		//drawField(&field)
	}
	return scoreField(&field, "[")
}

func scoreField(field *[][]string, target string) int {
	total := 0

	for i, r := range *field {
		for j, c := range r {
			if c == target {
				total += (100 * i) + j
			}
		}
	}
	return total
}

func moveRobot(field *[][]string, p point, dir string, pusher func(field *[][]string, p point, dir string) point) point {
	dp := doesPush(field, p, dir)
	switch dp {

	case stop:
		return p
	case move:
		(*field)[p.i][p.j] = "."
		p = moveDir(dir, p)
		(*field)[p.i][p.j] = "@"
		return p
	case push:
		return pusher(field, p, dir)
	}
	return p
}
func pushWideBox(field *[][]string, p point, dir string) point {
	if dir == left || dir == right {
		return pushWideBoxHorizontal(field, p, dir)
	} else {
		return pushWideBoxVertical(field, p, dir)
	}

}

func pushWideBoxVertical(field *[][]string, p point, dir string) point {
	boxesLeft := make([]point, 0)
	findBoxes(&boxesLeft, field, p, dir)
	movedInto := make([]point, 0)
	for _, lp := range boxesLeft {
		rp := moveDir(right, lp)

		if !slices.Contains(movedInto, lp) {
			(*field)[lp.i][lp.j] = "."

		}
		if !slices.Contains(movedInto, rp) {
			(*field)[rp.i][rp.j] = "."
		}

		lp = moveDir(dir, lp)
		rp = moveDir(dir, rp)

		(*field)[rp.i][rp.j] = "]"
		(*field)[lp.i][lp.j] = "["
		movedInto = append(movedInto, lp, rp)

	}

	rPos := moveDir(dir, p)
	(*field)[rPos.i][rPos.j] = "@"
	(*field)[p.i][p.j] = "."
	return rPos
}

func findBoxes(boxesLeft *[]point, field *[][]string, p point, dir string) {
	p = moveDir(dir, p)
	c := (*field)[p.i][p.j]
	if c == "." {
		return
	}
	if c == "#" {
		return
	}
	if c == "[" {
		*boxesLeft = append(*boxesLeft, p)
		findBoxes(boxesLeft, field, moveDir(right, p), dir)
	}
	if c == "]" {
		*boxesLeft = append(*boxesLeft, moveDir(left, p))
		findBoxes(boxesLeft, field, moveDir(left, p), dir)
	}
	findBoxes(boxesLeft, field, p, dir)
}

func pushWideBoxHorizontal(field *[][]string, p point, dir string) point {
	(*field)[p.i][p.j] = "."
	p = moveDir(dir, p)
	(*field)[p.i][p.j] = "@"
	rPos := p
	for {
		p = moveDir(dir, p)
		c := (*field)[p.i][p.j]
		if c == "[" {
			(*field)[p.i][p.j] = "]"
		}
		if c == "]" {
			(*field)[p.i][p.j] = "["
		}

		if (*field)[p.i][p.j] == "." {
			if dir == left {
				(*field)[p.i][p.j] = "["
			}
			if dir == right {
				(*field)[p.i][p.j] = "]"
			}

			return rPos
		}
	}
}

func pushBox(field *[][]string, p point, dir string) point {
	(*field)[p.i][p.j] = "."
	p = moveDir(dir, p)
	(*field)[p.i][p.j] = "@"
	rPos := p
	for {
		p = moveDir(dir, p)
		if (*field)[p.i][p.j] == "." {
			(*field)[p.i][p.j] = "O"
			return rPos
		}
	}
}

func doesPush(field *[][]string, p point, dir string) int {
	doPush := false
	for {
		p = moveDir(dir, p)
		c := (*field)[p.i][p.j]
		if c == "#" {
			return stop
		}
		if c == "O" || c == "[" || c == "]" {
			doPush = true
			if dir == up || dir == down {
				if c == "[" && doesPush(field, moveDir(right, p), dir) == stop {
					return stop
				}
				if c == "]" && doesPush(field, moveDir(left, p), dir) == stop {
					return stop
				}
			}
			continue
		}
		if c == "[" {
			doPush = true
			if doesPush(field, moveDir(right, p), dir) == stop {
				return stop
			}
			continue
		}
		if c == "]" {
			doPush = true
			if doesPush(field, moveDir(left, p), dir) == stop {
				return stop
			}
			continue
		}

		if (*field)[p.i][p.j] == "." {
			if doPush {
				return push
			}
			return move
		}
	}
}

func moveDir(dir string, p point) point {
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

func findStart(field *[][]string) point {
	for i, row := range *field {
		for j, v := range row {
			if v == "@" {
				return point{i, j}
			}
		}
	}
	panic("")
}

func scaleMap(field [][]string) [][]string {
	wideField := make([][]string, len(field))
	for i, r := range field {
		for _, c := range r {
			switch c {

			case "#":
				wideField[i] = append(wideField[i], "#", "#")
			case "O":
				wideField[i] = append(wideField[i], "[", "]")
			case ".":
				wideField[i] = append(wideField[i], ".", ".")
			case "@":
				wideField[i] = append(wideField[i], "@", ".")
			}
		}
	}
	return wideField
}

func parse() ([][]string, []string) {
	absPath, err := filepath.Abs("d15/input.txt")
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	field := make([][]string, 0)
	instructions := make([]string, 0)
	var isField = true
	for scanner.Scan() {
		if (scanner.Text()) == "" {
			isField = false
			continue
		}
		split := strings.Split(scanner.Text(), "")
		if isField {
			field = append(field, split)
		} else {
			for _, v := range split {
				instructions = append(instructions, v)
			}
		}

	}

	return field, instructions
}
