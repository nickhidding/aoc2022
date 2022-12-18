package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

type Valve struct {
	name    string
	flow    int
	tunnels []string
	routes  map[string]int
}

type State struct {
	valve    Valve
	minutes  int
	released int
	opened   []Open
}

type Open struct {
	minute int
	valve  string
}

func main() {
	mainStart := time.Now()
	valvesList := parseInput(strings.Split(input(), "\n"))
	calculateRoutes(&valvesList)

	fmt.Println("--- Part 1 ---")
	var start Valve = valvesList[slices.IndexFunc(valvesList, func(v Valve) bool { return v.name == "AA" })]
	stack := make([]State, 0)
	stack = append(stack, State{valve: start, minutes: 30, released: 0, opened: make([]Open, 0)})
	var mostReleasedState State

	for len(stack) != 0 {
		state := stack[0]
		stack = stack[1:]
		terminal := true
		for target, distance := range state.valve.routes {
			targetValve := valvesList[slices.IndexFunc(valvesList, func(v Valve) bool { return v.name == target })]
			var alreadyOpened bool = slices.IndexFunc(state.opened, func(o Open) bool { return o.valve == target }) != -1
			if state.minutes-distance > 0 && !alreadyOpened {
				terminal = false
				opened := make([]Open, len(state.opened))
				copy(opened, state.opened)
				opened = append(opened, Open{valve: target, minute: state.minutes})
				extraReleased := (state.minutes - distance) * targetValve.flow
				stack = append(stack, State{valve: targetValve, minutes: state.minutes - distance, released: state.released + extraReleased, opened: opened})
			}
		}
		if terminal {
			if state.released > mostReleasedState.released {
				mostReleasedState = state
			}
		}
	}
	fmt.Println("Most pressure that can be released:", mostReleasedState.released)
	fmt.Println("Path")
	for i := 0; i < len(mostReleasedState.opened); i++ {
		fmt.Printf("Valve %s opened at minute %2d\n", mostReleasedState.opened[i].valve, mostReleasedState.opened[i].minute)
	}

	elapsed := time.Since(mainStart)
	fmt.Printf("Part 1 took %s\n", elapsed)
	fmt.Println()
	fmt.Println("--- Part 2 ---")
}

func parseInput(valves []string) []Valve {
	valvesList := make([]Valve, 0)
	for _, valve := range valves {
		r := regexp.MustCompile(`Valve (.*) has flow rate=(\d+); tunnel([s]?) lead([s]?) to valve([s]?) (.*)`)
		matches := r.FindStringSubmatch(valve)
		name := matches[1]
		flow, _ := strconv.Atoi(string(matches[2]))
		tunnels := matches[6]
		valvesList = append(valvesList, Valve{name: name, flow: flow, tunnels: strings.Split(tunnels, ", ")})
	}
	return valvesList
}

func calculateRoutes(valves *[]Valve) {
	for i, v := range *valves {
		routes := make(map[string]int, 0)
		stack := make([]Valve, 0)
		stack = append(stack, v)
		distances := make(map[string]int, 0)
		visited := make([]Valve, 0)
		distances[v.name] = 0
		for len(stack) > 0 && len(routes) < len(*valves) {
			cur := stack[0]
			stack = stack[1:]
			visited = append(visited, cur)
			curDist := distances[cur.name]
			if cur.name != v.name && cur.flow > 0 {
				routes[cur.name] = curDist + 1
			}
			for _, v2 := range cur.tunnels {
				var alreadyVisisted bool = slices.IndexFunc(visited, func(c Valve) bool { return c.name == v2 }) != -1
				if !alreadyVisisted {
					stack = append(stack, (*valves)[slices.IndexFunc(*valves, func(c Valve) bool { return c.name == v2 })])
					distances[v2] = curDist + 1
				}
			}
		}
		(*valves)[i].routes = routes
	}
}

func input() string {
	return `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II`
}

func puzzleInput() string {
	return `Valve SY has flow rate=0; tunnels lead to valves GW, LW
Valve TS has flow rate=0; tunnels lead to valves CC, OP
Valve LU has flow rate=0; tunnels lead to valves PS, XJ
Valve ND has flow rate=0; tunnels lead to valves EN, TL
Valve PD has flow rate=0; tunnels lead to valves TL, LI
Valve VF has flow rate=0; tunnels lead to valves LW, RX
Valve LD has flow rate=0; tunnels lead to valves AD, LP
Valve DG has flow rate=0; tunnels lead to valves DR, SS
Valve IG has flow rate=8; tunnels lead to valves AN, YA, GA
Valve LK has flow rate=0; tunnels lead to valves HQ, LW
Valve TD has flow rate=14; tunnels lead to valves BG, CQ
Valve CQ has flow rate=0; tunnels lead to valves TD, HD
Valve AZ has flow rate=0; tunnels lead to valves AD, XW
Valve ZU has flow rate=0; tunnels lead to valves TL, AN
Valve HD has flow rate=0; tunnels lead to valves BP, CQ
Valve FX has flow rate=0; tunnels lead to valves LW, XM
Valve CU has flow rate=18; tunnels lead to valves BX, VA, RX, DF
Valve SS has flow rate=17; tunnels lead to valves DG, ZD, ZG
Valve BP has flow rate=19; tunnels lead to valves HD, ZD
Valve DZ has flow rate=0; tunnels lead to valves XS, CC
Valve PS has flow rate=0; tunnels lead to valves GH, LU
Valve TA has flow rate=0; tunnels lead to valves LI, AA
Valve BG has flow rate=0; tunnels lead to valves TD, ZG
Valve WP has flow rate=0; tunnels lead to valves OB, AA
Valve XS has flow rate=9; tunnels lead to valves EN, DZ
Valve AA has flow rate=0; tunnels lead to valves WG, GA, VO, WP, TA
Valve LW has flow rate=25; tunnels lead to valves LK, FX, SY, VF
Valve AD has flow rate=23; tunnels lead to valves DF, GW, AZ, LD, FM
Valve EN has flow rate=0; tunnels lead to valves ND, XS
Valve ZG has flow rate=0; tunnels lead to valves SS, BG
Valve LI has flow rate=11; tunnels lead to valves YA, XM, TA, PD
Valve VO has flow rate=0; tunnels lead to valves AA, OD
Valve AN has flow rate=0; tunnels lead to valves IG, ZU
Valve GH has flow rate=15; tunnels lead to valves VA, PS
Valve OP has flow rate=4; tunnels lead to valves AJ, TS, FM, BX, NM
Valve BX has flow rate=0; tunnels lead to valves OP, CU
Valve RX has flow rate=0; tunnels lead to valves CU, VF
Valve FM has flow rate=0; tunnels lead to valves OP, AD
Valve OB has flow rate=0; tunnels lead to valves WP, XW
Valve CC has flow rate=3; tunnels lead to valves QS, LP, DZ, OD, TS
Valve LP has flow rate=0; tunnels lead to valves LD, CC
Valve NM has flow rate=0; tunnels lead to valves WH, OP
Valve HQ has flow rate=0; tunnels lead to valves XW, LK
Valve GW has flow rate=0; tunnels lead to valves SY, AD
Valve QS has flow rate=0; tunnels lead to valves CC, XW
Valve DF has flow rate=0; tunnels lead to valves AD, CU
Valve XM has flow rate=0; tunnels lead to valves LI, FX
Valve VA has flow rate=0; tunnels lead to valves CU, GH
Valve GA has flow rate=0; tunnels lead to valves IG, AA
Valve YA has flow rate=0; tunnels lead to valves LI, IG
Valve XW has flow rate=20; tunnels lead to valves OB, HQ, QS, WH, AZ
Valve XJ has flow rate=24; tunnel leads to valve LU
Valve AJ has flow rate=0; tunnels lead to valves WG, OP
Valve WH has flow rate=0; tunnels lead to valves XW, NM
Valve TL has flow rate=13; tunnels lead to valves PD, DR, ZU, ND
Valve OD has flow rate=0; tunnels lead to valves CC, VO
Valve ZD has flow rate=0; tunnels lead to valves SS, BP
Valve DR has flow rate=0; tunnels lead to valves DG, TL
Valve WG has flow rate=0; tunnels lead to valves AJ, AA`
}
