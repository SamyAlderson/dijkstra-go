go.mod
module dijkstra-go

go 1.19

require (
	github.com/pkg/errors v0.0.0-20200624221458-97e15bb1c46f
)

require (
	github.com/stretchr/testify v1.7.0
) 

replace (
	github.com/pkg/errors => ./errors 
)