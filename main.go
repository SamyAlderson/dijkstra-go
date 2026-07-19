// main.go: Main entry point for dijkstra-go

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

import (
	"dijkstra"
	"graph"
)

func main() {
	// Create a sample weighted graph
	g := graph.NewWeightedGraph(5)
	g.AddEdge(0, 1, 3)
	g.AddEdge(0, 2, 8)
	g.AddEdge(1, 2, 2)
	g.AddEdge(1, 3, 5)
	g.AddEdge(2, 3, 10)
	g.AddEdge(2, 4, 4)
	g.AddEdge(3, 4, 6)

	// Run Dijkstra's algorithm on the graph
	distances, predecessors := dijkstra.Run(context.Background(), g, 0)

	// Print the shortest distances and predecessors
	fmt.Println("Shortest Distances:")
	for node, distance := range distances {
		fmt.Printf("Node %d: %d\n", node, distance)
	}
	fmt.Println("Predecessors:")
	for node, predecessor := range predecessors {
		fmt.Printf("Node %d: %d\n", node, predecessor)
	}

	// Example usage: find the shortest path from node 0 to node 4
	path := dijkstra.ReconstructPath(0, predecessors)
	fmt.Println("Shortest Path from Node 0 to Node 4:")
	fmt.Println(path)
}

func exampleUsage() {
	fmt.Println("Running Dijkstra's algorithm on a sample graph...")
}

func printError(err error) {
	log.Fatal(errors.Wrap(err, "Error running Dijkstra's algorithm"))
}