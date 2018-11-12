// https://www.geeksforgeeks.org/dynamic-programming-set-7-coin-change/

package main

import (
        "fmt"
)

func count(coin []int, sum int) int {
	numCoins := len(coin)
	change := make([][]int, sum+1)

	for i := 0; i <= sum; i++ {
		change[i] = make([]int, numCoins)
	}

	for i := 0; i < numCoins; i++ {
		change[0][i] = 1
	}

	for i := 1; i <= sum; i++ {
		for j := 0; j < numCoins; j++ {
			// Include jth coin.
			if i - coin[j] >= 0 {
				change[i][j] += change[i-coin[j]][j]
			}

			// Exclude jth coin.
			if j >= 1 {
				change[i][j] += change[i][j-1]
			}
		}
	}

	return change[sum][numCoins-1]
}

func count2(coin []int, sum int) int {
	numCoins := len(coin)
	change := make([]int, sum+1)

	change[0] = 1
	for i := 0; i < numCoins; i++ {
		for j := coin[i]; j <= sum; j++ {
			change[j] += change[j-coin[i]]
		}
	}

	return change[sum]
}

func main() {
	coin := []int{1,2,3,5}
	fmt.Println(count(coin, 10))
	fmt.Println(count2(coin, 10))
}
