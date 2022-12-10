package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	cycles      int
	instruction string
	value       int
}

func main() {
	program := strings.Split(puzzleInput(), "\n")
	executor := make([]Instruction, 0)
	signal := make(map[int]int)
	crt := make(map[int]string)

	for _, command := range program {
		if command == "noop" {
			executor = append(executor, Instruction{cycles: 1, instruction: command})
		}

		if strings.HasPrefix(command, "addx ") {
			split := strings.Split(command, " ")
			value, _ := strconv.Atoi(split[1])
			executor = append(executor, Instruction{cycles: 2, instruction: split[0], value: value})
		}
	}

	x := 1
	i := 0
	for true {
		i = i + 1
		if len(executor) <= 0 {
			break
		}

		signal[i] = i * x
		// fmt.Println("value X during cycle", fmt.Sprintf("%3d", i), "=", fmt.Sprintf("%4d", x), "total=", fmt.Sprintf("%4d", i*x))

		crtpos := i % 40
		if crtpos >= x && crtpos < x+3 {
			crt[i-1] = "#"
		} else {
			crt[i-1] = "."
		}

		if len(executor) > 0 {
			// fmt.Println(executor[0])
			executor[0].cycles = executor[0].cycles - 1
			if executor[0].cycles <= 0 {
				x = x + executor[0].value
				executor = executor[1:]
			}
		}

		// fmt.Println("value X after cycle", fmt.Sprintf("%3d", i), "=", fmt.Sprintf("%4d", x), "total=", fmt.Sprintf("%4d", i*x))
	}

	// fmt.Println("signal strength cycle 20", signal[20])
	// fmt.Println("signal strength cycle 60", signal[60])
	// fmt.Println("signal strength cycle 100", signal[100])
	// fmt.Println("signal strength cycle 140", signal[140])
	// fmt.Println("signal strength cycle 180", signal[180])
	// fmt.Println("signal strength cycle 220", signal[220])
	fmt.Println("total=", signal[20]+signal[60]+signal[100]+signal[140]+signal[180]+signal[220])
	fmt.Println()

	for i := 0; i < len(crt); i++ {
		fmt.Print(crt[i])
		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
}

func input() string {
	return `noop
addx 3
addx -5`
}

func largerInput() string {
	return `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`
}

func puzzleInput() string {
	return `addx 2
addx 3
addx 1
noop
addx 4
noop
noop
noop
addx 5
noop
addx 1
addx 4
addx -2
addx 3
addx 5
addx -1
addx 5
addx 3
addx -2
addx 4
noop
noop
noop
addx -27
addx -5
addx 2
addx -7
addx 3
addx 7
addx 5
addx 2
addx 5
noop
noop
addx -2
noop
addx 3
addx 2
addx 5
addx 2
addx 3
noop
addx 2
addx -29
addx 30
addx -26
addx -10
noop
addx 5
noop
addx 18
addx -13
noop
noop
addx 5
noop
noop
addx 5
noop
noop
noop
addx 1
addx 2
addx 7
noop
noop
addx 3
noop
addx 2
addx 3
noop
addx -37
noop
addx 16
addx -12
addx 29
addx -16
addx -10
addx 5
addx 2
addx -11
addx 11
addx 3
addx 5
addx 2
addx 2
addx -1
addx 2
addx 5
addx 2
noop
noop
noop
addx -37
noop
addx 17
addx -10
addx -2
noop
addx 7
addx 3
noop
addx 2
addx -10
addx 22
addx -9
addx 5
addx 2
addx -5
addx 6
addx 2
addx 5
addx 2
addx -28
addx -7
noop
noop
addx 1
addx 4
addx 17
addx -12
noop
noop
noop
noop
addx 5
addx 6
noop
addx -1
addx -17
addx 18
noop
addx 5
noop
noop
noop
addx 5
addx 4
addx -2
noop
noop
noop
noop
noop`
}
