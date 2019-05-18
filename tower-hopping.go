package main

import "fmt"
import "math"

func main() {
	//towers := []int{4, 2, 0, 0, 2, 0}
	towers := []int{1, 3, 5, 3, 1, 2, 0}
	fmt.Printf("\nHop tower %v? %v.\n", towers, canHopTowers(towers))
}

func canHopTowers(tower []int) int {
	var numTowers int = len(tower)
	var hopTowersFrom func(int)

	hops := make([]int, numTowers)
	path := make([]int, numTowers)
	seen := make([]bool, numTowers)

	for i := 0; i < numTowers; i++ {
		hops[i] = math.MaxInt32
	}

	// A function to compute hop state for given tower recursively.
	hopTowersFrom = func (source int) {
		// If already computed this tower, then return the state.
		if seen[source] {
			return
		}

		seen[source] = true
		for jumps := tower[source]; jumps > 0; jumps-- {
			next := source+jumps

			if next >= numTowers {
				hops[source] = 1
				return
			}

			hopTowersFrom(next)
			if hops[next] < hops[source]-1 {
				// The next tower is either crossing the given
				// set or hoppable. So, the source tower is
				// hoppable.
				hops[source] = hops[next]+1
				path[source] = next
			}
		}

		return
	}

	hopTowersFrom(0)
	for i := 0; path[i] != 0; i = path[i] {
		fmt.Printf("{%v,%v} ", i, path[i])
	}

	return hops[0]
}
