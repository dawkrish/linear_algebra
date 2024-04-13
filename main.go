package main

import (
	"math/rand"

	"github.com/dawkrish/linear_algebra/matrix"
)

func main() {
	A := [][]float64{{1, 2, 1}, {2, 1, 1}, {-1, 2, 1}}
	matrix.Print(A)
	println(matrix.Determinant(A))

	// M1, _ := matrix.REF(A)
	// matrix.Print(M1)
	// fmt.Println()
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
