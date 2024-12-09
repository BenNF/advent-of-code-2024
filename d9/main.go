package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	fmt.Println(doP1())
	fmt.Println(doP2())
}
func doP1() int {
	tape := parse()

	return checkSum(compact(tape))
}

func doP2() int {
	tape := parse()
	fmt.Println(len(tape))
	tape = continuousCompact(tape)

	return checkSum(tape)

}

func continuousCompact(tape []int) []int {
	id, idx, size := seekId(tape, math.MaxInt, len(tape)-1)

	for id != -1 {
		freeIdx := seekSpaceLeft(tape, size, idx)
		if freeIdx == -1 {
			id, idx, size = seekId(tape, id, idx)
			continue //not enough space to move id block
		}

		for i := freeIdx; i < freeIdx+size; i++ {
			tape[i] = id
		}
		for i := idx; i < idx+size; i++ {
			tape[i] = -1
		}

		id, idx, size = seekId(tape, id, idx)
	}
	return tape
}

func seekSpaceLeft(tape []int, size int, bound int) int {
	start := -1
	count := 0
	for i, v := range tape[:bound] {
		if v == -1 {
			if start == -1 {
				start = i
			}
			count++
			if count == size {
				return start
			}
		} else {
			start = -1
			count = 0
		}
	}
	return -1
}

func seekId(tape []int, idBound int, idxStart int) (int, int, int) {
	end := -1
	id := -1
	start := -1
	for i := idxStart; i >= 0; i-- {
		if tape[i] != -1 && tape[i] < idBound && id == -1 {
			end = i
			id = tape[i]
		}

		if id != -1 && tape[i] != id {
			start = i + 1
			break
		}
	}

	//block ends at the start of array
	if id != -1 && end != -1 && start == -1 {
		start = 0
	}
	return id, start, end - start + 1
}
func checkSum(tape []int) int {
	total := 0
	for i, v := range tape {
		if v == -1 {
			continue
		}
		total += v * i
	}
	return total
}

func compact(tape []int) []int {
	freeidx := seekFree(tape, 0)
	for i := len(tape) - 1; i >= 0; i-- {
		if tape[i] == -1 {
			continue
		}
		if freeidx >= i {
			break
		}
		id := tape[i]
		tape[freeidx] = id
		tape[i] = -1
		freeidx = seekFree(tape, freeidx+1)
	}
	return tape
}

func seekFree(in []int, start int) int {
	for i := start; i < len(in); i++ {
		if in[i] == -1 {
			return i
		}
	}
	return -1
}

func parse() []int {
	absPath, err := filepath.Abs("d9/input.txt")
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	strData := string(data)
	idSaturated := make([]int, 0)
	id := 0
	isData := true
	for i, _ := range strData {
		v, err := strconv.Atoi(string(strData[i]))
		if err != nil {
			panic(err)
		}
		if isData {
			for i := 0; i < v; i++ {
				idSaturated = append(idSaturated, id)
			}
			id++
			isData = false
		} else {
			for i := 0; i < v; i++ {
				idSaturated = append(idSaturated, -1)
			}
			isData = true
		}

	}

	return idSaturated
}
