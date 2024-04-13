package matrix

func REF(A [][]float64) ([][]float64, int) {
	matrix := Copy(A)
	cols := NumberOfCols(A)
	currentPivotRowIndex := 0
	determinantFactor := 1

	for colIndex := 0; colIndex < cols; colIndex++ {
		foundPivotRowIndex := GetPivotIndex(colIndex, currentPivotRowIndex, matrix)
		if foundPivotRowIndex == -1 {
			// skip , as there is no pivot in this column
		} else {
			// switch the rows with the "foundPivotRowIndex" with the "currentPivotRowIndex"
			if foundPivotRowIndex != currentPivotRowIndex {
				matrix, _ = RowSwitch(foundPivotRowIndex, currentPivotRowIndex, matrix)
				determinantFactor *= -1
			}

			// scale the below rows ...
			for i := currentPivotRowIndex + 1; i < NumberOfRows(matrix); i++ {
				scalar := -1 * float64(matrix[i][colIndex]) / float64(matrix[currentPivotRowIndex][colIndex])
				matrix, _ = RowAddition(scalar, i, currentPivotRowIndex, matrix)

			}
			currentPivotRowIndex += 1
		}
	}
	return matrix, determinantFactor
}

func RREF(A [][]float64) [][]float64 {
	matrix := Copy(A)
	cols := NumberOfCols(A)
	currentPivotRowIndex := 0

	for colIndex := 0; colIndex < cols; colIndex++ {
		foundPivotRowIndex := GetPivotIndex(colIndex, currentPivotRowIndex, matrix)
		if foundPivotRowIndex == -1 {
			// skip , as there is no pivot in this column
		} else {
			// switch the rows with the "foundPivotRowIndex" with the "currentPivotRowIndex"
			if foundPivotRowIndex != currentPivotRowIndex {
				matrix, _ = RowSwitch(foundPivotRowIndex, currentPivotRowIndex, matrix)
			}

			//scale the pivot row !
			inverseScalar := 1 / float64(matrix[currentPivotRowIndex][colIndex])
			matrix, _ = RowMultiplication(inverseScalar, currentPivotRowIndex, matrix)

			// scale the below rows ...
			for i := 0; i < NumberOfRows(matrix); i++ {
				if i == currentPivotRowIndex {
					continue
				}
				scalar := -1 * float64(matrix[i][colIndex]) / float64(matrix[currentPivotRowIndex][colIndex])
				matrix, _ = RowAddition(scalar, i, currentPivotRowIndex, matrix)

			}
			currentPivotRowIndex += 1
		}
	}
	return matrix
}

func GetPivotIndex(colIndex, rowIndex int, A [][]float64) int {
	// we need to find the pivot from the "currentRowIndex"
	for i := rowIndex; i < NumberOfRows(A); i++ {
		if A[i][colIndex] != 0 {
			return i
		}
	}
	return -1
}

// fmt.Printf("r_idx : %v\tc_idx : %v\n", r_idx, c_idx)
// fmt.Println("WE ARE HERE")
// fmt.Printf("i:%v\tj:%v\tcolumn:%v\n", i, j, column)
// fmt.Printf("Numerator:%v\tDenimnator:%v\tScalar:%v\n", float64(column[j]), float64(column[i]), scalar)
