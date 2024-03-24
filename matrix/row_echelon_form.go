package matrix

// import "fmt"

func LeftMostColumnWithNonZeroEntry(A [][]float64, currentRow int) (int, int) {
	for i := 0; i < NumberOfCols(A); i++ {
		for j := currentRow; j < NumberOfRows(A); j++ {
			if A[j][i] != 0 {
				return j, i
			}
		}
	}
	return -1, -1
}

func RowEchelonForm(A [][]float64) ([][]float64, int) {
	matrix := Copy(A)
	var determinantFactor = 1
	rows := NumberOfRows(A)

	for i := 0; i < rows; i++ {
		r_idx, c_idx := LeftMostColumnWithNonZeroEntry(matrix, i)
		// r_idx = represents the row_index of the entry which is non-zero; it needs to be same as "i"; or else swap it.
		// c_idx = represents the col_index of the entry which is non-zero; on that

		if r_idx == -1 || c_idx == -1 {
			break
		}
		if r_idx != i {
			matrix, _ = RowSwitch(r_idx, i, matrix)
			determinantFactor *= -1
		}
		column, _ := GetColumnAt(c_idx, matrix)

		for j := r_idx + 1; j < rows; j++ {
			scalar := -1 * (float64(column[j]) / float64(column[i]))
			matrix, _ = RowAddition(scalar, j, i, matrix)
		}

	}
	return matrix, determinantFactor
}

// fmt.Printf("r_idx : %v\tc_idx : %v\n", r_idx, c_idx)
// fmt.Println("WE ARE HERE")
// fmt.Printf("i:%v\tj:%v\tcolumn:%v\n", i, j, column)
// fmt.Printf("Numerator:%v\tDenimnator:%v\tScalar:%v\n", float64(column[j]), float64(column[i]), scalar)
