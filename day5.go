package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	items []string
}

func main() {
	split := strings.Split(puzzleInput(), "\n\n")
	startStacks := strings.Split(split[0], "\n")
	actions := strings.Split(split[1], "\n")

	stacks := make([]Stack, 0)
	totalStacks := strings.Split(startStacks[len(startStacks)-1], "   ")

	for i := 1; i <= len(totalStacks); i++ {
		// fmt.Println("stack (i) =", i)
		items := make([]string, 0)

		for j := len(startStacks) - 2; j >= 0; j-- {
			// fmt.Println("row (j) =", j)
			position := ((i - 1) * 4) + 1
			el := string(startStacks[j][position])

			if el != " " {
				items = append(items, el)
			}
		}

		stacks = append(stacks, Stack{items: items})
	}

	fmt.Println("start =", stacks)

	for _, action := range actions {
		// fmt.Println("action=", action)

		re, _ := regexp.Compile(`\d+`)
		a := re.FindAll([]byte(action), -1)
		amount, _ := strconv.Atoi(string(a[0]))
		from, _ := strconv.Atoi(string(a[1]))
		to, _ := strconv.Atoi(string(a[2]))

		tmpstack := make([]string, 0)

		for i := 0; i < amount; i++ {
			fromStack := stacks[from-1]
			item := fromStack.items[len(fromStack.items)-1]
			stacks[from-1].items = fromStack.items[0 : len(fromStack.items)-1]

			tmpstack = append(tmpstack, item)
		}

		for i := 0; i < amount; i++ {
			item := tmpstack[len(tmpstack)-1]
			tmpstack = tmpstack[0 : len(tmpstack)-1]

			stacks[to-1].items = append(stacks[to-1].items, item)
		}
	}

	fmt.Println("final =", stacks)

	combinedTop := ""
	for _, stack := range stacks {
		combinedTop = combinedTop + stack.items[len(stack.items)-1]
	}
	fmt.Println("combined =", combinedTop)
}

func input() string {
	return `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
}

func puzzleInput() string {
	return `[G]                 [D] [R]        
[W]         [V]     [C] [T] [M]    
[L]         [P] [Z] [Q] [F] [V]    
[J]         [S] [D] [J] [M] [T] [V]
[B]     [M] [H] [L] [Z] [J] [B] [S]
[R] [C] [T] [C] [T] [R] [D] [R] [D]
[T] [W] [Z] [T] [P] [B] [B] [H] [P]
[D] [S] [R] [D] [G] [F] [S] [L] [Q]
 1   2   3   4   5   6   7   8   9 

move 1 from 3 to 5
move 5 from 5 to 4
move 6 from 7 to 3
move 6 from 1 to 3
move 1 from 1 to 9
move 1 from 1 to 4
move 3 from 6 to 9
move 2 from 7 to 5
move 1 from 5 to 7
move 1 from 7 to 2
move 2 from 2 to 5
move 2 from 6 to 3
move 6 from 8 to 9
move 7 from 3 to 9
move 1 from 8 to 7
move 8 from 9 to 7
move 5 from 4 to 8
move 1 from 6 to 2
move 2 from 8 to 4
move 9 from 9 to 1
move 2 from 8 to 5
move 1 from 8 to 5
move 5 from 9 to 2
move 1 from 6 to 8
move 5 from 1 to 7
move 1 from 8 to 2
move 2 from 1 to 7
move 1 from 2 to 6
move 4 from 5 to 4
move 2 from 1 to 4
move 13 from 7 to 8
move 3 from 8 to 6
move 2 from 6 to 8
move 10 from 3 to 5
move 2 from 7 to 6
move 3 from 5 to 6
move 10 from 8 to 1
move 1 from 8 to 6
move 6 from 2 to 4
move 1 from 5 to 8
move 5 from 6 to 3
move 2 from 8 to 6
move 1 from 7 to 9
move 2 from 2 to 7
move 3 from 5 to 1
move 2 from 7 to 2
move 6 from 6 to 3
move 7 from 5 to 6
move 5 from 3 to 2
move 10 from 1 to 8
move 2 from 1 to 3
move 8 from 3 to 7
move 9 from 4 to 8
move 1 from 9 to 2
move 2 from 7 to 8
move 4 from 6 to 9
move 1 from 4 to 9
move 5 from 7 to 4
move 3 from 6 to 5
move 1 from 1 to 5
move 14 from 4 to 8
move 3 from 9 to 7
move 4 from 5 to 9
move 2 from 4 to 1
move 27 from 8 to 6
move 2 from 7 to 2
move 2 from 7 to 4
move 4 from 2 to 9
move 7 from 8 to 4
move 10 from 4 to 1
move 18 from 6 to 5
move 6 from 9 to 2
move 1 from 9 to 5
move 11 from 2 to 6
move 2 from 5 to 4
move 1 from 2 to 8
move 2 from 4 to 9
move 2 from 8 to 3
move 1 from 6 to 8
move 4 from 9 to 7
move 4 from 7 to 8
move 7 from 5 to 1
move 4 from 6 to 3
move 2 from 3 to 7
move 6 from 5 to 3
move 2 from 8 to 2
move 14 from 6 to 2
move 3 from 8 to 1
move 15 from 2 to 3
move 1 from 6 to 1
move 14 from 3 to 2
move 2 from 2 to 5
move 1 from 9 to 3
move 13 from 1 to 3
move 4 from 2 to 6
move 10 from 1 to 3
move 2 from 6 to 9
move 6 from 2 to 9
move 6 from 5 to 2
move 2 from 6 to 8
move 7 from 9 to 5
move 1 from 5 to 8
move 2 from 7 to 6
move 34 from 3 to 6
move 19 from 6 to 2
move 12 from 6 to 9
move 3 from 6 to 3
move 2 from 3 to 2
move 1 from 6 to 5
move 17 from 2 to 8
move 2 from 3 to 2
move 8 from 9 to 4
move 7 from 5 to 2
move 5 from 4 to 1
move 4 from 1 to 6
move 1 from 1 to 6
move 6 from 6 to 8
move 2 from 8 to 4
move 17 from 8 to 6
move 2 from 4 to 5
move 17 from 6 to 9
move 22 from 9 to 7
move 1 from 5 to 2
move 20 from 2 to 7
move 29 from 7 to 9
move 1 from 4 to 7
move 3 from 8 to 3
move 1 from 8 to 5
move 3 from 8 to 2
move 2 from 2 to 4
move 27 from 9 to 7
move 2 from 3 to 2
move 1 from 5 to 2
move 18 from 7 to 5
move 1 from 3 to 2
move 1 from 5 to 6
move 18 from 5 to 3
move 1 from 6 to 3
move 2 from 9 to 5
move 10 from 3 to 5
move 4 from 3 to 6
move 1 from 7 to 1
move 1 from 5 to 1
move 6 from 7 to 6
move 1 from 6 to 2
move 4 from 4 to 8
move 5 from 5 to 4
move 1 from 3 to 8
move 2 from 1 to 8
move 2 from 2 to 5
move 3 from 3 to 8
move 6 from 8 to 2
move 1 from 3 to 9
move 1 from 6 to 3
move 6 from 2 to 8
move 7 from 8 to 4
move 8 from 5 to 2
move 5 from 4 to 6
move 2 from 8 to 3
move 2 from 3 to 9
move 1 from 3 to 9
move 2 from 7 to 1
move 2 from 1 to 2
move 12 from 2 to 4
move 1 from 9 to 7
move 1 from 6 to 2
move 9 from 7 to 9
move 1 from 8 to 2
move 9 from 9 to 8
move 6 from 7 to 8
move 4 from 4 to 1
move 6 from 2 to 5
move 1 from 4 to 9
move 3 from 1 to 9
move 6 from 4 to 5
move 5 from 8 to 9
move 8 from 4 to 6
move 3 from 9 to 8
move 1 from 9 to 3
move 3 from 8 to 3
move 5 from 9 to 2
move 3 from 2 to 6
move 3 from 6 to 9
move 3 from 6 to 2
move 4 from 2 to 6
move 6 from 9 to 7
move 1 from 1 to 8
move 8 from 8 to 5
move 20 from 5 to 3
move 2 from 2 to 8
move 6 from 7 to 1
move 10 from 6 to 3
move 4 from 6 to 7
move 4 from 1 to 9
move 2 from 1 to 2
move 3 from 6 to 9
move 5 from 8 to 3
move 3 from 7 to 9
move 17 from 3 to 2
move 1 from 6 to 2
move 2 from 6 to 9
move 1 from 6 to 4
move 12 from 9 to 2
move 1 from 4 to 7
move 8 from 3 to 8
move 8 from 8 to 9
move 7 from 9 to 2
move 1 from 9 to 7
move 18 from 2 to 9
move 1 from 7 to 2
move 2 from 7 to 1
move 1 from 1 to 2
move 4 from 2 to 7
move 15 from 9 to 3
move 1 from 9 to 1
move 2 from 1 to 8
move 6 from 2 to 4
move 8 from 2 to 1
move 2 from 8 to 5
move 2 from 9 to 3
move 4 from 4 to 1
move 2 from 5 to 8
move 2 from 8 to 9
move 14 from 3 to 1
move 2 from 9 to 7
move 2 from 4 to 3
move 1 from 2 to 9
move 5 from 7 to 9
move 21 from 1 to 9
move 2 from 1 to 6
move 3 from 2 to 4
move 1 from 7 to 3
move 19 from 9 to 5
move 1 from 2 to 7
move 1 from 7 to 2
move 3 from 4 to 2
move 19 from 5 to 7
move 2 from 2 to 5
move 1 from 5 to 3
move 1 from 3 to 4
move 8 from 9 to 4
move 1 from 6 to 3
move 1 from 2 to 6
move 1 from 2 to 1
move 8 from 7 to 3
move 5 from 4 to 7
move 2 from 6 to 4
move 1 from 5 to 9
move 1 from 1 to 6
move 1 from 1 to 2
move 2 from 4 to 7
move 1 from 4 to 2
move 2 from 4 to 9
move 1 from 6 to 8
move 1 from 1 to 5
move 1 from 8 to 6
move 1 from 1 to 4
move 25 from 3 to 1
move 1 from 4 to 2
move 2 from 3 to 6
move 3 from 1 to 9
move 6 from 9 to 8
move 1 from 6 to 3
move 1 from 2 to 9
move 15 from 7 to 6
move 2 from 2 to 6
move 1 from 3 to 8
move 1 from 1 to 4
move 6 from 8 to 4
move 1 from 3 to 8
move 1 from 8 to 5
move 2 from 5 to 2
move 8 from 6 to 7
move 1 from 8 to 7
move 1 from 9 to 4
move 9 from 4 to 5
move 19 from 1 to 3
move 9 from 3 to 5
move 6 from 7 to 2
move 2 from 1 to 7
move 7 from 2 to 4
move 7 from 5 to 6
move 5 from 4 to 3
move 3 from 5 to 8
move 1 from 2 to 4
move 2 from 4 to 8
move 14 from 6 to 1
move 6 from 5 to 6
move 1 from 5 to 2
move 7 from 1 to 6
move 1 from 2 to 4
move 4 from 6 to 4
move 1 from 5 to 4
move 2 from 1 to 9
move 2 from 9 to 4
move 2 from 1 to 8
move 9 from 3 to 6
move 3 from 7 to 4
move 4 from 8 to 6
move 3 from 7 to 6
move 1 from 7 to 2
move 1 from 7 to 5
move 3 from 8 to 4
move 26 from 6 to 1
move 8 from 1 to 2
move 1 from 6 to 4
move 5 from 2 to 7
move 2 from 2 to 4
move 10 from 4 to 7
move 1 from 6 to 1
move 22 from 1 to 2
move 1 from 6 to 1
move 6 from 4 to 7
move 1 from 5 to 1
move 1 from 1 to 2
move 21 from 7 to 2
move 38 from 2 to 3
move 8 from 2 to 6
move 2 from 4 to 8
move 2 from 8 to 2
move 1 from 1 to 3
move 1 from 2 to 8
move 1 from 2 to 5
move 6 from 6 to 4
move 2 from 4 to 2
move 2 from 2 to 6
move 1 from 8 to 2
move 28 from 3 to 1
move 11 from 1 to 2
move 8 from 1 to 7
move 4 from 6 to 4
move 8 from 3 to 1
move 8 from 2 to 5
move 6 from 5 to 4
move 2 from 5 to 4
move 8 from 3 to 4
move 22 from 4 to 1
move 2 from 3 to 5
move 33 from 1 to 5
move 26 from 5 to 6
move 4 from 5 to 7
move 2 from 2 to 7
move 2 from 7 to 2
move 2 from 7 to 8
move 2 from 8 to 3
move 6 from 1 to 3
move 5 from 5 to 1
move 1 from 5 to 7
move 7 from 7 to 5
move 4 from 5 to 6
move 5 from 1 to 8
move 4 from 2 to 4
move 2 from 7 to 4
move 2 from 7 to 3
move 5 from 4 to 6
move 1 from 8 to 2
move 1 from 2 to 4
move 10 from 3 to 6
move 44 from 6 to 9
move 2 from 5 to 7
move 1 from 5 to 8
move 41 from 9 to 1
move 1 from 6 to 4
move 2 from 8 to 1
move 1 from 7 to 3
move 1 from 3 to 8
move 2 from 9 to 8
move 29 from 1 to 9
move 2 from 1 to 5
move 2 from 8 to 3
move 1 from 3 to 5
move 2 from 5 to 9
move 1 from 5 to 7
move 25 from 9 to 2
move 10 from 2 to 1
move 1 from 7 to 8
move 2 from 4 to 1
move 2 from 8 to 9
move 1 from 8 to 6
move 4 from 2 to 4
move 4 from 2 to 5
move 1 from 6 to 5
move 1 from 2 to 7
move 2 from 4 to 1
move 18 from 1 to 3
move 8 from 9 to 4
move 15 from 3 to 9
move 3 from 4 to 8
move 4 from 5 to 8
move 4 from 2 to 4
move 10 from 9 to 4
move 4 from 8 to 5
move 2 from 7 to 2
move 11 from 4 to 9
move 12 from 4 to 9
move 2 from 5 to 7
move 4 from 2 to 4
move 5 from 8 to 1
move 1 from 5 to 6
move 1 from 4 to 6
move 1 from 3 to 9
move 1 from 5 to 7
move 4 from 1 to 6
move 6 from 1 to 5
move 6 from 5 to 9
move 3 from 7 to 6
move 9 from 6 to 5
move 8 from 5 to 2
move 7 from 2 to 3
move 1 from 3 to 1
move 7 from 3 to 5
move 2 from 4 to 1
move 1 from 2 to 6
move 2 from 1 to 3
move 8 from 5 to 9
move 3 from 1 to 3
move 1 from 6 to 1
move 2 from 4 to 1
move 1 from 5 to 2
move 2 from 1 to 6
move 2 from 6 to 3
move 2 from 3 to 2
move 2 from 2 to 4
move 1 from 2 to 6
move 3 from 3 to 9
move 2 from 4 to 8
move 3 from 3 to 1
move 4 from 1 to 7
move 2 from 8 to 4
move 7 from 9 to 6
move 1 from 1 to 4
move 11 from 9 to 7
move 3 from 9 to 3
move 14 from 9 to 5
move 6 from 6 to 5
move 4 from 5 to 9
move 10 from 7 to 6
move 1 from 3 to 7
move 2 from 4 to 1
move 4 from 7 to 9
move 9 from 6 to 1
move 3 from 6 to 5
move 15 from 9 to 1
move 1 from 4 to 7
move 4 from 9 to 7
move 12 from 5 to 1
move 3 from 7 to 3
move 4 from 7 to 2
move 1 from 9 to 3
move 22 from 1 to 2
move 21 from 2 to 6
move 3 from 1 to 9
move 1 from 3 to 7
move 1 from 7 to 3
move 1 from 3 to 2
move 8 from 1 to 4
move 1 from 9 to 2
move 7 from 4 to 8
move 3 from 3 to 9
move 3 from 3 to 5
move 4 from 2 to 3
move 1 from 1 to 3
move 4 from 8 to 5
move 2 from 8 to 3
move 5 from 3 to 2
move 6 from 5 to 3
move 2 from 5 to 8
move 2 from 1 to 7
move 2 from 7 to 4
move 15 from 6 to 9
move 8 from 3 to 1
move 3 from 5 to 9
move 2 from 4 to 9
move 8 from 1 to 3
move 8 from 9 to 8
move 1 from 1 to 4
move 3 from 5 to 9
move 4 from 8 to 1
move 1 from 3 to 9
move 2 from 4 to 3
move 2 from 8 to 6
move 3 from 8 to 7
move 8 from 2 to 5
move 3 from 5 to 2
move 4 from 3 to 4
move 3 from 6 to 1
move 2 from 5 to 9
move 4 from 4 to 1
move 2 from 5 to 6
move 1 from 5 to 4
move 2 from 2 to 1
move 4 from 3 to 9
move 1 from 7 to 3
move 2 from 7 to 4
move 2 from 4 to 7
move 1 from 6 to 7
move 1 from 2 to 8
move 2 from 3 to 9
move 14 from 1 to 8
move 1 from 6 to 2
move 2 from 7 to 1
move 3 from 8 to 3
move 6 from 8 to 5`
}
