package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
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
	return score(computeRobots(100, &robots))
}

func doP2() int {
	init := parse()
	count := len(init)
	wg := sync.WaitGroup{}
	output := make(chan int, 1)
	stop := make(chan bool)
	for step := range 100000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-stop:
				return
			default:
				robots := computeRobots(step, &init)
				if len(robots) == count {
					output <- step
					close(stop)
				}
			}

		}()
	}
	wg.Wait()
	return <-output
}

func computeRobots(steps int, initialRobots *[]robot) map[point]int {
	outRobots := make(map[point]int, len(*initialRobots))
	for _, r := range *initialRobots {
		px := (r.px + (r.vx * steps)) % width
		py := (r.py + (r.vy * steps)) % height
		if px < 0 {
			px = width + px
		}
		if py < 0 {
			py = height + py
		}
		outRobots[point{px, py}]++
	}
	return outRobots
}

func score(robots map[point]int) int {
	var q1, q2, q3, q4 int
	horizontal := (width - 1) / 2
	vertical := (height - 1) / 2

	for p := range robots {
		if p.x < horizontal && p.y < vertical {
			q1 += robots[p]
		}
		if p.x > horizontal && p.y < vertical {
			q2 += robots[p]
		}
		if (p.x < horizontal) && p.y > vertical {
			q3 += robots[p]
		}
		if (p.x > horizontal) && p.y > vertical {
			q4 += robots[p]
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
