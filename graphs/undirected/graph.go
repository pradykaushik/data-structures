package undirected

import (
	"bytes"
	"fmt"
	"github.com/pradykaushik/data-structures/graphs"
	"github.com/pradykaushik/data-structures/linkedlist"
	"github.com/pradykaushik/data-structures/queue"
	"github.com/pradykaushik/data-structures/queue/fifo"
	"github.com/pradykaushik/data-structures/stack"
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

// All the vertices visited in a dfs are connected to the source vertex.
func (g UndirectedGraph) ConnectedVertices(source int) ([]int, bool) {
	if source >= len(g.gph) {
		return []int{}, false
	}

	var connected = make([]int, 0, 0)
	var visited = make(map[int]struct{})
	g.connectedVertices(source, &visited, &connected)
	return connected, true
}

func (g UndirectedGraph) connectedVertices(
	v int,
	visited *map[int]struct{},
	connected *[]int) {

	(*connected) = append(*connected, v)
	(*visited)[v] = struct{}{}
	for _, adjV := range g.gph[v].SerializeIntoArray() {
		if _, ok := (*visited)[adjV.Get().(int)]; !ok {
			g.connectedVertices(adjV.Get().(int), visited, connected)
		}
	}
}

// Creating the path while traversing the graph.
// If going down a path wasn't fruitful, then removing the corresponding vertices from path path.
// This way, by the time we're done traversing the graph, we'll have the path - (V + E).
//
// Important note - the max path length = V.
func (g UndirectedGraph) FindPathV2(source, dest int) ([]int, bool) {
	if (source >= len(g.gph)) || (dest >= len(g.gph)) {
		return []int{}, false
	}

	if source == dest {
		return []int{source}, true
	}

	// Using a circular queue for constant time pop() operation.
	// Could also use doubly linkedlist (might actually be more efficient).
	var path = stack.NewArrayStack(len(g.gph))
	var found = false
	var visited = make(map[int]struct{})
	g.findPathV2(source, dest, source, &visited, path, &found)
	var pathArr []int
	for !path.IsEmpty() {
		v, _ := path.Pop()
		pathArr = append([]int{v}, pathArr...)
	}
	return pathArr, found
}

func (g UndirectedGraph) findPathV2(
	source, dest int,
	curV int,
	visited *map[int]struct{},
	path stack.Stack,
	found *bool) {

	(*visited)[curV] = struct{}{} // marking as visited.
	if curV == dest {
		// we have found the path.
		(*found) = true
		path.Push(curV)
		return
	}

	// keeping track of whether there aren't any more vertices to explore.
	var moreToExplore = false
	var toExplore []int
	for _, adjV := range g.gph[curV].SerializeIntoArray() {
		if _, ok := (*visited)[adjV.Get().(int)]; !ok {
			moreToExplore = true
			toExplore = append(toExplore, adjV.Get().(int))
		}
	}

	if moreToExplore {
		path.Push(curV)
		for _, v := range toExplore {
			if _, ok := (*visited)[v]; !ok {
				g.findPathV2(source, dest, v, visited, path, found)
				if *found {
					// this exploration was fruitful.
					return
				}
			}
		}
		// If here, then none of the explorations from curV were fruitful.
		path.Pop()
	}
}
