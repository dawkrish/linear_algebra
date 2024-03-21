package main

import (
	"fmt"

	"github.com/dawkrish/linear_algebra/matrix"
)

func main() {

	B := [][]float64{
		{1, 2, -1, 2, 47},
		{0, 4, 9, 2, -93},
		{3, 7, 2, 2, 8},
		{-5, -2, 2, 7, 9},
		{1, 82, 7, 82, 0}}
	matrix.Print(B)
	fmt.Println()
	C, _ := matrix.RowEchelonForm(B)
	matrix.Print(C)
	det, _ := matrix.Determinant(B)
	fmt.Println()
	fmt.Println("Determinant : ", det)

}

// [][]float64{{0, 2, -1, 3, 5}, {0, 4, 9, 2, 7}, {0, 7, 2, 5, 4}, {0, -3, 6, 5, 14}}
