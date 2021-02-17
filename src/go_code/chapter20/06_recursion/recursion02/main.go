package main

import "fmt"

func solve(m *[8][7]int, x, y int) bool {
	// if [6][5] == 2 then done
	if m[6][5] == 2 {
		return true
	} else {
		if m[x][y] == 0 { // if x, y is unknown the could have a try
			// assume this point is good
			m[x][y] = 2 // need to check and decide following down right up left

			if solve(m, x + 1, y){ // up
				return true
			} else if solve(m, x, y + 1) { // down
				return true
			} else if solve(m, x - 1, y)  { // left
				return true
			} else if solve(m, x, y - 1) { // right
				return true
			} else { // dead
				m[x][y] = 3
				return false
			}

		} else { // start point is wall
			return false
		}
	}
}

func main() {
	// make a map with 2D array
	// 1 -> wall
	// 0 -> unknown point
	// 2 -> could path
	// 3 -> known dead path
	var Map [8][7]int
	for i := 0; i < len(Map) - 1; i++ {
		Map[0][i] = 1
		Map[i][0] = 1
		Map[len(Map) - 1][i] = 1
		Map[i][len(Map[0]) - 1] = 1
	}
	Map[3][1] = 1
	Map[3][2] = 1
	Map[1][2] = 1
	Map[2][2] = 1

	for i := 0; i < len(Map); i++ {
		for j := 0; j < len(Map[0]); j ++ {
			fmt.Printf("%v   ", Map[i][j])
		}
		fmt.Println()
	}

	fmt.Println()
	solve(&Map, 1, 1)

	for i := 0; i < len(Map); i++ {
		for j := 0; j < len(Map[0]); j ++ {
			fmt.Printf("%v   ", Map[i][j])
		}
		fmt.Println()
	}

}
