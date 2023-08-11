package main

//
//import (
//	"fmt"
//	"strconv"
//)
//
//type Node struct {
//	// Key идентификатор вершины
//	Key int
//	// Вершины будут описывать вершины, связанные с этой
//	// Ключ будет значением Key подключенной вершины
//	// со значением, являющимся указателем на него
//	Vertices map[int]*Node
//}
//
//// Функция-конструктор для Node
//func NewNode(key int) *Node {
//	return &Node{
//		Key:      key,
//		Vertices: map[int]*Node{},
//	}
//}
//
//func (v *Node) String() string {
//	s := strconv.Itoa(v.Key) + ":"
//
//	for _, neighbor := range v.Vertices {
//		s += " " + strconv.Itoa(neighbor.Key)
//	}
//
//	return s
//}
//
//type Graph struct {
//	// Вершины описывает все вершины, содержащиеся в графе
//	// Ключ будет значением Key подключенной вершины
//	// со значением, являющимся указателем на него
//	Vertices map[int]*Node
//	// Решает граф или орграф
//	directed bool
//}
//
//func NewDirectedGraph() *Graph {
//	return &Graph{
//		Vertices: map[int]*Node{},
//		directed: true,
//	}
//}
//
//func NewUndirectedGraph() *Graph {
//	return &Graph{
//		Vertices: map[int]*Node{},
//	}
//}
//
//// AddVertex создает новую вершину с заданным
//// ключом и добавляем его на граф
//func (g *Graph) AddVertex(key int) {
//	v := NewNode(key)
//	g.Vertices[key] = v
//}
//
//// Метод AddEdge добавляет ребро между двумя вершинами графа.
//func (g *Graph) AddEdge(k1, k2 int) {
//	v1 := g.Vertices[k1]
//	v2 := g.Vertices[k2]
//
//	// вернуть ошибку, если одна из вершин не существует
//	if v1 == nil || v2 == nil {
//		panic("not all vertices exist")
//	}
//
//	// ничего не делать, если вершины уже соединены
//	if _, ok := v1.Vertices[v2.Key]; ok {
//		return
//	}
//
//	// Добавляем направленное ребро между v1 и v2
//	// Если граф неориентированный, добавляем ребро v2 к v1,
//	// делая ребро между v1 и v2 двунаправленное
//	v1.Vertices[v2.Key] = v2
//	if !g.directed && v1.Key != v2.Key {
//		v2.Vertices[v1.Key] = v1
//	}
//
//	// Добавляем вершины на карту вершин графа.
//	g.Vertices[v1.Key] = v1
//	g.Vertices[v2.Key] = v2
//}
//
//func (g *Graph) String() string {
//	s := ""
//	i := 0
//	for _, v := range g.Vertices {
//		if i != 0 {
//			s += "\n"
//		}
//		s += v.String()
//		i++
//	}
//	return s
//}
//
//func DFS(g *Graph, startNode *Node, visitCb func(int)) {
//	// мы ведем карту посещенных узлов, чтобы предотвратить посещение одного и того же
//	// узел более одного раза
//	visited := map[int]bool{}
//
//	if startNode == nil {
//		return
//	}
//	visited[startNode.Key] = true
//	visitCb(startNode.Key)
//
//	// для каждой из соседних вершин рекурсивно вызвать функцию
//	// если он еще не был посещен
//	for _, v := range startNode.Vertices {
//		if visited[v.Key] {
//			continue
//		}
//		DFS(g, v, visitCb)
//	}
//}
//
//func main() {
//	g := NewDirectedGraph()
//
//	g.AddVertex(1)
//	g.AddVertex(2)
//	g.AddVertex(3)
//	g.AddVertex(4)
//	g.AddVertex(5)
//	g.AddVertex(6)
//	g.AddVertex(7)
//	g.AddVertex(8)
//	g.AddVertex(9)
//	g.AddVertex(10)
//
//	g.AddEdge(1, 9)
//	g.AddEdge(1, 5)
//	g.AddEdge(1, 2)
//	g.AddEdge(2, 2)
//	g.AddEdge(3, 4)
//	g.AddEdge(5, 6)
//	g.AddEdge(5, 8)
//	g.AddEdge(6, 7)
//	g.AddEdge(9, 10)
//
//	visitedOrder := []int{}
//	cb := func(i int) {
//		visitedOrder = append(visitedOrder, i)
//	}
//	DFS(g, g.Vertices[1], cb)
//
//	// add assertions here
//	fmt.Println(visitedOrder)
//}
