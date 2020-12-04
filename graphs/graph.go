package graphs

// Graph defines an API for an undirected graph.
// API taken from https://algs4.cs.princeton.edu/41graph/.
//
// Total number of vertices = V.
// Total number of edges = E.
// Vertices are numbered from 0 to V-1.
type Graph interface {
	GetV() int
	GetE() int
	// AddEdge adds an edge to connect the two vertices.
	// Return false if vertex does not exist.
	AddEdge(int, int) bool
	// Adjacent returns the list of vertices adjacent to the provided one.
	Adjacent(int) ([]int, bool)
	// Degree returns the number of edges incident on the given vertex.
	Degree(int) (int, bool)
	// InDegree returns the number of edges directed into the vertex.
	// Note that for an undirected graph, indegree = degree.
	InDegree(int) (int, bool)
	// OutDegree returns the number of edges directed out of the vertex.
	// Note that for an undirected graph, outdegree, degree.
	OutDegree(int) (int, bool)
	// String representation of the graph.
	String() string

	// Traversals.
	Dfs() []int
	Bfs() []int

	// Graph based algorithms.
	ConnectedVertices(int) ([]int, bool)
	FindPath(int, int) ([]int, bool)
	FindPathV2(int, int) ([]int, bool)
}
