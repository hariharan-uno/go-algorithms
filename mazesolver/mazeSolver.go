package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Maze struct {
	x [5][5]int
}

var Goal = [2]int{4, 3}
var Start = [2]int{0, 1}
var sol = [][]int{}
var m = Maze{}
var defaultBlockers = "[0,0]:[4,0];[0,4]:[4,4];[0,2]:[0,3];[4,1]:[4,2]"

var givenBlockers = "[2,2]:[2,3];[2,2]:[3,2]" //Update this variable according to the given inputs.

func main() {
	addBlockers(defaultBlockers, &m) //Core definition
	addBlockers(givenBlockers, &m)   //Given blockers

	if findPath(Start[0], Start[1]) {
		fmt.Println(sol)
		fmt.Println()
		prettyPrint(&m)
	} else {
		fmt.Println("Blocked")
	}
}

//findPath is a recursive function with the following algorithm:
//
//  FIND-PATH(x, y)
//
//  if (x,y outside maze) return false
//  if (x,y is goal) return true
//  if (x,y not open) return false
//  mark x,y as part of solution path
//  if (FIND-PATH(North of x,y) == true) return true
//  if (FIND-PATH(East of x,y) == true) return true
//  if (FIND-PATH(South of x,y) == true) return true
//  if (FIND-PATH(West of x,y) == true) return true
//  unmark x,y as part of solution path
//  return false
//
func findPath(x int, y int) bool {
	if notInRange(x, y) {
		return false
	}
	if isGoal(x, y) {
		m.x[x][y] = 3
		sol = append(sol, []int{x, y})
		return true
	}
	if notOpen(x, y) {
		return false
	}
	m.x[x][y] = 3
	sol = append(sol, []int{x, y})
	if findPath(x, y+1) {
		return true
	}
	if findPath(x+1, y) {
		return true
	}
	if findPath(x, y-1) {
		return true
	}
	if findPath(x-1, y) {
		return true
	}
	m.x[x][y] = 0
	sol = sol[:len(sol)-1]
	return false
}

//notOpen checks whether a coordinate is open or not.
//It returns true if the coordinate is not open and false if it is open.
func notOpen(x int, y int) bool {
	if m.x[x][y] != 0 {
		return true
	}
	return false
}

//isGoal checks whether a coordinate is the Goal or not.
func isGoal(x int, y int) bool {
	if x == Goal[0] && y == Goal[1] {
		return true
	}
	return false
}

//notInRange checks whether a coordinate is not in range or is in range.
//It returns true if the coordinate is not in range and false if it is in range.
func notInRange(x int, y int) bool {
	if x > 4 || y > 4 || x < 0 || y < 0 {
		return true
	}
	return false
}

//prettyPrint prints the 2D array into a visual grid
func prettyPrint(m *Maze) {
	for i := 4; i >= 0; i-- {
		for j := 0; j < 5; j++ {
			fmt.Printf("%v ", m.x[j][i])
		}
		fmt.Println()
	}
}

//addBlockers adds blockers to a Maze
func addBlockers(s string, m *Maze) {
	blocks := strings.Split(s, ";")
	for _, block := range blocks {
		x0, _ := strconv.Atoi(string(block[1]))
		y0, _ := strconv.Atoi(string(block[3]))
		x1, _ := strconv.Atoi(string(block[7]))
		y1, _ := strconv.Atoi(string(block[9]))
		if x0 == x1 {
			for y0 <= y1 {
				m.x[x0][y0] = 1
				y0++
			}
		}
		if y0 == y1 {
			for x0 <= x1 {
				m.x[x0][y0] = 1
				x0++
			}
		}
	}
}
