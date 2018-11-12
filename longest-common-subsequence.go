// Longest Common Subsequence
// https://www.geeksforgeeks.org/longest-common-substring-dp-29/

package main

import (
	"fmt"
	"log"
	"os"
)

func lcs(str [][][]byte, x, y []byte, lx, ly int) int {
	if lx == 0 || ly == 0 {
		return 0
	} else if x[lx-1] == y[ly-1] {
		len := lcs(str, x, y, lx-1, ly-1)
		str[lx][ly] = append(str[lx-1][ly-1], x[lx-1])
		return 1 + len
	} else {
		len1 := lcs(str, x, y, lx, ly-1)
		len2 := lcs(str, x, y, lx-1, ly)
		if len1 > len2 {
			str[lx][ly] = make([]byte, len(str[lx][ly-1]))
			copy(str[lx][ly], str[lx][ly-1])
			return len1
		} else {
			str[lx][ly] = make([]byte, len(str[lx-1][ly]))
			copy(str[lx][ly], str[lx-1][ly])
			return len2
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage")
	}

	X, Y := []byte(os.Args[1]), []byte(os.Args[2])

	strXY := make([][][]byte, len(X)+1)
	for i := 0; i < len(X)+1; i++ {
		strXY[i] = make([][]byte, len(Y)+1)
	}

	fmt.Printf("LCS(%c, %c) = {%d, %c}\n", X, Y,
		lcs(strXY, X, Y, len(X), len(Y)), strXY[len(X)][len(Y)])
}
