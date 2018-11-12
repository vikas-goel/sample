package main

import "fmt"
import "math"

type ScanGraph interface {
	Shorter(srcWeight, edgeWeight, tgtWeight int) bool
	Update(tgtWeight *int, srcWeight, edgeWeight int)
}

type ScanPrims struct{}
func (g ScanPrims) Shorter(srcWeight, edgeWeight, tgtWeight int) bool {
	return edgeWeight < tgtWeight
}

func (g ScanPrims) Update(tgtWeight *int, srcWeight, edgeWeight int) {
	if tgtWeight == nil {
		return
	}

	*tgtWeight = edgeWeight
}

type ScanDijkstras struct{}
func (g ScanDijkstras) Shorter(srcWeight, edgeWeight, tgtWeight int) bool {
	return (srcWeight+edgeWeight) < tgtWeight
}

func (g ScanDijkstras) Update(tgtWeight *int, srcWeight, edgeWeight int) {
	if tgtWeight == nil {
		return
	}

	*tgtWeight = edgeWeight + srcWeight
}

func minimum(set []bool, array []int, len int) (index int) {
	if len == 0 {
		return -1
	}

	min := math.MaxInt32
	for i := 0; i < len; i++ {
		if !set[i] && array[i] < min {
			index = i
			min = array[index]
		}
	}

	return
}

func printPath(path []int, len, target int) {
	if len == 0 {
		return
	}

	fmt.Print("{", target)
	for ; path[target] != -1; target = path[target] {
		fmt.Print(", ", path[target])
	}
	fmt.Println("}")
}

func traverse(g ScanGraph, graph [][]int, V, src int, path []int) {
	if V == 0 {
		return
	}

	visited := make([]bool, V)
	weight := make([]int, V)
	for i := 0; i < V; i++ {
		weight[i] = math.MaxInt32
	}

	weight[src] = 0
	path[src] = -1

	for i := 0; i < V; i++ {
		u := minimum(visited, weight, V)
		visited[u] = true

		for v := 0; v < V; v++ {
			if !visited[v] && graph[u][v] != 0 &&
				g.Shorter(weight[u], graph[u][v], weight[v]) {
				g.Update(&weight[v], weight[u], graph[u][v])
				path[v] = u
			}
		}
	}
}

func main() {
	V := 9
	path := make([]int, V)
	graph := [][]int{ {0, 4, 0, 0, 0, 0, 0, 8, 0 },
			{4, 0, 8, 0, 0, 0, 0, 11, 0},
			{0, 8, 0, 7, 0, 4, 0, 0, 2},
			{0, 0, 7, 0, 9, 14, 0, 0, 0},
			{0, 0, 0, 9, 0, 10, 0, 0, 0},
			{0, 0, 4, 14, 10, 0, 2, 0, 0},
			{0, 0, 0, 0, 0, 2, 0, 1, 6},
			{8, 11, 0, 0, 0, 0, 1, 0, 7},
			{0, 0, 2, 0, 0, 0, 6, 7, 0} }

	traverse(ScanDijkstras{}, graph, V, 0, path)
	printPath(path, V, 3)
	printPath(path, V, 4)
	printPath(path, V, 8)

	traverse(ScanPrims{}, graph, V, 0, path)
	printPath(path, V, 3)
	printPath(path, V, 4)
	printPath(path, V, 8)
}
