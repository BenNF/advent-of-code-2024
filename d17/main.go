package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("p1: ", doP1(), "in: ", time.Since(now))
	now = time.Now()
	fmt.Println("p2: ", doP2(), "in: ", time.Since(now))
}

func doP1() string {
	a, _, _, program := parse()
	output := eval(a, program)
	strOutput := make([]string, len(output))

	for i, c := range output {
		strOutput[i] = strconv.Itoa(c)
	}
	return strings.Join(strOutput, ",")
}

func doP2() int {
	_, _, _, program := parse()
	success := math.MaxInt
	dfs(program, 0, 0, &success)
	return success
}

func dfs(program []int, cur, pos int, success *int) {
	for i := 0; i < 8; i++ {
		nextNum := (cur << 3) + i
		result := eval(nextNum, program)

		if slices.Compare(result, program[len(program)-pos-1:]) != 0 {
			continue
		}
		if pos == len(program)-1 {
			*success = int(math.Min(float64(*success), float64(nextNum)))
			return
		}
		dfs(program, nextNum, pos+1, success)
	}
}
func eval(a int, program []int) []int {
	b := 0
	c := 0
	idx := 0
	output := make([]int, 0)
	var opcode, operand int
	for idx < len(program) {
		opcode = program[idx]
		operand = program[idx+1]
		didJump := false

		switch opcode {
		case 0: //adiv
			pow := int(math.Pow(2, float64(evalComboOperand(operand, a, b, c))))
			a = a / pow
		case 1: //xor b literal operand
			b = b ^ operand
		case 2: //mod 8, trim to 3 bit
			b = evalComboOperand(operand, a, b, c) % 8
		case 3: //jump instruction
			if a != 0 {
				idx = operand
				didJump = true
			}
		case 4: //xor b/c
			b = b ^ c
		case 5: //out
			output = append(output, evalComboOperand(operand, a, b, c)%8)
		case 6: //bdiv
			pow := int(math.Pow(2, float64(evalComboOperand(operand, a, b, c))))
			b = a / pow
		case 7: //cdiv
			pow := int(math.Pow(2, float64(evalComboOperand(operand, a, b, c))))
			c = a / pow
		default:
			panic("Unknown opcode")
		}
		if !didJump {
			idx += 2
		}
	}
	return output
}

func evalComboOperand(operand int, a, b, c int) int {
	switch operand {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		return operand
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	default:
		panic("Unknown operand")
	}
}

func parse() (int, int, int, []int) {
	absPath, err := filepath.Abs("d17/input.txt")
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(absPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(data)))

	rExpr := regexp.MustCompile("Register [A-Z]: (\\d+)")
	scanner.Scan()
	a, _ := strconv.Atoi(rExpr.FindStringSubmatch(scanner.Text())[1])
	scanner.Scan()
	b, _ := strconv.Atoi(rExpr.FindStringSubmatch(scanner.Text())[1])
	scanner.Scan()
	c, _ := strconv.Atoi(rExpr.FindStringSubmatch(scanner.Text())[1])
	scanner.Scan()
	scanner.Scan()
	progChars := strings.Split(scanner.Text(), ",")
	prog := make([]int, len(progChars))

	for i, c := range progChars {
		prog[i], _ = strconv.Atoi(c)
	}

	return a, b, c, prog
}
