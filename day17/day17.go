package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	mainStart := time.Now()
	fmt.Println("--- Part 1 ---")

	jetPattern := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"
	// jetPattern := puzzleInput()
	rockFormations := strings.Split(formations(), "\n\n")

	chamber := make(map[int]map[int]string, 0) //yx-map
	currentJet := 0
	currentRock := 0

	iterations := 2022
	// iterations = 1000000000000
	for i := 0; i < iterations; i++ {
		// Get next rock
		rockToDraw := strings.Split(rockFormations[currentRock], "\n")
		currentRock++
		if currentRock >= len(rockFormations) {
			currentRock = 0
		}
		for i, j := 0, len(rockToDraw)-1; i < j; i, j = i+1, j-1 {
			rockToDraw[i], rockToDraw[j] = rockToDraw[j], rockToDraw[i]
		}
		rockWidth := len(rockToDraw[0])
		rockHeight := len(rockToDraw)
		rockX := 2
		rockY := towerHeight(&chamber) + 2 + rockHeight

		// play jets until rock settled
		settled := false
		for settled == false {
			// Jet
			jet := string(jetPattern[currentJet])
			currentJet++
			if currentJet >= len(jetPattern) {
				currentJet = 0
			}
			switch jet {
			case ">":
				collision := false
				for y := 0; y < rockHeight; y++ {
					for x := rockWidth - 1; x >= 0; x-- {
						if rockToDraw[rockHeight-1-y][x] == '#' {
							if chamber[rockY-y][rockX+x+1] == "#" {
								collision = true
							}
							break
						}
					}
				}

				if rockX+rockWidth < 7 && !collision {
					rockX++
				}
				break
			case "<":
				collision := false
				for y := 0; y < rockHeight; y++ {
					for x := 0; x < rockWidth; x++ {
						if rockToDraw[rockHeight-1-y][x] == '#' {
							if chamber[rockY-y][rockX+x-1] == "#" {
								collision = true
							}
							break
						}
					}
				}

				if rockX > 0 && !collision {
					rockX--
				}
				break
			}

			// Check if rock settled
			for x := 0; x < rockWidth; x++ {
				for y := 0; y < rockHeight; y++ {
					if rockToDraw[y][x] == '#' {
						if chamber[rockY-(rockHeight-y)][rockX+x] == "#" || rockY-(rockHeight-y) < 0 {
							settled = true
						}
						break
					}
				}
			}

			// Fall down
			if !settled {
				rockY--
			}
		}

		// Place rock in chamber
		for y := 0; y < rockHeight; y++ {
			for x := 0; x < rockWidth; x++ {
				if rockToDraw[rockHeight-1-y][x] == '#' {
					addRow(&chamber, rockY-y)
					chamber[rockY-y][rockX+x] = "#"
				}
			}
		}

		// fmt.Printf("Total tower height after %d iterations: %d\n", i+1, towerHeight(&chamber))
	}

	fmt.Println()
	for x := towerHeight(&chamber); x >= 0; x-- {
		fmt.Printf("%4d ", x)
		for y := 0; y < 7; y++ {
			fmt.Print(chamber[x][y])
		}
		fmt.Println()
	}
	fmt.Printf("\nTotal tower height after %d iterations: %d\n", iterations, towerHeight(&chamber))

	elapsed := time.Since(mainStart)
	fmt.Printf("Part 1 took %s\n", elapsed)
}

func addRow(chamber *map[int]map[int]string, row int) {
	if (*chamber)[row] == nil {
		(*chamber)[row] = make(map[int]string)
		for i := 0; i < 7; i++ {
			(*chamber)[row][i] = "."
		}
	}
}

func towerHeight(chamber *map[int]map[int]string) int {
	height := 0
	for n := range *chamber {
		if n+1 > height {
			height = n + 1
		}
	}
	return height
}

func formations() string {
	return `####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##`
}

func puzzleInput() string {
	return `>>><<>>><><<>>>><<<>><>>><<<<>>>><<><<>>><<<<><>>>><<<<>>><<>><>><<<>>>><><<<>>>><<>><<>>><<<>>><<>>>><>>>><<<>>>><<<<><>>>><<<<>><<<<>>><>><>>>><>>>><><<<>>>><<<<>>>><<>>>><<<<>>><<<>>><<<<><<>><<<><<<>>>><>><>>>><<<<>><>>><<>>><<><<>>>><<<<>>><<>>><>>><>>>><<<<>><<<>><<<<>>>><<>>><<><<>><<>><>><<>>>><><<<<>>>><<><>><><<>><<<<>><<<<>>><<<<>>>><>>><><<<>>>><<<<>><<>><<<<>>>><<<<>>><>>>><<<><<><<>><><<>>><>>>><<>>><<<<><<<><>>><>><<<>>><<<>>>><<<>>><<>>>><<<>><<<<>>>><<<>>><<<><<<<><<<>>><>><>>>><<<<><<>>><<<<>>><<><<<<><<<>><<>><>><<<<>>><<<>>><<<<><<>>><<<<><<<><<<<>>><<<<>>>><<>>><<><<<<>>>><<>>><<><<<><<>><<<<><<<>>><<<<>>>><<>><<<<>><>><>><>>><><<<>><<>><>>><>><<<><>>>><>>><<><>>><<>>>><>>>><>><<>>>><>><<<<><<<<>>><<<<>><<>>><>>><<<<><<<<>>>><<<<>><<<<>>>><<<<>><<<><<<>>><<><<>>>><>><<<<>>><<>>>><<<>><<<>>>><<>><<><<<<>>>><><<<<>>>><><<>><<<<>>><<>>>><<<<>>>><><>><<<>>><<<<><<<>><<>>>><<<>>><<<<>>>><>>><>>><<>>>><<<>>>><<<>>><<>>>><<>><<>>><>><<<>>><<<>>>><<<>><<<<><<<>>><<<>>><<>><<<>><<<<>>><<>>>><<<<>>>><>>><>><<<>><<<><<<<><>><<<<>>><<<<><>>>><<<>><<<><<<<><<<<>>>><<<<>><<<>><<<><<>><<<>>><<<<><<>>>><<>>><<>>>><>>>><<<>>><<<>><<<<>><<<><>><><<<<>>><>>>><<>><<<<>><>>>><<<>><<><>>>><<<<>><>>>><<>>><<<<>><>>><>>>><>>><<<><<<>>><<><<>>>><<<<><><<<<>><<<>>><<<><<>><<>><<<<><<>>><<>><<<<><<>><<<>><>>><>><<<<>><<<>><<<>>>><<<>>>><<<<>>>><>><><<<>>>><<<><<<>>><<>>><<>><<<>>><>>>><<<><>>><<><>>>><>><>><<<<>>>><<>><>>>><<>>><>>><<>>><>>>><<<>><<>>>><<<>>>><<><<<<>>><>>><>>><<<>>><<>>><<<<>>><<<>><>>><<<>>>><>>>><>><<<<>>><>>><>>><>>><<<<>>><<><<<>><<<>>><<<<>>><<<><<<>>><>><><>><<>>><<<<>>>><<<>>><<<><<<>>>><<<>>><<<<>>><<<<>>><<<<>>>><>>><<><<>><<>>><<<<>>><<>><<<<>>>><<<><>>>><<<>>>><<<>><<>><<><<>>>><<<<>><<>><<<<><<<<><<<<>>>><<>>>><<<>>><><<<<>>><>>><<>>>><<<>>>><<<>><>>>><<<<>>>><<<><<<>>>><<><<<<>>>><<>>>><<>><<><<<<><<>><<><<<<>>>><>><<<>>>><><<>>>><>>>><>>><<<<>><<><<<>><<>>>><><<><<>>><>><<<><<<>><<<<>><<>>><<<><<<<>>>><>>>><>><>>>><<<>><<><<<<><>>>><<>><<<>>><<<>>><<<>><<<<><><<<>>><<<<>><<<>>>><><<>>>><<>><<<>>><<<<><<<<><>><<>><><<><<>>>><<<>><><<><>><<<<><>><>><>>><<<>><<<>>>><<>><<>><<<<>><>>>><<<<>><>>>><<<<>>><><<<>>><<>><<<>>><<<<><<<>>><<>>>><<<>>><>>>><<<<>>><><>><><><<<>>>><>>>><<<>>><><<><<>>>><<<<>>><>><<<<>>>><>>>><>><>>><<<<>>><<>>><>>><<><<<><>><<><<>>>><<<<>><<>><<<<>>><<<><<><<>>><<<<>>>><<<<>>><<>>>><<<>>><<<<>>>><<<<>>>><<<<>><<<>>><<<<>>><<<>>><<<>><<<>>>><<<<>>>><>>><>>>><<<>>><<<>><<<><<><>>>><<><<>><<<<>>><><<<<><<>>>><<><<<>>><<>><<<<><>>><<>>><<<>>>><<><<<>>>><<><<>>>><><>>>><<>>><>><<>><<>><<><<<<>>>><<<><>><<<<><<<<>><<>><<<>>><<><<><<<<>>>><<<><<<<>>>><<<>><<>>>><<<>><<>><>><><<<>>>><<><<<<>>>><<<>>><<<>>><>>><<<>>><<<>>><<<><<>>><<>>><<<<><<<>><<<>>><<<<>>><<<<>><<<>>>><>><<<>>>><<>><<><<>>>><<<><<>>><<><<<><<<<><>>>><<>>>><<<<>><>><<<<>>><<<<>><>><<<<><><<>>><><<>><<<<>><>>><>>><<<<>>>><<<<><<>><><<<>><><<<<>>><>>>><<<>>>><>><>>>><<<>>>><<>>><<<<>>><<><>>><<<<><<<>>>><<>><<<<><<<>><<>>>><<><><<>><<<>>>><<<>><<><<<>>><<>>><>>><>>>><<>>><<<<>><<<><<>>><<<><<>>><>>><<<<>><<<>>><<>><<<<>><<><<>>><<>><<<><>>>><>><<<><>>><<>>>><<>><><<<<><<<><<<><<<>>><<<>>><<<<>><<<><><>><<<>>>><<>>><<<>><<>>><><>>><>><>>>><<>><<>>><<><<<>>><<<>>>><<<>><<<>>>><>><<><><<><<<<>><<<<>><<>>><<<><>>><<<><<<<>>><<<><<<>>>><<>><<<<>><<<><<>>>><<>><<<<>><<>><<>>><<<<>><<>><><<<><<<>><<<<><<>><>>>><<<>>><>>><<<>>>><<>>>><>>>><<<<>>>><>>>><>>><<<>>>><<>><<<><><<<><<<>>>><<<>>>><<<<>><<<>>>><<<>>>><<<>><><<<<>>>><>>><<<<>>>><<<<><>>><<<<><<<><<<<>>>><><<>>><<<<><<>>><>>><<>>>><>><<<<>>><<<>><>>><<><<<>>><<<<>>><<>>><<<>>><<>>>><<<<>>>><>>>><<<>><<<<>>><<><>><<>>><<<<><<>><<<<><<>>><<<<><<<><<<>>><<><>><<<<>>><<<<>><<<<>>><<><><>>>><>>><<<>>><><>>>><<<>>><>>>><<<>>>><<<>><<>>><<<>>><>>>><>>>><>>><<>>><<<><>><<<>>><<<>>><<<<>>>><<>>><<<<>><<<<>>><<<<><<<<>>><<>><<<<><<>>><<<>>>><<<>><<<><<>>>><<<>>>><><<>>><>>>><<<>><<<>>>><<<<>><>>>><<><<<>>>><<<<>><<<<>>><<>>><<<>>><><<<>>>><<<><<<><>><>>><><<<>>><><<<<>>>><<<<>>>><<<>>>><>>><<>>><>><<<<>>>><<>><<<<><<<<>><>>><<<<><<<<>><<<>>>><<<><<<<><<<<><>>><<<>>><<<<>><<<>>><<><<>>><<<>>><>><<><<>>><<<<>>>><><>>><<<><<<>><<<<>>><<<>><<>><<<>>><<<<><<<<>>>><<<<>><<>><<>>>><<><<<>>>><<<>>>><<<<>>><<<<><>>>><>>><<><<>>><<<>>>><>><<<><<>>>><<><<<>>><>><<>>><<<<>>>><<<>>>><>>><<><<<<><<><>>><<>><<>><<>><<<<>><<<<>>>><<<<>><<<><<>>>><>>>><>>><<<<>><<<><<>>><>>><>><<<>><><<>><<>>>><<><<<<><<<<>>><<>>>><<>><<<><<<<><<>>><<><<<>>>><<<<><>>>><<<<>>>><>>>><<>>>><>><>>><<><<<><<>>>><<<><<><<>>><<><<<<><<<>><<<<>><<><<<>><<<><<<>>><<>><<<>>><<<>>>><>>>><<<>>>><<>><><<>>>><<<<>><<<<>><<<>>>><><<>>>><<<>>>><<<>>>><<>><<<<>>><<<<>><<<>>>><>>><<<<>>><<<>>><><<<<>>><<<>><<>>>><<<>>><<<<>>><<>>>><<>>><<<<>>>><>>><<><<<<>>>><<<>>>><<<<>>><><<<<>><<<<>>><<>>><<<<>>>><<<><<<<>><<<<>>>><<>>><>><><<<>>><<<><>>><>>>><<<><><<><><>><<><>>><<<<>><<>>><<<>>>><>>>><<<<><<<>>><<<<>><<<><>>>><<<<>>><<<<>><<<<><<<<>><<<>>><<<>>><<>>>><>>><<>>>><<<><<><<>>>><<>><<><<<<>>>><><<>><<<>>><>><<<<><<<<>>><<>>><<>><<<>>><<<>><<>>><<<<>>><>>><>>>><>>><>>><>>>><<>><<<<>><<<<>>>><<<<><<<>>><<>>>><<<><<><<<>>><<>>>><<><><>>>><<<>>>><>><<><<<<><>>><>>>><<<>><<<<><<<<><>>><<<<><<>><<<<>>>><<<>>><<<><>>><<><<<<>><>>><<<<>>><<<>>><>>>><<<<>>>><>><<>>><<<><<><<<<>><<<>>>><><<<>><<<>>><<<><<>>><<<>><<<<>>>><<><<>><>><<>>>><<<>>><<>>><<<>>><>>>><>>>><<>>><<<>>>><>><>>>><<<>>>><><<><<<<>><><<>>><<<>>>><<<<><<<<>><<<<>>>><<<<>>><>>><>>><<>>>><><<><<<>>><<>>>><<<>>>><>>>><<>>>><<<><<><<<<>>>><<>>><<<<><<<><<>>>><<<<><<>>><>>>><<>>><<<<><<>><<<<>><<<>><<<><<<<><<<>>><>><<<>>>><<<<>>>><<<<>><<<<>><>><>>>><>>>><<>>><<<><>>><><<<<>>>><<>>><<><<<>>>><<>><<>><<<<>><<<<>>><<<><<>>><<<<>>>><<<>>>><<<>><<<<>><<<>>>><<<<>><>>>><<><<<>><<<><<<>>>><<<>>><<>>><<>>>><<<>>>><>>><<<<>>><>>>><<<>>>><<<><<>>><<<<>><<<>><>><>>><>>>><<><<<>>>><<<><<<><<<<><<<>>><<<<>>><>><<>>>><<<>><<<>>><<>>><>><<<<>>><<><<<<>><<>>>><<<<><<><>>>><<<>>><<<>><>><><<<<>>>><<<<><<<>>><<>><<>>><<<>>><>>><<<<>>>><<<>>><>><<><<<>>>><<<>>>><<<<>>>><<<><<<>>>><<<><<<>>><><<<>>><<>>><><<>>><>>>><<<>>><><<<<>>><<><<<>><>>><>><<<<>>><<<<><<<>>><>><<<<>>><<<>><<><<<<><<<><<>>>><<<<>><<>><<<><<>>>><<>>><<>>>><<<<>><<<<>>><<>>>><>>>><<<<>>><<<>>>><<<>>><>><<<<>>>><<<<>>>><<<><>>><<<<><<<<><<><<<>><<<<>>>><<>>>><<<>><<<<>>>><>>>><<<<>><<<>><<<<>>>><<<>>>><<<>>><>>><>>><<>>><<<>>><>>><<<<>>><<><<>><<<>>><<<<>>>><<<>><<<>><>><<<><<>>><<<<>>><<>><<>>><<<>>>><<<<><><>>>><>><<<>><<<<><<<<>>><<<>><<>><<>><<><<<>>>><>><<><<<<>><<<<>>>><<<<><<<<>><<<<><<<<><<<>>>><<<><<<>>><<<>>><>>><<>>>><<<<>>><>>>><<<<>><<<<><<<><<<><>><><<<<>>>><<>>>><<<<>>>><<>>><<<<>>>><<<>>>><<>>><<<<><<>><<<<><>>>><<<>>>><><<<<>>><<>>><<><<<>>><<>><>>><<>><><<<<>><<<<><<>><>>>><>>>><<>><<<>>><<<<>><<<>>><<<><<<<>>><>>><<>><<<<>>>><<<<>>>><>>>><<<<><>>>><<<<><<<>><<<<>>><>>><>><><<<>>>><<>>><<<<>><>><<<>><<>>><>><<>>><<>>><>>><>><<>>>><>>>><<<>>>><<<>>><>><<><<>>>><<>><<<<>>><<><><<>>>><<<>><>>>><<<>><<<<>>><<<<>><<>>>><<<>>><<<<><>>>><><<<>>><<<<>>><<<<>>><>>><<>><<<>><>>>><<>><><<<>><<<<>>>><>>><>><<<<>>>><<<>>>><>><<><<<>><<<<>>><<<>>><>>>><>><>><<<<>>><>>>><<>>>><<<>><<<>>>><<<<><<<><><<<<>><<>>><<>><<>>>><<>>>><<<<>>><<<>><<>>>><<<<><<<>>><<<>>>><<<<>><<<>><<<<>>><<>>><>>>><<<<>><<>>>><<<<>>>><<>><<<<>>>><><<<<>><<<>><<<<>>><<<<><<<<>>>><<<>><<>>>><<<<>><<>>><<>><<>>>><<<<>><<>><<>>>><<><<<>><<>>>><<>><<>>>><<<>><>>>><<>><<<>>>><<>>>><>><<>><<<>>>><<<<>><>>>><<<>><><<<<>>><<<<><<>>>><<<<>><<<<>><<<<>>><>>><<<>>>><<<<><<<<>>><<<>>>><<>><><<<<>>>><<<<>>><<<>>>><<>>>><<>><<<<>><<<<><<<<>>>><<<>>>><<<>>><>>><<<><<<>><<>>><>><<<>>>><<>><>>><<<>>><<<>>><<<>>>><>>><<<>>>><>>>><<<>>><<<<>>><<<>><<>><<<<>><<<<><<<<>><<<<><<<<>>>><<>>><<<>><<>><<<>><<>>>><>><<<<>><<<>>>><>><>>>><<>>><<>><<<<>>>><<><<<>><<<<><>>><<<<>>><>><<<>>><>><>><>><<<>><>>>><<<>>>><<<>>>><<<>><<<<>><<<><<>>><<>>>><>>>><<>>>><<<<>>><<<<><>>><<<<><<>>><<<>>>><<<><<<<>><<<>>>><<<><<<<>><<><<<<>><<<<>><<<>><><><<<>>><><<<>><<<<>><<><<>>><<<>>><<<>><<><<<>>>><<<>><<>>>><<<>>>><>><<<>>><<<<>><><<>>><<<<>>>><><<<<>><<<>><<<>>><<<>><<<<><<><<<><<<>><<<<>>>><<<<>>>><<<<>>>><<>>><<<<>>><<<<>>>><<<><<>><<<<>>>><<>><<>>><<>>>><<>><<>><><<><<><<<<>><>><<><>>><>>><<<><<<>>>><<>>>><<>>><<<>>><<<<>><>>>><<>>>><<><<<<>><<<<>><><<<>><>>>><<><>>>><<<>>>><><>><<><<><><<<>>><<<<>>><<>><<<>>>><<<<>>><<><>>><><<>>><<<><<><<<>>>><<<<>>><<<>><<<<>><<><<<>><<<<><<<><<<<>>><<<<>>><<<>>>><<><<<>>><>>><<>>><><<>>><<<<><<>><<>>><>>><<<<>>><<<<><<>>>><<>>><<>><<<<><>>>><<<>>>><>><<>>>><<<<>>><>><>>><<<>>>><><<<<>><<<<>>>><<<>>><>>>><<>>><><<<<><>>>><<>>><<<<><<<>>>><<<>>><>>>><><>>><<>><>><>><<<>><<<<>><<<><<>>>><<<>>><><<<>><>><<<<>><<>><<<<><<<<>><<<>>>><<<><<<<>>>><<>><<><<<>><>><<<>>>><<>>><<<<><><>>>><>><<<>><<<<><>>><<>>><><>>><<>>>><><<<>>>><<<>>><<<>><>><<>>>><><>><<<<>>>><<<><<<>>><<><<<<><<>>>><<>>>><<<>>>><>>>><<<<>>>><<<>><<<<>><<><<<<>>><<<>>>><<<><<<>>>><<<<>>>><><>>>><<><<<<><>>>><>>>><<<>>><<>>>><<<>>><<<>><<>>>><<>><><<>><<>>><><>>>><>>><<<<>>><<<<>><<<<>>>><<>><<<<>><<<<><<<<>>>><><<><<>>>><<<>><<<<><<<>>>><<<>>><<><<>><<<><<<>>>><><><<<<><><>>>><<>>>><<>><<<<><<<>>><<<><>>><<<<>>>><<<<>>>><<>>><<<<>><<<<>>><<>><<>>><<>>>><>>>><<>>>><>>>><<>>>><<<>><>>><<><<<>>><>><<>><<>>><>>>><<><<>>>><>><<<>>><<>>>><<<<>>><><<>><<>><<<><<<<>>>><<>>>><>><<<>><<<<>><<<<>>><<>>>><<<<><<>><<<><><<>>>><<>>>><<<>>><<<<>>>><<<>>>><<>>><<<<>>><>><<<<>>><<<>>><<>>>><>>>><<<><<>><>><<<<><>>>><>><<<>><>><>>><<<><>>><<<<>><<>>>><>>>><<>>>><<>><<<<>>>><>><<<>>><>>><<>>>><<>>><<<<>>>><<<><<<<>><<<>>>><<<><<<>><<<<>><<<>>><<>>>><><<<>><<<>><<<><<<<>>><<<>>>><<<<>><<>><<<>>><<<<><<<><><<<<>><<<<>>>><<><<<>><<>>>><<>>>><>>>><<>><<<><<>>>><<>><<<>>><<<>>>><>>><>><<<>>>><<><<>>><<<<>>>><<<><<<>><>><<><<<<><<<>>>><<<>>>><<<<>>>><<>><<>><<<<>><>><<>><<>>><<<<>><<<<><<<<>>>><<<>>>><<><<>>>><<<>>><>><<>>><<<<>><<<>>>><<<<>>>><<<<>>>><>>>><><><<<<><<>>><>><<>>><>><<<>>>><<>><<>>><<<<>><<><<>><<<<>><<<<><<<<>><<<<>><<<>><<<<>>><>>>><<>>><>>><<>>>><<>><<<>>><<<>>><<<>><>><<<<>>>><>><>>><<>>>><<>>>><<<>>>><<<<>>>><<<><<<>>>><>><<<>>>><<><<<><>><<<<>>>><<<>>><<<>>>><<><<>>>><<<<>>`
}
