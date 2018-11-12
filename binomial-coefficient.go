// https://www.geeksforgeeks.org/dynamic-programming-set-9-binomial-coefficient/

package main

import "fmt"

func binCoeff(n, k int, coeff [][]int) int {
	if k == 0 || k == n {
		coeff[n][k] = 1
	} else if coeff[n][k] == 0 {
		// Include nth elem + Exclude nth elem.
		coeff[n][k] = binCoeff(n-1, k-1, coeff) + binCoeff(n-1, k, coeff)
	}

	return coeff[n][k]
}

func main() {
	n, k := 4, 2
	coeff := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		coeff[i] = make([]int, k+1)
	}

	fmt.Println(binCoeff(n, k, coeff))
}
