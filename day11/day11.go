package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Monkey struct {
	Items       []big.Int `json:"items"`
	Operation   string    `json:"operation"`
	Test        int       `json:"test"`
	IfTrue      int       `json:"true"`
	IfFalse     int       `json:"false"`
	Inspections int       `json:"inspections"`
}

var monkeys = make(map[int]Monkey)
var sem = make(chan int, 1)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)

	data, err := ioutil.ReadFile("./input.txt")
	// data, err := ioutil.ReadFile("./puzzleInput.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n\n")

	for i, m := range input {
		mSplit := strings.Split(m, "\n")

		startingItems := strings.Split(strings.Replace(mSplit[1], "  Starting items: ", "", 1), ", ")
		items := make([]big.Int, 0)
		for _, item := range startingItems {
			worry, _ := strconv.Atoi(item)
			items = append(items, *big.NewInt(int64(worry)))
		}
		operation := strings.Replace(mSplit[2], "  Operation: new = ", "", 1)
		test, _ := strconv.Atoi(strings.Replace(mSplit[3], "  Test: divisible by ", "", 1))
		ifTrue, _ := strconv.Atoi(strings.Replace(mSplit[4], "    If true: throw to monkey ", "", 1))
		ifFalse, _ := strconv.Atoi(strings.Replace(mSplit[5], "    If false: throw to monkey ", "", 1))

		monkeys[i] = Monkey{
			Items:       items,
			Operation:   operation,
			Test:        test,
			IfTrue:      ifTrue,
			IfFalse:     ifFalse,
			Inspections: 0,
		}
	}

	// Reddit hint
	// The problem is that the size of the numbers explode. So can we still get a solution by using smaller numbers? Yes we can.
	// The key is this observation: we never need the exact values of the worry levels to get the answer. We just need how often they are thrown.
	// Suppose an item is divisible by 13,17 and 19. If you divide this value by the product (13 * 17 * 19), it is still divisible by 13,17 and 19.
	// This works in general for any set of divisors.
	lcm := monkeys[0].Test
	for i := 1; i < len(monkeys); i++ {
		lcm = lcm * monkeys[i].Test
	}

	for round := 0; round < 10000; round++ {
		// fmt.Println("round", round)

		for i := 0; i < len(monkeys); i++ {
			var wg sync.WaitGroup
			wg.Add(len(monkeys[i].Items))
			for _, item := range monkeys[i].Items {
				sem <- 1
				go calculateItem(item, monkeys[i], lcm, &wg)
				<-sem
			}
			wg.Wait()

			if entry, ok := monkeys[i]; ok {
				entry.Inspections = entry.Inspections + len(entry.Items)
				entry.Items = make([]big.Int, 0)
				monkeys[i] = entry
			}
		}
	}

	for i := 0; i < len(monkeys); i++ {
		fmt.Print("monkey ", i, " ")
		fmt.Println(printMonkey(monkeys[i]))
	}

	fmt.Println("sorting...")

	keys := make([]int, 0, len(monkeys))
	for k := range monkeys {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return monkeys[keys[i]].Inspections > monkeys[keys[j]].Inspections
	})
	for _, k := range keys {
		fmt.Println(k, printMonkey(monkeys[k]))
	}

	fmt.Println("2 most active monkeys", monkeys[keys[0]].Inspections*monkeys[keys[1]].Inspections)
}

func calculateItem(item big.Int, monkey Monkey, lcm int, wg *sync.WaitGroup) {
	defer wg.Done()
	worryLevel := item
	// fmt.Println("   Monkey inspects an item with a worry level of ", item)
	worryLevel = executeOperation(monkey.Operation, worryLevel)
	// fmt.Println("    worry level after operation", worryLevel)
	// worryLevel = *worryLevel.Div(&worryLevel, big.NewInt(int64(3)))
	worryLevel = *worryLevel.Mod(&worryLevel, big.NewInt(int64(lcm)))
	// fmt.Println("    Monkey gets bored with item. Worry level is divided by 3 to", worryLevel)

	test := big.NewInt(int64(monkey.Test))
	mod := *big.NewInt(0).Mod(&worryLevel, test)
	var toMonkey int
	if mod.Int64() == 0 {
		// fmt.Println("    Current worry level is divisible by", monkeys[i].Test)
		toMonkey = monkey.IfTrue
	} else {
		// fmt.Println("    Current worry level is not divisible by", monkeys[i].Test)
		toMonkey = monkey.IfFalse
	}

	// fmt.Println("    Item with worry level", worryLevel, "is thrown to monkey ", toMonkey)
	sem <- 1
	if entry, ok := monkeys[toMonkey]; ok {
		entry.Items = append(entry.Items, worryLevel)
		monkeys[toMonkey] = entry
	}
	<-sem
}

func executeOperation(operation string, old big.Int) big.Int {
	split := strings.Split(operation, " ")

	var left big.Int
	var right big.Int
	if split[0] == "old" {
		left = old
	} else {
		o1, _ := strconv.Atoi(split[0])
		left = *big.NewInt(int64(o1))
	}
	if split[2] == "old" {
		right = old
	} else {
		o2, _ := strconv.Atoi(split[2])
		right = *big.NewInt(int64(o2))
	}

	switch split[1] {
	case "*":
		return *left.Mul(&left, &right)
	case "+":
		return *left.Add(&left, &right)
	}
	panic("oh no")
}

func printMonkey(monkey Monkey) string {
	buff, _ := json.Marshal(&monkey)
	return string(buff)
}
