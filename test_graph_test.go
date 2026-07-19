package graph

import (
	"errors"
	"testing"

	"github.com/pkg/errors"
)

func TestGraph(t *testing.T) {
	graph := NewGraph()

	if err := graph.AddEdge("A", "B", 1); err != nil {
		t.Fatal(err)
	}

	if err := graph.AddEdge("B", "C", 2); err != nil {
		t.Fatal(err)
	}

	if err := graph.AddEdge("C", "A", 3); err != nil {
		t.Fatal(err)
	}

	if err := graph.AddEdge("A", "C", 4); err != nil {
		t.Fatal(err)
	}

	if err := graph.AddVertex("D"); err != nil {
		t.Fatal(err)
	}

	// This should fail because there's already a vertex with this name
	if err := graph.AddVertex("A"); err != nil {
		if !errors.Is(err, ErrVertexExists) {
			t.Errorf("expected ErrVertexExists, got %v", err)
		}
	}

	paths, err := graph.Paths()
	if err != nil {
		t.Fatal(err)
	}

	if len(paths) != 3 {
		t.Errorf("expected 3 paths, got %v", len(paths))
	}

	// We expect all vertices to have a path to themselves
	for _, path := range paths {
		if path.Source != path.Destination {
			t.Errorf("expected path from vertex to itself, got %+v", path)
		}
	}
}

func TestGraphEdges(t *testing.T) {
	graph := NewGraph()

	if err := graph.AddEdge("A", "B", 1); err != nil {
		t.Fatal(err)
	}

	if err := graph.AddEdge("B", "C", 2); err != nil {
		t.Fatal(err)
	}

	if err := graph.AddEdge("C", "A", 3); err != nil {
		t.Fatal(err)
	}

	if err := graph.AddEdge("A", "C", 4); err != nil {
		t.Fatal(err)
	}

	edges, err := graph.Edges()
	if err != nil {
		t.Fatal(err)
	}

	if len(edges) != 4 {
		t.Errorf("expected 4 edges, got %v", len(edges))
	}

	for _, edge := range edges {
		if edge.Source == edge.Destination {
			t.Errorf("expected non-self edges, got %+v", edge)
		}
	}
}

func TestGraphVertices(t *testing.T) {
	graph := NewGraph()

	if err := graph.AddVertex("A"); err != nil {
		t.Fatal(err)
	}

	if err := graph.AddVertex("B"); err != nil {
		t.Fatal(err)
	}

	if err := graph.AddVertex("C"); err != nil {
		t.Fatal(err)
	}

	if err := graph.AddVertex("A"); err != nil {
		if !errors.Is(err, ErrVertexExists) {
			t.Errorf("expected ErrVertexExists, got %v", err)
		}
	}

	vertices, err := graph.Vertices()
	if err != nil {
		t.Fatal(err)
	}

	if len(vertices) != 3 {
		t.Errorf("expected 3 vertices, got %v", len(vertices))
	}
}