package matrix

import "fmt"

type Matrix [][]float64

func (m Matrix) Rows() int {
	return len(m)
}

func (m Matrix) Cols() int {
	if m.Rows() == 0 {
		return 0
	}
	return len(m[0])
}

func (m Matrix) Dims() (int, int) {
	return m.Rows(), m.Cols()
}

func (a Matrix) Add(b Matrix) Matrix {
	rowsA, colsA := a.Dims()
	rowsB, colsB := b.Dims()

	if rowsA != rowsB || colsA != colsB {
		panic(
			fmt.Sprintf(
				"Matrices have different dimensions: a is %dx%d, b is %dx%d",
				rowsA, colsA, rowsB, colsB,
			),
		)
	}

	result := make(Matrix, rowsA) // [[], [], ..., []] con rowsA filas
	for i, row := range a {
		result[i] = make([]float64, len(row))
		for j, val := range row {
			result[i][j] = val + b[i][j]
		}
	}
	return result
}
