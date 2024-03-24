package main

import (
	"fmt"
	"math/rand"

	"github.com/dawkrish/linear_algebra/matrix"
)

func main() {
	A := [][]float64{
		{1, 3, -2, 5},
		{3, 5, 6, 7},
		{2, 4, 3, 8}}
	M, _ := matrix.RowEchelonForm(A)
	matrix.Print(M)

}

func randomMatrix(n int) [][]float64 {
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
		for j := range matrix[i] {
			matrix[i][j] = float64(rand.Intn(6) - 3)
		}
	}
	return matrix

}
