package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type point struct {
	i, j int
}

type region struct {
	perimeter, area int
	contained       []point
}

func main() {
	fmt.Println(doP1())
	fmt.Println(doP2())

}

func doP1() int {
	field := parse()
	regions := make([]region, 0)
	for i, v := range field {
		for j, c := range v {
			p := point{i, j}
			check := true
			for _, r := range regions {
				if slices.Contains(r.contained, p) {
					check = false
					break
				}
			}
			if !check {
				continue
			}
			prev := make(map[point]bool)
			region := findRegion(p, c, &field, &prev)
			regions = append(regions, region)

		}
	}
	return scorePerimeter(regions)
}

func doP2() int {

	field := parse()
	regions := make([]region, 0)
	for i, v := range field {
		for j, c := range v {
			p := point{i, j}
			check := true
			for _, r := range regions {
				if slices.Contains(r.contained, p) {
					check = false
					break
				}
			}
			if !check {
				continue
			}
			prev := make(map[point]bool)
			region := findRegion(p, c, &field, &prev)
			regions = append(regions, region)

		}
	}
	return scoreSides(regions)

}

func scoreSides(regions []region) int {
	total := 0
	for _, r := range regions {
		total += r.area * countSides(r)
	}
	return total
}

// actually counting corners
func countSides(r region) int {
	sides := 0
	for _, p := range r.contained {

		//outer corner
		if !slices.Contains(r.contained, point{p.i - 1, p.j}) && !slices.Contains(r.contained, point{p.i, p.j - 1}) {
			sides++
		}
		if !slices.Contains(r.contained, point{p.i + 1, p.j}) && !slices.Contains(r.contained, point{p.i, p.j - 1}) {
			sides++
		}
		if !slices.Contains(r.contained, point{p.i - 1, p.j}) && !slices.Contains(r.contained, point{p.i, p.j + 1}) {
			sides++
		}
		if !slices.Contains(r.contained, point{p.i + 1, p.j}) && !slices.Contains(r.contained, point{p.i, p.j + 1}) {
			sides++
		}

		//# Inner corners
		if slices.Contains(r.contained, point{p.i - 1, p.j}) && slices.Contains(r.contained, point{p.i, p.j - 1}) && !slices.Contains(r.contained, point{p.i - 1, p.j - 1}) {
			sides++
		}
		if slices.Contains(r.contained, point{p.i + 1, p.j}) && slices.Contains(r.contained, point{p.i, p.j - 1}) && !slices.Contains(r.contained, point{p.i + 1, p.j - 1}) {
			sides++
		}
		if slices.Contains(r.contained, point{p.i - 1, p.j}) && slices.Contains(r.contained, point{p.i, p.j + 1}) && !slices.Contains(r.contained, point{p.i - 1, p.j + 1}) {
			sides++
		}
		if slices.Contains(r.contained, point{p.i + 1, p.j}) && slices.Contains(r.contained, point{p.i, p.j + 1}) && !slices.Contains(r.contained, point{p.i + 1, p.j + 1}) {
			sides++
		}

	}
	fmt.Println(r.area, sides, r.contained)
	return sides
}

func findRegion(p point, c string, field *[][]string, visited *map[point]bool) region {
	if !isInBounds(p, field) {
		(*visited)[p] = true
		return region{1, 0, []point{}}

	}
	if (*field)[p.i][p.j] != c {
		(*visited)[p] = true
		return region{1, 0, []point{}}
	}

	if _, ok := (*visited)[p]; ok {
		return region{0, 0, []point{}}
	}

	(*visited)[p] = true
	up := findRegion(point{p.i - 1, p.j}, c, field, visited)
	down := findRegion(point{p.i + 1, p.j}, c, field, visited)
	left := findRegion(point{p.i, p.j - 1}, c, field, visited)
	right := findRegion(point{p.i, p.j + 1}, c, field, visited)

	contained := append(make([]point, 0), p)
	contained = slices.Concat(contained, up.contained, down.contained, left.contained, right.contained)
	return region{up.perimeter + down.perimeter + right.perimeter + left.perimeter, 1 + up.area + down.area + left.area + right.area, contained}
}

func isInBounds(p point, field *[][]string) bool {
	return p.i >= 0 && p.i < len(*field) && p.j >= 0 && p.j < len((*field)[p.i])
}

func scorePerimeter(regions []region) int {
	total := 0
	for _, r := range regions {
		total += r.area * r.perimeter
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
