package main

import (
	// "fmt"

	"github.com/dawkrish/linear_algebra/matrix"
)

func main() {
	A := [][]int{{1, 2, 3}, {1, 4, 5}}
	// B := [][]int{{1, 0, 1}, {4, 5, 2}, {7, 0, 3}}

	matrix.Print(A)
	// fmt.Print(matrix.Transpose(A))
}
