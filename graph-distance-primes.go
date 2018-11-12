//https://www.geeksforgeeks.org/shortest-path-reach-one-prime-changing-single-digit-time/

package main

import "fmt"
import "log"
import "os"
import "strconv"

type graph struct {
	v int
	edge [][]int
}

func newGraph(V int) (g *graph) {
	g = new(graph)
	g.v = V
	g.edge = make([][]int, V)
	for i := 0; i < V; i++ {
		g.edge[i] = make([]int, 0, 1)
	}

	return
}

func (g *graph) addEdge(v1, v2 int) {
	if g == nil {
		return
	}

	g.edge[v1] = append(g.edge[v1], v2)
	g.edge[v2] = append(g.edge[v2], v1)
}

func (g *graph) getPath(v1, v2 int) []int {
	visited := make([]bool, g.v)
	queue := make([]int, g.v)
	parent := make([]int, g.v)

	queue[0] = v1
	parent[v1] = -1
	visited[v1] = true
	qlen := 1
	for qpos := 0; qpos != qlen; qpos++ {
		vert := queue[qpos]

		for _, e := range g.edge[vert] {
			if !visited[e] {
				queue[qlen] = e
				parent[e] = vert
				visited[e] = true
				qlen++
			}
		}
	}

	return parent
}


func getPrimes() []int {
	num := 9999
	notPrime := make([]bool, num+1)

	for n := 2; n*n < num; n++ {
		if notPrime[n] {
			continue
		}

		for i := n*n; i <= num; i += n {
			notPrime[i] = true
		}
	}

	primes := make([]int, 0, 5)
	for n := 1000; n < num; n++ {
		if !notPrime[n] {
			primes = append(primes, n)
		}
	}

	return primes
}

func connected(num1, num2 int) bool {
	diff := 0

	for div := 1000; div >= 1; div /= 10 {
		n1, n2 := (num1%(div*10))/div, (num2%(div*10))/div
		if n1 != n2 {
			diff++
		}
	}

	if diff == 1 {
		return true
	}

	return false
}

func printPath(p1, p2 int) {
	idx1, idx2 := -1, -1
	primes := getPrimes()
	g := newGraph(len(primes))

	for i := 0; i < len(primes); i++ {
		for j := i+1; j < len(primes); j++ {
			if connected(primes[i], primes[j]) {
				g.addEdge(i, j)
			}
		}

		if primes[i] == p1 {
			idx1 = i
		} else if primes[i] == p2 {
			idx2 = i
		}
	}

	path := g.getPath(idx1, idx2)
	fmt.Print("[ ")
	for i := idx2; i != -1; i = path[i] {
		fmt.Print(primes[i], " ")
	}
	fmt.Println("]")
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Usage: ", os.Args[0], " <prime-1> <prime-2>")
	}

	prime1, _ := strconv.Atoi(os.Args[1])
	prime2, _ := strconv.Atoi(os.Args[2])

	printPath(prime1, prime2)
}
