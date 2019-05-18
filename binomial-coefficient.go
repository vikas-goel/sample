// https://www.geeksforgeeks.org/dynamic-programming-set-9-binomial-coefficient/

package main

import "fmt"

// Recursive binomial coefficient.
func rBC(n, k int, coeff [][]int) int {
	if k == 0 || k == n {
		coeff[n][k] = 1
	} else if coeff[n][k] == 0 {
		// Include nth elem + Exclude nth elem.
		coeff[n][k] = rBC(n-1, k-1, coeff) + rBC(n-1, k, coeff)
	}

	return coeff[n][k]
}

// Iterative binomial coefficient.
func iBC(n, k int) int {
	coeff := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		coeff[i] = make([]int, k+1)
		coeff[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			coeff[i][j] = coeff[i-1][j-1] + coeff[i-1][j]
		}
	}

	return coeff[n][k]
}

func main() {
	n, k := 4, 2
	coeff := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		coeff[i] = make([]int, k+1)
	}

	fmt.Println(rBC(n, k, coeff))
	fmt.Println(iBC(n, k))
}
