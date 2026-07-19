# Dijkstra's Algorithm in Go

Implementation of Dijkstra's algorithm in Go for finding the shortest path in a weighted graph.

## What's this?

This project provides a Go implementation of Dijkstra's algorithm for finding the shortest path in a weighted graph. It includes a weighted graph data structure and unit tests for both the algorithm and the graph.

## Install

To install the project, run:

```bash
go get -u github.com/samyalderson/dijkstra-go
```

## Usage

To use the project, run the `main.go` file:

```bash
go run main.go
```

This will prompt you to enter the graph edges and the starting node. The program will then output the shortest path from the starting node to all other nodes in the graph.

## Build from Source

To build the project from source, run:

```bash
make build
```

## Project Structure

The project consists of the following files:

* `main.go`: Main entry point
* `dijkstra.go`: Dijkstra's algorithm implementation
* `graph.go`: Weighted graph data structure
* `test_dijkstra_test.go`: Unit tests for Dijkstra's algorithm
* `test_graph_test.go`: Unit tests for weighted graph
* `go.mod`: Go module file
* `go.sum`: Go checksum file
* `Makefile`: Build script for Go project

## License

This project is licensed under the MIT license. See the `LICENSE` file for details.

## Contributing

Pull requests are welcome. If you'd like to contribute, please fork the repository and submit a pull request.