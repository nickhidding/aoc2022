package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	startY int
	endY   int
}

func main() {
	sensors := strings.Split(input(), "\n")
	checkY := 10
	maxXY := 20

	sensors = strings.Split(puzzleInput(), "\n")
	checkY = 2000000
	maxXY = 4000000

	tunnelMap := make(map[int]map[int]string, 0) // xy
	ranges := make(map[int][]Range, 0)

	for _, sensor := range sensors {
		fmt.Println(sensor)
		r := regexp.MustCompile(`Sensor at x=([-]?\d+), y=([-]?\d+): closest beacon is at x=([-]?\d+), y=([-]?\d+)`)
		matches := r.FindStringSubmatch(sensor)
		sensorX, _ := strconv.Atoi(string(matches[1]))
		sensorY, _ := strconv.Atoi(string(matches[2]))
		beaconX, _ := strconv.Atoi(string(matches[3]))
		beaconY, _ := strconv.Atoi(string(matches[4]))

		initXMap(&tunnelMap, sensorX)
		tunnelMap[sensorX][sensorY] = "S"
		initXMap(&tunnelMap, beaconX)
		tunnelMap[beaconX][beaconY] = "B"

		// Area that cannot contain beacon
		manhattanDistance := int(math.Abs(float64(sensorX-beaconX))) + int(math.Abs(float64(sensorY-beaconY)))

		// Mark manhattan distance
		i := 0
		for x := sensorX - manhattanDistance; x <= sensorX+manhattanDistance; x++ {

			// Used for part 1
			dist := int(math.Abs(float64(sensorX-x))) + int(math.Abs(float64(sensorY-checkY)))
			if dist <= manhattanDistance {
				initXMap(&tunnelMap, x)
				if tunnelMap[x][checkY] == "" {
					tunnelMap[x][checkY] = "#"
				}
			}

			// Used for part 2
			ranges[x] = append(ranges[x], Range{startY: sensorY - i, endY: sensorY + i})
			if x < sensorX {
				i++
			} else {
				i--
			}
		}
	}

	fmt.Println("--- Part 1 ---")
	for y := 0; y <= 22; y++ {
		fmt.Printf("%3d ", y)
		for x := -4; x <= 26; x++ {
			if tunnelMap[x][y] == "" {
				fmt.Print(".")
			} else {
				fmt.Print(tunnelMap[x][y])
			}
		}
		fmt.Println()
	}

	positions := 0
	for i := range tunnelMap {
		if tunnelMap[i][checkY] == "#" {
			positions++
		}
	}
	fmt.Println("In the row where y =", checkY, " there are", positions, "positions where a beacon cannot be present.")
	fmt.Println()

	fmt.Println("--- Part 2 ---")
	for x := 0; x <= maxXY; x++ {
		r := ranges[x]
		sort.SliceStable(r, func(i, j int) bool {
			return r[i].startY < r[j].startY
		})
		previousEnd := 0
		for _, k := range r {
			if previousEnd < k.startY-1 {
				fmt.Println("found possible beacon position at x", x, "y", k.startY-1)
				fmt.Println("Frequency =", (x*4000000)+k.startY-1)
			}
			if k.endY >= previousEnd {
				previousEnd = k.endY
			}
		}
	}
}

func initXMap(tunnelMap *map[int]map[int]string, x int) {
	if (*tunnelMap)[x] == nil {
		(*tunnelMap)[x] = make(map[int]string, 0)
	}
}

func input() string {
	return `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`
}

func puzzleInput() string {
	return `Sensor at x=489739, y=1144461: closest beacon is at x=-46516, y=554951
Sensor at x=2543342, y=3938: closest beacon is at x=2646619, y=229757
Sensor at x=3182359, y=3999986: closest beacon is at x=3142235, y=3956791
Sensor at x=3828004, y=1282262: closest beacon is at x=3199543, y=2310713
Sensor at x=871967, y=3962966: closest beacon is at x=-323662, y=4519876
Sensor at x=1323641, y=2986163: closest beacon is at x=2428372, y=3303736
Sensor at x=2911492, y=2576579: closest beacon is at x=3022758, y=2461675
Sensor at x=3030965, y=2469848: closest beacon is at x=3022758, y=2461675
Sensor at x=3299037, y=3402462: closest beacon is at x=3142235, y=3956791
Sensor at x=1975203, y=1672969: closest beacon is at x=1785046, y=2000000
Sensor at x=3048950, y=2452864: closest beacon is at x=3022758, y=2461675
Sensor at x=336773, y=2518242: closest beacon is at x=1785046, y=2000000
Sensor at x=1513936, y=574443: closest beacon is at x=2646619, y=229757
Sensor at x=3222440, y=2801189: closest beacon is at x=3199543, y=2310713
Sensor at x=2838327, y=2122421: closest beacon is at x=2630338, y=2304286
Sensor at x=2291940, y=2502068: closest beacon is at x=2630338, y=2304286
Sensor at x=2743173, y=3608337: closest beacon is at x=2428372, y=3303736
Sensor at x=3031202, y=2452943: closest beacon is at x=3022758, y=2461675
Sensor at x=3120226, y=3998439: closest beacon is at x=3142235, y=3956791
Sensor at x=2234247, y=3996367: closest beacon is at x=2428372, y=3303736
Sensor at x=593197, y=548: closest beacon is at x=-46516, y=554951
Sensor at x=2612034, y=2832157: closest beacon is at x=2630338, y=2304286
Sensor at x=3088807, y=3929947: closest beacon is at x=3142235, y=3956791
Sensor at x=2022834, y=2212455: closest beacon is at x=1785046, y=2000000
Sensor at x=3129783, y=3975610: closest beacon is at x=3142235, y=3956791
Sensor at x=3150025, y=2333166: closest beacon is at x=3199543, y=2310713
Sensor at x=3118715, y=2376161: closest beacon is at x=3199543, y=2310713
Sensor at x=3951193, y=3181929: closest beacon is at x=4344952, y=3106256
Sensor at x=2807831, y=2401551: closest beacon is at x=2630338, y=2304286
Sensor at x=3683864, y=2906786: closest beacon is at x=4344952, y=3106256
Sensor at x=2723234, y=3206978: closest beacon is at x=2428372, y=3303736
Sensor at x=3047123, y=3891244: closest beacon is at x=3142235, y=3956791
Sensor at x=3621967, y=3793314: closest beacon is at x=3142235, y=3956791
Sensor at x=2384506, y=1814055: closest beacon is at x=2630338, y=2304286
Sensor at x=83227, y=330275: closest beacon is at x=-46516, y=554951
Sensor at x=3343176, y=75114: closest beacon is at x=2646619, y=229757`
}
