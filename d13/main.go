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

type machine struct {
	ax, ay, bx, by, tx, ty int
}

func main() {
	now := time.Now()
	fmt.Println("p1: ", doP1(), "in: ", time.Since(now))
	now = time.Now()
	fmt.Println("p2: ", doP2(), "in: ", time.Since(now))
}

func doP1() int {
	total := 0
	for _, m := range parse() {
		total += solveMachine(m)
	}
	return total
}

func doP2() int {
	total := 0
	for _, m := range parse() {
		m.tx = m.tx + 10000000000000
		m.ty = m.ty + 10000000000000
		total += solveMachine(m)
	}
	return total

}

// manual algebra solving the linear equations produces this set of values
func solveMachine(m machine) int {
	div := m.ay*m.bx - m.by*m.ax
	a := (m.ty*m.bx - m.tx*m.by) / div
	b := (m.tx*m.ay - m.ty*m.ax) / div

	checkx := a*m.ax + b*m.bx
	checky := a*m.ay + b*m.by
	if checkx == m.tx && checky == m.ty {
		return 3*a + b
	}
	return 0
}
func parse() []machine {
	absPath, err := filepath.Abs("d13/input.txt")
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	buttonLineRegex := regexp.MustCompile("X\\+(\\d+), Y\\+(\\d+)")
	prizeLineRegex := regexp.MustCompile("X=(\\d+), Y=(\\d+)")

	output := make([]machine, 0)

	var ax, ay, bx, by, tx, ty int
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		line := scanner.Text()

		if strings.Contains(line, "Button A") {
			finds := buttonLineRegex.FindStringSubmatch(line)
			ax, _ = strconv.Atoi(finds[1])
			ay, _ = strconv.Atoi(finds[2])
		}
		if strings.Contains(line, "Button B") {
			finds := buttonLineRegex.FindStringSubmatch(line)
			bx, _ = strconv.Atoi(finds[1])
			by, _ = strconv.Atoi(finds[2])
		}

		if strings.Contains(line, "Prize") {
			finds := prizeLineRegex.FindStringSubmatch(line)
			tx, _ = strconv.Atoi(finds[1])
			ty, _ = strconv.Atoi(finds[2])

			output = append(output, machine{ax, ay, bx, by, tx, ty})
			ax, ay, bx, by, tx, ty = 0, 0, 0, 0, 0, 0
		}

	}
	return output
}
