package matrix

import (
	"errors"
	"fmt"
	"math"
)

// Returns a matrix with m-rows & n-cols
// All the elements are 0
func NewMatrix(m, n int) [][]int {
	var matrix [][]int
	for i := 0; i < m; i++ {
		var row []int
		for j := 0; j < n; j++ {
			row = append(row, 0)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

// Returns a matrix with n-rows & n-cols.
// Diagonal Elements are 1 , rest are 0
func NewIdentityMatrix(n int) [][]int {
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

// Returns matrix which belongs to n x m
func Transpose(A [][]int) [][]int {
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

// Returns number of rows
func NumberOfRows(matrix [][]int) int {
	return len(matrix)
}

// Returns number of cols
func NumberOfCols(matrix [][]int) int {
	return len(matrix[0])
}

// Returns error if the matrixes are of not same dimensions
// Returns C = A + B, where C[i][j] = A[i][j] + B[i][j]
func Addition(A, B [][]int) ([][]int, error) {
	if NumberOfRows(A) != NumberOfRows(B) {
		return [][]int{}, errors.New("number of rows are different")
	}
	if NumberOfCols(A) != NumberOfCols(B) {
		return [][]int{}, errors.New("number of col are different")
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
func Multiply(A, B [][]int) ([][]int, error) {
	if NumberOfCols(A) != NumberOfRows(B) {
		return [][]int{}, errors.New("number of cols of A != number of rows of B")
	}

	m := NumberOfRows(A)
	n := NumberOfCols(A)
	p := NumberOfCols(B)
	C := NewMatrix(m, p)

	for i := 0; i < m; i++ {
		for j := 0; j < p; j++ {
			var sum int
			for k := 0; k < n; k++ {
				sum += (A[i][k] * B[k][j])
			}
			C[i][j] = sum
		}
	}
	return C, nil
}

func Print(A [][]int) {
	m := NumberOfRows(A)
	n := NumberOfCols(A)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				fmt.Print("[")
			}
			fmt.Printf("%6d,", A[i][j])
			if j == n-1 {
				fmt.Print("]")
			}
		}
		fmt.Println()
	}
}

func largestNumber(A [][]int) int {
	var max = math.MinInt
	for i := 0; i < NumberOfRows(A); i++ {
		for j := 0; j < NumberOfCols(A); j++ {
			if A[i][j] > max {
				max = A[i][j]
			}
		}
	}
	return max
}
