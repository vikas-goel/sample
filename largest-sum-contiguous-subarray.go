// Largest Sum Contiguous Subarray
// https://www.geeksforgeeks.org/largest-sum-contiguous-subarray/

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage:", os.Args[0])
	}

	var lscFinal, lscCurrent []int
	maxFinal, maxCurrent := -999999, 0
	sequence := make([]int, len(os.Args)-1)
	for i := 0; i < len(sequence); i++ {
		sequence[i], _ = strconv.Atoi(os.Args[i+1])
		if maxCurrent+sequence[i] > sequence[i] {
			maxCurrent += sequence[i]
			lscCurrent = append(lscCurrent, sequence[i])
		} else {
			maxCurrent = sequence[i]
			lscCurrent = []int{ sequence[i] }
		}

		if maxCurrent >= maxFinal {
			maxFinal = maxCurrent
			lscFinal = lscCurrent
		}
	}

	fmt.Println(maxFinal, lscFinal)
}
