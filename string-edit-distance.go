// https://www.geeksforgeeks.org/edit-distance-dp-5/

package main

import (
	"fmt"
	"log"
	"os"
)

func minimumOf(nums ...int) (min int) {
	min = 999999
	for _, i := range nums {
		if i < min {
			min = i
		}
	}

	return
}

func getDistanceBetween(str1, str2 []byte, lstr1, lstr2 int) (d int) {
	if lstr1 == 0 {
		d = lstr2
	} else if lstr2 == 0 {
		d = lstr1
	} else if str1[lstr1-1] == str2[lstr2-1] {
		d = getDistanceBetween(str1, str2, lstr1-1, lstr2-1)
	} else {
		d = 1 + minimumOf(
			getDistanceBetween(str1, str2, lstr1, lstr2-1),
			getDistanceBetween(str1, str2, lstr1-1, lstr2),
			getDistanceBetween(str1, str2, lstr1-1, lstr2-1))
	}

	return
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage", os.Args[0], "<string1> <string2>")
	}

	str1 := []byte(os.Args[1])
	str2 := []byte(os.Args[2])

	oper := getDistanceBetween(str1, str2, len(str1), len(str2))
	fmt.Printf("%c -> %c = %v operations.\n", str1, str2, oper)
}
