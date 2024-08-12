package main

import (
	"fmt"
)

// Node represents a node in the graph.
type Node struct {
	ID       string
	Category int
}

// Edge represents an edge between two nodes.
type Edge struct {
	Source string
	Target string
}

// EdgeNode represents an edge in the adjacency list.
type EdgeNode struct {
	Adjvex int
	Next   *EdgeNode
}

// VertexNode represents a vertex in the adjacency list.
type VertexNode struct {
	Data      Node
	FirstEdge *EdgeNode
}

// findNodeIndex returns the index of a node with the given ID in the nodes slice.
func findNodeIndex(nodes []Node, id string) int {
	for i, node := range nodes {
		if node.ID == id {
			return i
		}
	}
	return -1
}

// dfs performs a depth-first search on the graph.
func dfs(adjList []VertexNode, currentIndex int, targetIndex, visited map[int]bool, traversedEdges *[]Edge, originalEdges []Edge, originalNodes []Node, modelNodes map[int]bool) {
	if targetIndex[currentIndex] {
		visited[currentIndex] = true
		return
	}

	if modelNodes[currentIndex] {
		visited[currentIndex] = true
		for _, edge := range originalEdges {
			if findNodeIndex(originalNodes, edge.Target) == currentIndex {
				*traversedEdges = append(*traversedEdges, edge)
			}
		}
		return
	}

	visited[currentIndex] = true
	temp := adjList[currentIndex].FirstEdge
	for temp != nil {
		if !visited[temp.Adjvex] {
			sourceIndex := currentIndex
			targetIndex1 := temp.Adjvex
			for _, edge := range originalEdges {
				if findNodeIndex(originalNodes, edge.Source) == sourceIndex && findNodeIndex(originalNodes, edge.Target) == targetIndex1 {
					*traversedEdges = append(*traversedEdges, edge)
				} else if findNodeIndex(originalNodes, edge.Source) == targetIndex1 && findNodeIndex(originalNodes, edge.Target) == sourceIndex {
					*traversedEdges = append(*traversedEdges, edge)
				}
			}
			dfs(adjList, temp.Adjvex, targetIndex, visited, traversedEdges, originalEdges, originalNodes, modelNodes)
		}
		temp = temp.Next
	}
}

// getNoAssociatedNodes identifies nodes and edges that are not associated with target nodes.
func getNoAssociatedNodes(nodes []Node, edges []Edge, baseNode Node, targetNodes []Node, modelNodes []Node) ([]string, []Edge) {
	size := len(nodes)
	adjList := make([]VertexNode, size)

	for i := 0; i < size; i++ {
		adjList[i].Data = nodes[i]
	}

	// Build adjacency list
	for _, edge := range edges {
		sourceIndex := findNodeIndex(nodes, edge.Source)
		targetIndex := findNodeIndex(nodes, edge.Target)

		if sourceIndex != -1 && targetIndex != -1 {
			newEdge1 := &EdgeNode{Adjvex: targetIndex, Next: adjList[sourceIndex].FirstEdge}
			adjList[sourceIndex].FirstEdge = newEdge1

			newEdge2 := &EdgeNode{Adjvex: sourceIndex, Next: adjList[targetIndex].FirstEdge}
			adjList[targetIndex].FirstEdge = newEdge2
		}
	}

	// Perform DFS from baseNode
	visitedFromBase := make(map[int]bool)
	targetNodesInt := make(map[int]bool)
	modelNodesInt := make(map[int]bool)

	for _, targetNode := range targetNodes {
		x := findNodeIndex(nodes, targetNode.ID)
		targetNodesInt[x] = true
	}

	for _, modelNode := range modelNodes {
		x := findNodeIndex(nodes, modelNode.ID)
		modelNodesInt[x] = true
	}

	var traversedEdges []Edge
	dfs(adjList, findNodeIndex(nodes, baseNode.ID), targetNodesInt, visitedFromBase, &traversedEdges, edges, nodes, modelNodesInt)

	// Collect nodes that are not associated with any targetNodes
	var resultNodes []string
	for index := 0; index < size; index++ {
		if visitedFromBase[index] {
			resultNodes = append(resultNodes, adjList[index].Data.ID)
		}
	}

	return resultNodes, traversedEdges
}

func main() {
	nodes := []Node{
		{"0", 0}, {"1", 0}, {"2", 0}, {"3", 1}, {"4", 0},
		{"5", 0}, {"6", 1}, {"7", 2}, {"8", 1}, {"9", 1},
		{"10", 1}, {"11", 1}, {"12", 1}, {"13", 1}, {"14", 1},
		{"15", 1}, {"16", 0}, {"17", 0}, {"18", 0}, {"19", 0},
		{"20", 0}, {"21", 0}, {"22", 0}, {"23", 0}, {"24", 0},
		{"25", 0}, {"26", 0}, {"27", 0}, {"28", 0}, {"29", 0},
		{"30", 0}, {"31", 0}, {"32", 0}, {"33", 0}, {"34", 0},
		{"35", 0}, {"36", 0}, {"37", 0}, {"38", 0}, {"39", 0},
		{"40", 0}, {"41", 0}, {"42", 0},
	}

	edges := []Edge{
		{"7", "0"}, {"7", "1"}, {"2", "7"}, {"3", "2"}, {"4", "7"},
		{"3", "4"}, {"5", "7"}, {"6", "5"}, {"5", "8"}, {"5", "9"},
		{"5", "10"}, {"5", "11"}, {"5", "12"}, {"5", "13"}, {"5", "14"},
		{"5", "15"}, {"8", "16"}, {"8", "17"}, {"8", "18"}, {"8", "19"},
		{"6", "20"}, {"20", "8"}, {"2", "8"}, {"6", "21"}, {"21", "8"},
		{"7", "22"}, {"7", "23"}, {"9", "24"}, {"9", "25"}, {"9", "26"},
		{"9", "27"}, {"4", "9"}, {"2", "9"}, {"10", "28"}, {"10", "29"},
		{"10", "30"}, {"10", "31"}, {"20", "10"}, {"2", "10"}, {"21", "10"},
		{"11", "32"}, {"11", "33"}, {"11", "34"}, {"11", "35"}, {"20", "11"},
		{"2", "11"}, {"12", "36"}, {"12", "37"}, {"12", "38"}, {"12", "39"},
		{"20", "12"}, {"2", "12"}, {"13", "40"}, {"13", "41"}, {"13", "42"},
		{"13", "43"}, {"45", "46"}, {"46", "13"}, {"6", "47"}, {"47", "13"},
		{"3", "48"}, {"48", "13"}, {"14", "49"}, {"14", "50"}, {"14", "51"},
		{"14", "52"}, {"4", "14"}, {"2", "14"},
	}

	var baseNode Node
	var targetNodes, modelNodes []Node

	for _, node := range nodes {
		switch node.Category {
		case 2:
			baseNode = node
		case 1:
			targetNodes = append(targetNodes, node)
		case 0:
			modelNodes = append(modelNodes, node)
		}
	}

	noAssociatedNodes, edgeRelationships := getNoAssociatedNodes(nodes, edges, baseNode, targetNodes, modelNodes)

	fmt.Println("剩余节点:")
	for _, node := range noAssociatedNodes {
		fmt.Print(node, " ")
	}
	fmt.Println()

	fmt.Println("节点关系:")
	for _, edge := range edgeRelationships {
		fmt.Printf("%s -> %s\n", edge.Source, edge.Target)
	}
}
