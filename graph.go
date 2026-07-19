package graph

import (
	"errors"
	"fmt"
	"sort"

	"github.com/pkg/errors"
)

// Graph represents a weighted graph
type Graph struct {
	Nodes   map[string]*Node
	Edges   map[string][]*Edge
	Weight  map[string]float64
	Reverse bool
}

// Node represents a node in the graph
type Node struct {
	ID        string
	Distance  float64
	Predecessor string
}

// Edge represents a directed edge between nodes
type Edge struct {
	From  string
	To    string
	Weight float64
}

// NewGraph creates a new graph
func NewGraph(reverse bool) *Graph {
	return &Graph{
		Nodes:   make(map[string]*Node),
		Edges:   make(map[string][]*Edge),
		Weight:  make(map[string]float64),
		Reverse: reverse,
	}
}

// AddNode adds a new node to the graph
func (g *Graph) AddNode(id string) *Node {
	node := &Node{ID: id}
	g.Nodes[id] = node
	return node
}

// AddEdge adds a new edge to the graph
func (g *Graph) AddEdge(from, to string, weight float64) {
	if g.Reverse {
		g.Edges[to] = append(g.Edges[to], &Edge{From: to, To: from, Weight: weight})
	} else {
		g.Edges[from] = append(g.Edges[from], &Edge{From: from, To: to, Weight: weight})
	}
	g.Weight[from+":"+to] = weight
}

// Dijkstra finds the shortest path between nodes using Dijkstra's algorithm
func (g *Graph) Dijkstra(start, end string) (float64, error) {
	startNode := g.Nodes[start]
	if startNode == nil {
		return 0, errors.New("start node not found")
	}

	endNode := g.Nodes[end]
	if endNode == nil {
		return 0, errors.New("end node not found")
	}

	for _, node := range g.Nodes {
		node.Distance = math.MaxFloat64
	}

	startNode.Distance = 0

	priorityQueue := make([]string, 0)
	priorityQueue = append(priorityQueue, start)

	for len(priorityQueue) > 0 {
		currentNode := g.Nodes[priorityQueue[0]]
		priorityQueue = priorityQueue[1:]

		for _, neighbor := range g.getNeighbors(currentNode.ID) {
			neighborNode := g.Nodes[neighbor.To]
			alt := currentNode.Distance + g.Weight[currentNode.ID+":"+neighbor.To]
			if neighborNode.Distance > alt {
				neighborNode.Distance = alt
				neighborNode.Predecessor = currentNode.ID
				priorityQueue = append(priorityQueue, neighbor.To)
			}
		}
		sort.Strings(priorityQueue)
	}

	return endNode.Distance, nil
}

func (g *Graph) getNeighbors(id string) []*Edge {
	if g.Reverse {
		for _, edge := range g.Edges[id] {
			return []*Edge{edge}
		}
		return nil
	}
	for _, edge := range g.Edges[id] {
		return []*Edge{edge}
	}
	return nil
}

func (g *Graph) String() string {
	var sb strings.Builder
	for id, node := range g.Nodes {
		sb.WriteString(fmt.Sprintf("%s: distance=%f predecessor=%s\n", id, node.Distance, node.Predecessor))
	}
	return sb.String()
}