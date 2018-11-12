// https://www.geeksforgeeks.org/subset-sum-problem-dp-25/

package main

import "fmt"

func subsetSum(set []int, sum, index int, subset *[]int) bool {
	if sum > 0 && index == 0 {
		return false
	} else if sum == 0 {
		return true
	} else if subsetSum(set, sum, index-1, subset) {
		return true
	} else if sum-set[index-1] >= 0 {
		if subsetSum(set, sum-set[index-1], index-1, subset) {
			*subset = append(*subset, set[index-1])
			return true
		}
	}

	return false
}

func main() {
	var subset []int
	set := []int{3, 8, 34, 4, 12, 7, 5, 2}
	fmt.Println(subsetSum(set, 33, len(set), &subset), subset)
}
