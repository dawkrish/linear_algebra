package main

import (
	"fmt"
	"math/rand"

	"github.com/dawkrish/linear_algebra/matrix"
)

func main() {
	A := [][]float64{{1, 1, 1, 10}, {1, 1, 2, 15}, {1, 1, 4, 20}}
	matrix.Print(A)
	println()
	// M1, _ := matrix.REF(A)
	// matrix.Print(M1)
	// fmt.Println()
	M2 := matrix.RREF(A)
	matrix.Print(M2)
	fmt.Println()
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
