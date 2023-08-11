package main

import (
	"fmt"
)

type Graph struct {
	Nodes int
	Edges map[int][]int
}

// NewGraph создает новый граф
func NewGraph(nodes int) *Graph {
	return &Graph{
		Nodes: nodes,
		Edges: map[int][]int{},
	}
}

// AddEdges добавляет ребра
func (g *Graph) AddEdge(parent, child int) {
	g.Edges[parent] = append(g.Edges[parent], child)
}

// AddReversEdge добавялет двухсторонее ребро, если граф неориентирован
func (g *Graph) AddReversEdge(parent, child int) {
	g.Edges[parent] = append(g.Edges[parent], child)
	g.Edges[child] = append(g.Edges[child], parent)
}

// DFSUtil рекурсивен, он помечает уже пройденные вершины
// и вызывает сам себя столько раз, сколько ребер у нас есть в графе
func (g *Graph) DFSUtil(node int, visited map[int]int) {
	if visited[node] == 0 {
		visited[node] = 1
		fmt.Println(node)
		for _, n := range g.Edges[node] {
			if visited[n] == 0 {
				g.DFSUtil(n, visited)
				//visited[node] = 2
			}

		}
		visited[node] = 2
	} else if visited[node] == 1 {
		fmt.Println("Found cycle!!!")
	}

}

// DFS инициализирует карту посещенных вершин и впервые вызывает DFSUtil
func (g *Graph) DFS(startNode int, visited map[int]int) map[int]int {
	//visited := make(map[int]bool)
	g.DFSUtil(startNode, visited)
	return visited
}

func main() {

	var nodes int
	fmt.Print("Enter number of nodes ")
	fmt.Scanf("%d \n", &nodes)
	g := NewGraph(nodes)
	visited := make(map[int]int)
	for i := 1; i < nodes+1; i++ {
		visited[i] = 0
	}

	var edges int
	var parent, child int
	var graphType int
	fmt.Print("Enter number of edges ")
	fmt.Scanf("%d \n", &edges)
	fmt.Print("Choose graph type: 1 - oriented / 2 - unoriented")
	fmt.Scanf("%d \n", &graphType)
	switch graphType {
	case 1:
		fmt.Print("You chose oriented graph")
		for i := 0; i < edges; i++ {
			fmt.Println("Add edge by entering parent and child ")
			fmt.Print("Add parent ")
			fmt.Scanf("%d \n", &parent)
			fmt.Print("Add child ")
			fmt.Scanf("%d \n", &child)
			g.AddEdge(parent, child)
		}
	
	case 2:
		fmt.Print("You chose unoriented graph")
		for i := 0; i < edges; i++ {
			fmt.Println("Add edge by entering parent and child ")
			fmt.Print("Add parent ")
			fmt.Scanf("%d \n", &parent)
			fmt.Print("Add child ")
			fmt.Scanf("%d \n", &child)
			g.AddReversEdge(parent, child)
	}

	// g.AddEdge(1, 2)
	// g.AddEdge(1, 3)
	// g.AddEdge(3, 4)
	// g.AddEdge(3, 6)
	// g.AddEdge(4, 5)

	visitedNodes := g.DFS(1, visited)

	fmt.Println(visitedNodes)

	for n := range visitedNodes {
		if visitedNodes[n] == 0 {
			fmt.Println("skipped node:", n)
		}
	}

}
