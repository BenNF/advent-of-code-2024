package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type point struct {
	i, j int
}

type region struct {
	perimeter int
	contained map[point]bool
}

func main() {
	now := time.Now()
	fmt.Println("p1: ", doP1(), "in: ", time.Since(now))
	now = time.Now()
	fmt.Println("p2: ", doP2(), "in: ", time.Since(now))
}

func doP1() int {
	regions := findRegions()
	return scorePerimeter(regions)
}

func doP2() int {
	regions := findRegions()
	return scoreSides(regions)

}

func findRegions() []region {
	field := parse()
	regions := make([]region, 0)
	globalContained := make(map[point]bool)
	for i, v := range field {
		for j, c := range v {
			p := point{i, j}
			if _, ok := globalContained[p]; ok {
				continue
			}
			prev := make(map[point]bool)
			perimeter := findRegion(p, c, &field, &prev)
			regions = append(regions, region{perimeter, prev})
			for k, v := range prev {
				globalContained[k] = v
			}

		}
	}
	return regions
}

func scoreSides(regions []region) int {
	total := 0
	for _, r := range regions {
		total += len(r.contained) * countSides(r)
	}
	return total
}

// actually counting corners
func countSides(r region) int {
	sides := 0
	for p := range r.contained {
		//outer corner
		if !mapContains(&r.contained, point{p.i - 1, p.j}) && !mapContains(&r.contained, point{p.i, p.j - 1}) {
			sides++
		}
		if !mapContains(&r.contained, point{p.i + 1, p.j}) && !mapContains(&r.contained, point{p.i, p.j - 1}) {
			sides++
		}
		if !mapContains(&r.contained, point{p.i - 1, p.j}) && !mapContains(&r.contained, point{p.i, p.j + 1}) {
			sides++
		}
		if !mapContains(&r.contained, point{p.i + 1, p.j}) && !mapContains(&r.contained, point{p.i, p.j + 1}) {
			sides++
		}
		//Inner corners
		if mapContains(&r.contained, point{p.i - 1, p.j}) && mapContains(&r.contained, point{p.i, p.j - 1}) && !mapContains(&r.contained, point{p.i - 1, p.j - 1}) {
			sides++
		}
		if mapContains(&r.contained, point{p.i + 1, p.j}) && mapContains(&r.contained, point{p.i, p.j - 1}) && !mapContains(&r.contained, point{p.i + 1, p.j - 1}) {
			sides++
		}
		if mapContains(&r.contained, point{p.i - 1, p.j}) && mapContains(&r.contained, point{p.i, p.j + 1}) && !mapContains(&r.contained, point{p.i - 1, p.j + 1}) {
			sides++
		}
		if mapContains(&r.contained, point{p.i + 1, p.j}) && mapContains(&r.contained, point{p.i, p.j + 1}) && !mapContains(&r.contained, point{p.i + 1, p.j + 1}) {
			sides++
		}
	}
	return sides
}

func mapContains(contains *map[point]bool, p point) bool {
	_, ok := (*contains)[p]
	return ok
}

func findRegion(p point, c string, field *[][]string, visited *map[point]bool) int {
	if !isInBounds(p, field) {
		return 1
	}
	if (*field)[p.i][p.j] != c {
		return 1
	}
	if _, ok := (*visited)[p]; ok {
		return 0
	}
	(*visited)[p] = true
	return findRegion(point{p.i - 1, p.j}, c, field, visited) + findRegion(point{p.i + 1, p.j}, c, field, visited) + findRegion(point{p.i, p.j - 1}, c, field, visited) + findRegion(point{p.i, p.j + 1}, c, field, visited)
}

func isInBounds(p point, field *[][]string) bool {
	return p.i >= 0 && p.i < len(*field) && p.j >= 0 && p.j < len((*field)[p.i])
}

func scorePerimeter(regions []region) int {
	total := 0
	for _, r := range regions {
		total += len(r.contained) * r.perimeter
	}
	return total
}

func parse() [][]string {
	absPath, err := filepath.Abs("d12/input.txt")
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
