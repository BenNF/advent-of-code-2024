package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	add    = iota
	mul    = iota
	concat = iota
)

type row struct {
	target     int64
	components []int64
}

type opTree struct {
	value     int64
	total     int64
	mulBranch *opTree
	addBranch *opTree
	conBranch *opTree
	parent    *opTree
	op        int
}

func main() {
	fmt.Println(doP1())
	fmt.Println(doP2())
}

func doP1() int64 {
	return runAsWg(func(r row) int64 {
		return solveTree(r, false)
	})
}

func doP2() int64 {
	return runAsWg(func(r row) int64 {
		return solveTree(r, true)
	})
}

func solveTree(r row, doConcat bool) int64 {
	tree := buildTree(r, doConcat)
	if doesTreeSolve(&tree, r.target, doConcat) {
		return r.target
	}
	return 0
}

func runAsWg(f func(r row) int64) int64 {
	rows := parse()
	start := time.Now()

	wg := sync.WaitGroup{}
	var total atomic.Int64

	for _, v := range rows {
		wg.Add(1)
		go func(r row) {
			defer wg.Done()
			total.Add(f(r))
		}(v)
	}
	wg.Wait()
	fmt.Println("took", time.Since(start))
	return total.Load()
}

func doesTreeSolve(tree *opTree, target int64, doConcat bool) bool {
	stack := []*opTree{tree}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.op == add && current.parent != nil {
			current.total = current.parent.total + current.value
		}
		if current.op == mul && current.parent != nil {
			current.total = current.parent.total * current.value
		}

		if doConcat && current.op == concat && current.parent != nil {
			strTotal := strconv.FormatInt(current.parent.total, 10)
			strValue := strconv.FormatInt(current.value, 10)
			var err error
			current.total, err = strconv.ParseInt(strTotal+strValue, 10, 64)
			if err != nil {
				panic(err)
			}
		}

		//leaf and value matches, we're done it solves
		if current.total == target && current.mulBranch == nil && current.addBranch == nil {
			return true
		}

		//we never shrink, don't search deeper
		if current.total > target {
			continue
		}

		if current.addBranch != nil && current.mulBranch != nil {
			stack = append(stack, current.addBranch, current.mulBranch)
		}
		if doConcat && current.conBranch != nil {
			stack = append(stack, current.conBranch)
		}
	}
	return false

}

func buildTree(r row, doCon bool) opTree {
	root := opTree{r.components[0], r.components[0], nil, nil, nil, nil, -1}
	leafs := []*opTree{&root}
	for i := 1; i < len(r.components); i++ {
		tmpLeafs := make([]*opTree, 0)
		for _, t := range leafs {
			t.mulBranch = &opTree{r.components[i], 0, nil, nil, nil, t, mul}
			t.addBranch = &opTree{r.components[i], 0, nil, nil, nil, t, add}
			if doCon {
				t.conBranch = &opTree{r.components[i], 0, nil, nil, nil, t, concat}
			}
			if doCon {
				tmpLeafs = append(tmpLeafs, t.conBranch)
			}
			tmpLeafs = append(tmpLeafs, t.mulBranch, t.addBranch)
		}
		leafs = make([]*opTree, len(tmpLeafs))
		copy(leafs, tmpLeafs)
	}
	return root
}

func parse() []row {
	absPath, err := filepath.Abs("d7/input.txt")
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	output := make([]row, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		split := strings.Split(scanner.Text(), ":")
		target, err := strconv.ParseInt(split[0], 10, 64)
		if err != nil {
			panic(err)
		}

		fields := strings.Fields(split[1])
		commponents := make([]int64, len(fields))
		for i, s := range strings.Fields(split[1]) {
			commponents[i], err = strconv.ParseInt(s, 10, 64)
			if err != nil {
				panic(err)
			}
		}

		output = append(output, row{target, commponents})
	}
	return output
}
