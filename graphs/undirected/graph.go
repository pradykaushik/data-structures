package undirected

import (
	"bytes"
	"fmt"
	"github.com/pradykaushik/data-structures/graphs"
	"github.com/pradykaushik/data-structures/linkedlist"
	"github.com/pradykaushik/data-structures/queue"
	"github.com/pradykaushik/data-structures/queue/fifo"
)

// UndirectedGraph is a Graph where the edges are not directed.
// This means that one could traverse from a vertex to another vertex as long as there
// is at least one edge connecting the two.
type UndirectedGraph struct {
	gph         []*linkedlist.LinkedList // using adjancency list representation.
	numVertices int
	numEdges    int
}

// Vertex implements util.Value and represents a vertex in the graph.
type Vertex int

func (v Vertex) Get() interface{} {
	return int(v)
}

// NewUndirectedGraph creates an undirected with the provided number of vertices.
// Note that this undirected graph will have no edges to begin with.
func NewUndirectedGraph(v int) graphs.Graph {
	g := &UndirectedGraph{
		gph:         make([]*linkedlist.LinkedList, v),
		numVertices: v,
		numEdges:    0,
	}

	for i := 0; i < g.numVertices; i++ {
		g.gph[i] = linkedlist.New()
	}

	return g
}

func (g UndirectedGraph) GetV() int {
	return g.numVertices
}

func (g UndirectedGraph) GetE() int {
	return g.numEdges
}

func (g *UndirectedGraph) AddEdge(v1 int, v2 int) bool {
	// As this is an undirected graph, we need to add v1-v2 and v2-v1.
	if (v1 >= len(g.gph)) || (v2 >= len(g.gph)) {
		return false
	}

	g.gph[v1].AddToFront(Vertex(v2))
	g.gph[v2].AddToFront(Vertex(v1))
	return true
}

func (g UndirectedGraph) Adjacent(v int) ([]int, bool) {
	// we need to convert from []util.Value to []int.
	var adjVertices []int
	if v >= len(g.gph) {
		return adjVertices, false
	}

	for _, adjVertex := range g.gph[v].SerializeIntoArray() {
		adjVertices = append(adjVertices, adjVertex.Get().(int))
	}

	return adjVertices, true
}

func (g UndirectedGraph) Degree(v int) (int, bool) {
	if v >= len(g.gph) {
		return -1, false
	}
	return g.gph[v].Size(), true
}

func (g UndirectedGraph) InDegree(v int) (int, bool) {
	return g.Degree(v)
}

func (g UndirectedGraph) OutDegree(v int) (int, bool) {
	return g.Degree(v)
}

func (g UndirectedGraph) String() string {
	var buf = new(bytes.Buffer)
	for v, adjL := range g.gph {
		buf.WriteString(fmt.Sprintf("%d => ", v))
		var adjVertices []int
		for _, connectedV := range adjL.SerializeIntoArray() {
			adjVertices = append(adjVertices, connectedV.Get().(int))
		}
		buf.WriteString(fmt.Sprintf("%v\n", adjVertices))
	}
	return buf.String()
}

func (g UndirectedGraph) Dfs() []int {
	if len(g.gph) == 0 {
		return []int{}
	}
	var visited = make(map[int]struct{})
	var result = make([]int, 0, 0)
	for v := range g.gph {
		if _, ok := visited[v]; !ok {
			g.dfs(v, &visited, &result)
		}
	}
	return result
}

func (g UndirectedGraph) dfs(
	v int,
	visited *map[int]struct{},
	result *[]int) {

	(*result) = append(*result, v)
	(*visited)[v] = struct{}{}
	for _, adjV := range g.gph[v].SerializeIntoArray() {
		if _, ok := (*visited)[adjV.Get().(int)]; !ok {
			g.dfs(adjV.Get().(int), visited, result)
		}
	}
}

func (g UndirectedGraph) Bfs() []int {
	if len(g.gph) == 0 {
		return []int{}
	}

	var nextV = fifo.NewLinearQueueArr(len(g.gph))
	var visited = make(map[int]struct{})
	var result = make([]int, 0, 0)

	for i := 0; i < len(g.gph); i++ {
		if _, ok := visited[i]; !ok {
			nextV.Enqueue(Vertex(i))
			g.bfs(nextV, &visited, &result)
		}
	}
	return result
}

func (g UndirectedGraph) bfs(
	next queue.Queue,
	visited *map[int]struct{},
	result *[]int) {

	for !next.IsEmpty() {
		nextV, _ := next.Dequeue()
		v := nextV.Get().(int)
		(*visited)[v] = struct{}{} // marking visited.
		(*result) = append(*result, v)
		for _, adjV := range g.gph[v].SerializeIntoArray() {
			if _, ok := (*visited)[adjV.Get().(int)]; !ok {
				next.Enqueue(adjV)
				(*visited)[adjV.Get().(int)] = struct{}{} // marking visited.
			}
		}
	}
}
