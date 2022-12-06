package main

import (
	"fmt"
	"strings"
)

func main() {
	rucksacks := strings.Split(input(), "\n")
	// fmt.Println(rucksacks)

	priorities := make([]int, 0)

	for _, rucksack := range rucksacks {
		compartment1 := rucksack[0 : len(rucksack)/2]
		compartment2 := rucksack[len(rucksack)/2:]
		// fmt.Println(rucksack)
		// fmt.Println(compartment1)
		// fmt.Println(compartment2)

		common := findCommon(compartment1, compartment2)
		// fmt.Println(common)
		for key := range common {
			priority := getPriority(key)
			// fmt.Println(priority)
			priorities = append(priorities, priority)
		}
	}

	fmt.Println(priorities)
	sum := 0
	for _, priority := range priorities {
		sum = sum + priority
	}
	fmt.Println(sum)
	fmt.Println()

	badgePriorities := make([]int, 0)
	for i := 0; i < len(rucksacks); i = i + 3 {
		rucksack1 := rucksacks[i]
		rucksack2 := rucksacks[i+1]
		rucksack3 := rucksacks[i+2]

		// fmt.Println(rucksack1)
		// fmt.Println(rucksack2)
		// fmt.Println(rucksack3)
		// fmt.Println()
		r1r2 := findCommon(rucksack1, rucksack2)
		r1r2string := ""
		for key := range r1r2 {
			r1r2string = r1r2string + key
		}
		common := findCommon(r1r2string, rucksack3)
		// fmt.Println(r1r2)
		// fmt.Println(common)
		for key := range common {
			priority := getPriority(key)
			badgePriorities = append(badgePriorities, priority)
		}
	}

	fmt.Println(badgePriorities)
	badgeSum := 0
	for _, priority := range badgePriorities {
		badgeSum = badgeSum + priority
	}
	fmt.Println(badgeSum)
}

func getPriority(character string) int {
	if character[0] >= 'a' && character[0] <= 'z' {
		return (int)(character[0] - 96)
	}
	if character[0] >= 'A' && character[0] <= 'Z' {
		return (int)(character[0] - 64 + 26)
	}
	return 0
}

func findCommon(compartment1 string, compartment2 string) map[string]bool {
	common := map[string]bool{}
	for i := 0; i < len(compartment1); i++ {
		for j := 0; j < len(compartment2); j++ {
			if compartment1[i] == compartment2[j] {
				common[string(compartment1[i])] = true
			}
		}
	}
	return common
}

func input() string {
	return `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
}
