// Package dijkstra provides an implementation of Dijkstra's algorithm for finding the shortest path in a weighted graph.
package dijkstra

import (
	"errors"
	"fmt"

	"github.com/pkg/errors"
)

// Graph represents a weighted graph with nodes and edges.
type Graph struct {
	Nodes map[string]map[string]float64
}

// NewGraph creates a new weighted graph.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]map[string]float64),
	}
}

// AddNode adds a node to the graph with an empty adjacency list.
func (g *Graph) AddNode(node string) {
	g.Nodes[node] = make(map[string]float64)
}

// AddEdge adds an edge between two nodes with a specified weight.
func (g *Graph) AddEdge(from, to string, weight float64) error {
	if from != to {
		g.Nodes[from][to] = weight
	}
	return nil
}

// Dijkstra implements Dijkstra's algorithm for finding the shortest path between two nodes.
func Dijkstra(graph *Graph, start, end string) (map[string]float64, error) {
	// Initialize the distance map with infinity for all nodes except the start node.
	distances := make(map[string]float64)
	for node := range graph.Nodes {
		distances[node] = float64(Infinity)
	}
	distances[start] = 0

	// Initialize the priority queue with the start node.
	pq := make(PriorityQueue, len(graph.Nodes))
	for node := range graph.Nodes {
		pq = append(pq, &Node{node, 0})
	}

	for len(pq) > 0 {
		// Extract the node with the minimum distance from the priority queue.
		current := pq.ExtractMin()
		if current.Distance > distances[current.Node] {
			continue
		}

		// Update the distances of the current node's neighbors.
		for neighbor, weight := range graph.Nodes[current.Node] {
			alt := distances[current.Node] + weight
			if alt < distances[neighbor] {
				distances[neighbor] = alt
				pq = append(pq, &Node{neighbor, alt})
			}
		}
	}

	// Check if the end node is reachable.
	if distances[end] == float64(Infinity) {
		return nil, errors.New("end node is unreachable")
	}

	return distances, nil
}

type Node struct {
	Node   string
	Distance float64
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() (interface{}, bool) {
	old := *pq
	n := len(old)
	p := old[n-1]
	*pq = old[0 : n-1]
	return p, true
}

func (pq *PriorityQueue) ExtractMin() *Node {
	pq.Pop()
	return pq.Pop()
}

const (
	Infinity = 1e9
)