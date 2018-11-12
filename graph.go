package main

import "fmt"

type list struct {
	vertex int
	next *list
}

type graph struct {
	directed bool
	vertices int
	edges []*list
}

func graphNew(directed bool, vertices int) (g *graph) {
	if vertices < 1 {
		return
	}

	g = new(graph)
	g.directed, g.vertices = directed, vertices
	g.edges = make([]*list, g.vertices)

	return
}

func (g *graph) edgeAdd(v1, v2 int) bool {
	if v1 < 0 || v1 >= g.vertices || v2 < 0 || v2 >= g.vertices {
		return false
	}

	v1v2 := new(list)
	v1v2.vertex = v2
	v1v2.next = g.edges[v1]
	g.edges[v1] = v1v2

	if g.directed {
		return true
	}

	v2v1 := new(list)
	v2v1.vertex = v1
	v2v1.next = g.edges[v2]
	g.edges[v2] = v2v1

	return true
}

func (g *graph) bfs(vertex int) *[]int {
	if vertex >= g.vertices || vertex < 0 {
		return nil
	}

	visited := make([]bool, g.vertices)
	queue := make([]int, 0, g.vertices)
	queue = append(queue, vertex)

	for top := 0; top < len(queue); top++ {
		vertex = queue[top]
		if visited[vertex] {
			continue
		}

		visited[vertex] = true
		for e := g.edges[vertex]; e != nil; e = e.next {
			if !visited[e.vertex] {
				queue = append(queue, e.vertex)
			}
		}
	}

	return &queue
}

func (g *graph) dfs(vertex int, visited *[]bool, queue *[]int) *[]int {
	if vertex >= g.vertices || vertex < 0 {
		return nil
	}

	if queue == nil {
		visited = new([]bool)
		*visited = make([]bool, g.vertices)
		queue = new([]int)
		*queue = make([]int, 0)
	}

	*queue = append(*queue, vertex)
	(*visited)[vertex] = true

	for e := g.edges[vertex]; e != nil; e = e.next {
		if !(*visited)[e.vertex] {
			g.dfs(e.vertex, visited, queue)
		}
	}

	return queue
}

func (g *graph) cycle() bool {
	if g == nil || g.vertices < 2 {
		return false
	}

	parent := make([]int, g.vertices)
	for i := 0; i < len(parent); i++ {
		parent[i] = -1
	}

	parentOf := func(vertex int) int {
		for ; parent[vertex] != -1; vertex = parent[vertex] {}
		return vertex
	}

	eVisited := map[int]map[int]bool{}
	for i := 0; i < g.vertices; i++ {
		eVisited[i] = map[int]bool{}
	}

	for src := 0; src < g.vertices; src++ {
		sParent := parentOf(src)
		for e := g.edges[src]; e != nil; e = e.next {
			dst := e.vertex
			if eVisited[src][dst] {
				continue
			}

			dParent := parentOf(dst)
			if sParent == dParent {
				return true
			}
			parent[dParent] = sParent

			if !g.directed {
				eVisited[src][dst] = true
				eVisited[dst][src] = true
			}
		}
	}

	fmt.Print("{ ")
	for i := 0; i < len(parent); i++ {
		if parent[i] == -1 {
			fmt.Print(i, " ")
		}
	}
	fmt.Print("}")

	return false
}

func (g *graph) print() {
	for i := 0; i < g.vertices; i++ {
		fmt.Printf("[%2d] -> { ", i)
		for e := g.edges[i]; e != nil; e = e.next {
			fmt.Print(e.vertex, " ")
		}
		fmt.Println("}")
	}
}

func main() {
	g := graphNew(false, 4)
	g.edgeAdd(0, 1)
	g.edgeAdd(0, 2)
	//g.edgeAdd(1, 2)
	g.edgeAdd(2, 0)
	g.edgeAdd(2, 3)
	//g.edgeAdd(3, 3)

	g.print()

	fmt.Println("BFS:", *g.bfs(0), *g.bfs(1), *g.bfs(2))

	fmt.Println("DFS:", *g.dfs(0, nil, nil), *g.dfs(1, nil, nil), *g.dfs(2, nil, nil))
	fmt.Print("Graph has a cycle: ")
	fmt.Println("", g.cycle())
}
