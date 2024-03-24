package matrix

import "fmt"

func Print(A [][]float64) {
	m := NumberOfRows(A)
	n := NumberOfCols(A)

	maxWidth := make([]int, n)
	for j := 0; j < n; j++ {
		maxWidth[j] = 0
		for i := 0; i < m; i++ {
			width := len(fmt.Sprintf("%v", A[i][j]))
			if width > maxWidth[j] {
				maxWidth[j] = width
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				fmt.Print("[")
			}
			fmt.Printf("%*v", maxWidth[j]+1, A[i][j]) // Adjust width based on maxWidth
			if j == n-1 {
				fmt.Print("]")
			} else {
				fmt.Print(",")
			}
		}
		fmt.Println()
	}
}
