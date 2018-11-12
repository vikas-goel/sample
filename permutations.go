package main

import (
	"fmt"
	"log"
	"os"
)

/*
type Interface interface {
	Copy() *interface{}
	Swap(x, y int)
}
*/

type charset struct {
	elems []byte
}

func Permutations(allSet *[]charset, set charset, start, end int) {
	if start == end {
		//newSet := set.Copy()
		//*allSet = append(*allSet, *newSet)
		fmt.Printf("%c\n", set.elems)
	} else {
		for i := start; i <= end; i++ {
			set.Swap(start, i)
			Permutations(allSet, set, start+1, end)
			set.Swap(start, i)
		}
	}
}

func (c *charset) Copy() *charset {
	newSet := new(charset)
	newSet.elems = make([]byte, len(c.elems))
	copy(newSet.elems, c.elems)
	return newSet
}

func (c *charset) Swap(x, y int) {
	//fmt.Printf("%c\n", c.elems)
	c.elems[x], c.elems[y] = c.elems[y], c.elems[x]
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "<string>")
	}

	orig := []byte(os.Args[1])
	//fmt.Printf("%c\n", orig)

	var all []charset
	set := charset{orig}

	Permutations(&all, set, 0, len(orig)-1)
}
