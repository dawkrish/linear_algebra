package matrix

import "fmt"

func RowEchelonForm(A [][]float64) [][]float64 {
	matrix := Copy(A)

	colIndex := LeftMostColumnWithNonZeroEntry(A)
	column, _ := GetColumnAt(colIndex, A)
	fmt.Println("Column > ", column)
	return matrix
}

func LeftMostColumnWithNonZeroEntry(A [][]float64) int {
	for i := 0; i < NumberOfCols(A); i++ {
		for j := 0; j < NumberOfRows(A); j++ {
			if A[j][i] != 0 {
				return i
			}
		}
	}
	return -1
}
