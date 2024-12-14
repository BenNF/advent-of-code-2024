package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y int
}

type robot struct {
	px, py, vx, vy int
}

const width = 101
const height = 103

func main() {
	now := time.Now()
	fmt.Println("p1: ", doP1(), "in: ", time.Since(now))
	now = time.Now()
	fmt.Println("p2: ", doP2(), "in: ", time.Since(now))
}

func doP1() int {
	robots := parse()
	return score(computeRobots(100, width, height, &robots), height, width)
}

func doP2() int {
	init := parse()
	for step := range 100000 {
		robots := computeRobots(step, width, height, &init)
		if lookLikeChristmasTree(&robots) {
			return step
		}
	}
	return 0
}

func lookLikeChristmasTree(robots *[]robot) bool {
	unique := make(map[point]bool, len(*robots))
	for _, r := range *robots {
		p := point{r.px, r.py}
		if _, ok := unique[p]; ok {
			return false
		}
		unique[p] = true
	}

	return true

}

// debug functions for manual checking
func drawField(field [][]int, step int) {
	fmt.Println(step, ":", strings.Repeat("-", width*2))
	for _, r := range field {
		fmt.Println(r)
	}
}

func mapField(width, height int, robots *[]robot) [][]int {
	field := make([][]int, height)
	for i, _ := range field {
		field[i] = make([]int, width)
	}

	for _, r := range *robots {
		field[r.py][r.px] = 1
	}
	return field
}

func computeRobots(steps, width, height int, initalRobots *[]robot) []robot {
	outRobots := make([]robot, len(*initalRobots))
	for i, r := range *initalRobots {
		px := (r.px + (r.vx * steps)) % width
		py := (r.py + (r.vy * steps)) % height

		if px < 0 {
			px = width + px
		}

		if py < 0 {
			py = height + py
		}

		outRobots[i].px = px
		outRobots[i].py = py
	}
	return outRobots
}

func score(robots []robot, height, width int) int {
	var q1, q2, q3, q4 int
	horizontal := (width - 1) / 2
	vertical := (height - 1) / 2

	for _, r := range robots {
		if r.px < horizontal && r.py < vertical {
			q1++
		}
		if r.px > horizontal && r.py < vertical {
			q2++
		}
		if (r.px < horizontal) && r.py > vertical {
			q3++
		}
		if (r.px > horizontal) && r.py > vertical {
			q4++
		}
	}
	return q1 * q2 * q3 * q4
}

func parse() []robot {
	absPath, err := filepath.Abs("d14/input.txt")
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	lineMatcher := regexp.MustCompile("p=(-?\\d+),(-?\\d+) v=(-?\\d+),(-?\\d+)")
	output := make([]robot, 0)

	for scanner.Scan() {
		matches := lineMatcher.FindStringSubmatch(scanner.Text())
		px, _ := strconv.Atoi(matches[1])
		py, _ := strconv.Atoi(matches[2])
		vx, _ := strconv.Atoi(matches[3])
		vy, _ := strconv.Atoi(matches[4])

		r := robot{px, py, vx, vy}
		output = append(output, r)
	}
	return output
}
