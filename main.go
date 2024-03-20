package main

import (
	"fmt"

	"github.com/dawkrish/linear_algebra/matrix"
)

func main() {
	A := [][]float64{{1, 2, 3}, {1, 4, 5}, {1, 7, 8}}
	matrix.Print(A)
	fmt.Println()
	B, err := matrix.RowSwitch(0, 2, A)
	if err != nil {
		fmt.Println("error : ", err)
	}
	matrix.Print(B)
	fmt.Println()
	C, err := matrix.RowMultiplication(1.0/3, 1, A)
	if err != nil {
		fmt.Println("error : ", err)
	}
	matrix.Print(C)
	fmt.Println()
	D, err := matrix.RowAddition(2, 0, 2, A)
	matrix.Print(D)

	matrix.RowEchelonForm(A)
	// fmt.Println(matrix.GetColumnAt(1, A))
	// fmt.Println(matrix.GetRowAt(1, A))

}
