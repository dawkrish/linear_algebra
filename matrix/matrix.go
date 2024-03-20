package matrix

import (
	"errors"
	"fmt"
)

// Returns a matrix with m-rows & n-cols
// All the elements are 0
func NewMatrix(m, n int) [][]float64 {
	var matrix [][]float64
	for i := 0; i < m; i++ {
		var row []float64
		for j := 0; j < n; j++ {
			row = append(row, 0)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

// Returns a matrix with n-rows & n-cols.
// Diagonal Elements are 1 , rest are 0
func NewIdentityMatrix(n int) [][]float64 {
	matrix := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				matrix[i][j] = 1
			} else {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}

func GetRowAt(index int, A [][]float64) ([]float64, error) {
	if index < 0 || index >= NumberOfRows(A) {
		return []float64{}, errors.New("index out of range")
	}
	return A[index], nil
}

func GetColumnAt(index int, A [][]float64) ([]float64, error) {
	if index < 0 || index >= NumberOfCols(A) {
		return []float64{}, errors.New("index out of range")
	}
	col := []float64{}
	for i := 0; i < NumberOfRows(A); i++ {
		col = append(col, A[i][index])
	}
	return col, nil
}

// Returns matrix which belongs to n x m
func Transpose(A [][]float64) [][]float64 {
	m := NumberOfRows(A)
	n := NumberOfCols(A)
	matrix := NewMatrix(n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i][j] = A[j][i]
		}
	}
	return matrix
}

func Determinant(A [][]float64) (int, error) {
	if NumberOfRows(A) != NumberOfCols(A) {
		return 0, errors.New("number of rows != number of cols")
	}
	det := 0
	for i := 0; i < NumberOfRows(A); i++ {

	}
	return det, nil
}

// Returns number of rows
func NumberOfRows(matrix [][]float64) int {
	return len(matrix)
}

// Returns number of cols
func NumberOfCols(matrix [][]float64) int {
	if len(matrix) == 0 {
		return 0
	}
	return len(matrix[0])
}

// Returns error if the matrixes are of not same dimensions
// Returns C = A + B, where C[i][j] = A[i][j] + B[i][j]
func Addition(A, B [][]float64) ([][]float64, error) {
	if NumberOfRows(A) != NumberOfRows(B) {
		return [][]float64{}, errors.New("number of rows are different")
	}
	if NumberOfCols(A) != NumberOfCols(B) {
		return [][]float64{}, errors.New("number of col are different")
	}

	m := NumberOfRows(A)
	n := NumberOfCols(A)
	C := NewMatrix(m, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}

	return C, nil
}

// Returns error if the matrixes don't follow dimensions for multiplaction
// Returns C = A*B, where C[i][j] = A[i][k]*B[k][j]
func Multiply(A, B [][]float64) ([][]float64, error) {
	if NumberOfCols(A) != NumberOfRows(B) {
		return [][]float64{}, errors.New("number of cols of A != number of rows of B")
	}

	m := NumberOfRows(A)
	n := NumberOfCols(A)
	p := NumberOfCols(B)

	C := NewMatrix(m, p)

	for i := 0; i < m; i++ {
		for j := 0; j < p; j++ {
			for k := 0; k < n; k++ {
				C[i][j] += (A[i][k] * B[k][j])
			}
		}
	}
	return C, nil
}

func Print(A [][]float64) {
	m := NumberOfRows(A)
	n := NumberOfCols(A)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				fmt.Print("[")
			}
			fmt.Printf("%6.2f,", A[i][j])
			if j == n-1 {
				fmt.Print("]")
			}
		}
		fmt.Println()
	}
}

func Copy(A [][]float64) [][]float64 {
	m := NumberOfRows(A)
	n := NumberOfCols(A)
	matrix := NewMatrix(m, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = A[i][j]
		}
	}
	return matrix
}
