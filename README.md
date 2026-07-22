# Dijkstra's Algorithm in Go
Find the shortest path in a weighted graph with ease

## What it does

Dijkstra's algorithm is a well-known solution for finding the shortest path in a weighted graph. This Go implementation provides a simple and efficient way to compute the shortest distances and paths between nodes in a graph.

## Install

To install dijkstra-go, run:
```bash
go get github.com/samyaldrson/dijkstra-go
```
## Usage

To use dijkstra-go, import the package and run the algorithm on your graph:
```go
import (
	"github.com/samyaldrson/dijkstra-go"
)

func main() {
	graph := dijkstra.Graph{
		// define your graph here
	}
	distances, paths := dijkstra.Dijkstra(graph)
	fmt.Println(distances)
	fmt.Println(paths)
}
```
## Build from Source

To build from source, run:
```bash
go build -o dijkstra-go main.go
```
## Running Tests

To run the test suite, run:
```bash
go test -v
```
## Project Structure

* `dijkstra.go`: Dijkstra's algorithm implementation
* `graph.go`: Graph data structure
* `test_dijkstra.go`: Test suite for dijkstra's algorithm
* `test_graph.go`: Test suite for graph data structure
* `main.go`: Example usage

## License

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