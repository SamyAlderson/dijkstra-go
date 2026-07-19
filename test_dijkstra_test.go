package dijkstra

import (
	"testing"

	"github.com/pkg/errors"
)

func TestDijkstra(t *testing.T) {
	g := NewGraph()
	g.AddEdge("A", "B", 3)
	g.AddEdge("A", "C", 2)
	g.AddEdge("B", "C", 1)
	g.AddEdge("C", "D", 4)

	dist, prev := Dijkstra(g, "A")

	if dist["D"] != 6 {
		t.Errorf("expected shortest distance to node D is 6, got %d", dist["D"])
	}

	if prev["D"] != "C" {
		t.Errorf("expected shortest path to node D is C, got %s", prev["D"])
	}
}

func TestDijkstraEmptyGraph(t *testing.T) {
	g := NewGraph()
	_, _ = Dijkstra(g, "A")
}

func TestDijkstraNegativeWeight(t *testing.T) {
	e := errors.New("negative weight is not allowed")
	g := NewGraph()
	g.AddEdge("A", "B", -1)
	if _, err := Dijkstra(g, "A"); err == nil {
		t.Errorf("expected error on negative weight, got nil")
	}
}

func TestDijkstraNonExistingNode(t *testing.T) {
	g := NewGraph()
	g.AddEdge("A", "B", 3)
	if _, err := Dijkstra(g, "C"); err == nil {
		t.Errorf("expected error on non-existing node, got nil")
	}
}