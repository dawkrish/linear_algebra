package matrix

import "errors"

// Ri <-> Rj
func RowSwitch(i, j int, A [][]float64) ([][]float64, error) {
	if i < 0 || i > NumberOfRows(A)-1 {
		return [][]float64{}, errors.New("i is in invalid range")
	}
	if j < 0 || j > NumberOfRows(A)-1 {
		return [][]float64{}, errors.New("j is in invalid range")
	}

	matrix := Copy(A)
	matrix[i], matrix[j] = matrix[j], matrix[i]
	return matrix, nil
}

func RowMultiplication(scalar float64, i int, A [][]float64) ([][]float64, error) {
	if scalar == 0 {
		return [][]float64{}, errors.New("k must not be 0")
	}
	matrix := Copy(A)
	for j := 0; j < NumberOfCols(A); j++ {
		matrix[i][j] *= scalar
	}
	return matrix, nil
}

// Ri = Ri + Rj * scalar
func RowAddition(scalar float64, i, j int, A [][]float64) ([][]float64, error) {
	if i == j {
		return [][]float64{}, errors.New("i must not be equal to j")
	}
	matrix := Copy(A)
	for l := 0; l < NumberOfCols(A); l++ {
		matrix[i][l] += matrix[j][l] * scalar
	}
	return matrix, nil
}
