package main

import "fmt"

func main() {
	num := 7
	f := make([]int, num+1)

	f[0], f[1] = 1, 1

	for i := 2; i <= num; i++ {
		f[i] = f[i-1] + f[i-2]
		fmt.Println("{", i, ",", f[i], "}")
	}
}
