package undirected

import (
	"fmt"
	"github.com/pradykaushik/data-structures/graphs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUndirectedGraph(t *testing.T) {
	ug := NewUndirectedGraph(13).(*UndirectedGraph)
	assert.NotNil(t, ug)
	assert.Equal(t, ug.numVertices, 13)
	assert.Equal(t, ug.numEdges, 0)
}

func TestGetV(t *testing.T) {
	ug := NewUndirectedGraph(13).(*UndirectedGraph)
	assert.NotNil(t, ug)
	assert.Equal(t, ug.GetV(), 13)
	assert.Equal(t, ug.GetE(), 0)
}

func getUndirectedGraph(t *testing.T) graphs.Graph {
	ug := NewUndirectedGraph(13)
	assert.NotNil(t, ug)
	// Edges to add.
	// The pairs are taken from https://algs4.cs.princeton.edu/41graph/.
	var pairs [][]int
	pairs = append(pairs, []int{0, 5})
	pairs = append(pairs, []int{4, 3})
	pairs = append(pairs, []int{0, 1})
	pairs = append(pairs, []int{9, 12})
	pairs = append(pairs, []int{6, 4})
	pairs = append(pairs, []int{5, 4})
	pairs = append(pairs, []int{0, 2})
	pairs = append(pairs, []int{11, 12})
	pairs = append(pairs, []int{9, 10})
	pairs = append(pairs, []int{0, 6})
	pairs = append(pairs, []int{7, 8})
	pairs = append(pairs, []int{9, 11})
	pairs = append(pairs, []int{5, 3})

	for _, p := range pairs {
		ug.AddEdge(p[0], p[1])
	}

	return ug
}

func TestAddEdge(t *testing.T) {
	ug := NewUndirectedGraph(13).(*UndirectedGraph)
	assert.NotNil(t, ug)
	// Edges to add.
	// The pairs are taken from https://algs4.cs.princeton.edu/41graph/.
	var pairs [][]int
	pairs = append(pairs, []int{0, 5})
	pairs = append(pairs, []int{4, 3})
	pairs = append(pairs, []int{0, 1})
	pairs = append(pairs, []int{9, 12})
	pairs = append(pairs, []int{6, 4})
	pairs = append(pairs, []int{5, 4})
	pairs = append(pairs, []int{0, 2})
	pairs = append(pairs, []int{11, 12})
	pairs = append(pairs, []int{9, 10})
	pairs = append(pairs, []int{0, 6})
	pairs = append(pairs, []int{7, 8})
	pairs = append(pairs, []int{9, 11})
	pairs = append(pairs, []int{5, 3})

	for _, p := range pairs {
		assert.True(t, ug.AddEdge(p[0], p[1]))
	}
	assert.False(t, ug.AddEdge(14, 0))
}

func compareArrays(t *testing.T, arr1, arr2 []int) {
	assert.Equal(t, len(arr1), len(arr2))
	for i := 0; i < len(arr1); i++ {
		assert.Equal(t, arr1[i], arr2[i])
	}
}

func compare2dArrays(t *testing.T, arr1, arr2 [][]int) {
	assert.Equal(t, len(arr1), len(arr2))
	for i := 0; i < len(arr1); i++ {
		assert.Equal(t, len(arr1[i]), len(arr2[i]))
		for j := 0; j < len(arr1[i]); j++ {
			assert.Equal(t, arr1[i][j], arr2[i][j])
		}
	}
}

func TestAdjacent(t *testing.T) {
	ug := getUndirectedGraph(t)
	var expectedAdjLists = [][]int{
		{6, 2, 1, 5},
		{0},
		{0},
		{5, 4},
		{5, 6, 3},
		{3, 4, 0},
		{0, 4},
		{8},
		{7},
		{11, 10, 12},
		{9},
		{9, 12},
		{11, 9},
	}

	assert.Equal(t, ug.GetV(), 13)
	for i := 0; i < 13; i++ {
		adjL, validVertex := ug.Adjacent(i)
		assert.True(t, validVertex)
		compareArrays(t, expectedAdjLists[i], adjL)
		t.Log(fmt.Sprintf("adj[%d] = %v", i, adjL))
	}
}

func TestDegrees(t *testing.T) {
	ug := getUndirectedGraph(t)
	var expectedDegrees = []int{4, 1, 1, 2, 3, 3, 2, 1, 1, 3, 1, 2, 2}
	assert.Equal(t, ug.GetV(), 13)
	for i := 0; i < 13; i++ {
		var deg, indeg, outdeg int
		var validVertex bool
		deg, validVertex = ug.Degree(i)
		assert.True(t, validVertex)
		assert.Equal(t, expectedDegrees[i], deg)

		indeg, validVertex = ug.Degree(i)
		assert.True(t, validVertex)
		assert.Equal(t, expectedDegrees[i], indeg)

		outdeg, validVertex = ug.Degree(i)
		assert.True(t, validVertex)
		assert.Equal(t, expectedDegrees[i], outdeg)
	}
}

func TestPrintDfs(t *testing.T) {
	ug := getUndirectedGraph(t)
	dfsResult := ug.Dfs()
	compareArrays(t, []int{0, 6, 4, 5, 3, 2, 1, 7, 8, 9, 11, 12, 10}, dfsResult)
}

func TestPrintBfs(t *testing.T) {
	ug := getUndirectedGraph(t)
	bfsResult := ug.Bfs()
	compareArrays(t, []int{0, 6, 2, 1, 5, 4, 3, 7, 8, 9, 11, 10, 12}, bfsResult)
}

func TestConnectedVertices(t *testing.T) {
	ug := getUndirectedGraph(t)
	for i := 0; i < 7; i++ {
		connectedV, validVertex := ug.ConnectedVertices(i)
		assert.True(t, validVertex)
		assert.ElementsMatch(t, []int{0, 1, 2, 3, 4, 5, 6}, connectedV)
	}

	for i := 7; i < 9; i++ {
		connectedV, validVertex := ug.ConnectedVertices(i)
		assert.True(t, validVertex)
		assert.ElementsMatch(t, []int{7, 8}, connectedV)
	}

	for i := 9; i < 13; i++ {
		connectedV, validVertex := ug.ConnectedVertices(i)
		assert.True(t, validVertex)
		assert.ElementsMatch(t, []int{9, 10, 11, 12}, connectedV)
	}
}

func TestFindPathV2(t *testing.T) {
	ug := getUndirectedGraph(t)
	for i := 0; i < 13; i++ {
		for j := 0; j < 13; j++ {
			path, found := ug.FindPathV2(i, j)
			if found {
				t.Logf("path[%d -> %d] => %v", i, j, path)
			}
		}
	}
}
