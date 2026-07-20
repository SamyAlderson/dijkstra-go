# dijkstra-go
A simple Go implementation of Dijkstra's algorithm for shortest paths in weighted graphs.

## What it does
Dijkstra's algorithm is used to find the shortest path between nodes in a weighted graph. This implementation uses a priority queue to efficiently explore the graph and calculate the shortest distances.

## Installation
To install the package, run:
```bash
go get github.com/samyalderson/dijkstra-go
```
## Usage
To use the package, import and instantiate the `Dijkstra` struct, then call the `Run` method on it:
```go
import (
	"github.com/samyalderson/dijkstra-go/graph"
	"github.com/samyalderson/dijkstra-go/dijkstra"
)

func main() {
	g := graph.NewWeightedGraph()
	// populate the graph with nodes and edges...

	d := dijkstra.NewDijkstra(g)
	d.Run("start", "end")
	// d.Distances now contains the shortest distances from the start node to all other nodes
}
```
## Build from source
To build the package from source, navigate to the project root and run:
```bash
go build .
```
This will generate an executable named `dijkstra-go` in the current directory.

## Running tests
To run the package's unit tests, navigate to the project root and run:
```bash
go test .
```
This will run the `graph_test.go` and `dijkstra_test.go` test files.

## Project structure
- `graph.go`: implementation of a weighted graph data structure
- `dijkstra.go`: implementation of Dijkstra's algorithm
- `graph_test.go`: unit tests for the weighted graph
- `dijkstra_test.go`: unit tests for Dijkstra's algorithm
- `README.md`: this file
- `LICENSE`: MIT license text

## License
MIT License

Copyright (c) 2026 SamyAlderson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.