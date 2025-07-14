package matrix

import (
	"fmt"
	"math"
	"reflect"
)

var ErrorTolerance = 1e-9 // Tolerance for comparisons

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

func (m Matrix) T() Matrix {
	rows := m.Rows()
	if rows == 0 {
		return Matrix{}
	}
	cols := m.Cols()

	result := make(Matrix, cols)
	for i := range cols {
		result[i] = make([]float64, rows)
		for j := range rows {
			result[i][j] = m[j][i]
		}
	}
	return result
}

func (m Matrix) Scale(factor float64) Matrix {
	rows := m.Rows()
	cols := m.Cols()
	result := make(Matrix, rows)
	for i := range rows {
		result[i] = make([]float64, cols)
		for j := range cols {
			result[i][j] = m[i][j] * factor
		}
	}
	return result
}

func (a Matrix) Add(b Matrix) Matrix {
	rowsA, colsA := a.Dims()
	rowsB, colsB := b.Dims()

	if rowsA != rowsB || colsA != colsB {
		panic(
			fmt.Sprintf(
				"Matrices have different dimensions: A is %dx%d, B is %dx%d",
				rowsA, colsA, rowsB, colsB,
			),
		)
	}

	result := make(Matrix, rowsA) // [[], [], ..., []] with rowsA rows
	for i := range rowsA {
		result[i] = make([]float64, colsA)
		for j := range colsA {
			result[i][j] = a[i][j] + b[i][j]
		}
	}
	return result
}

func (a Matrix) Mul(b Matrix) Matrix {
	rowsA, colsA := a.Dims()
	rowsB, colsB := b.Dims()

	if colsA != rowsB {
		panic(fmt.Sprintf(
			"Cannot multiply: incompatible dimensions. A is %dx%d, B is %dx%d. cols(A) must be equal to rows(B).",
			rowsA, colsA, rowsB, colsB,
		))
	}

	result := make(Matrix, rowsA)
	for i := range rowsA {
		result[i] = make([]float64, colsB)
		for j := range colsB {
			sum := 0.0
			for k := range colsA {
				sum += a[i][k] * b[k][j]
			}
			result[i][j] = sum
		}
	}
	return result
}

func (a Matrix) Eq(b Matrix) bool {
	return reflect.DeepEqual(a, b)
}

func (a Matrix) Neq(b Matrix) bool {
	return !a.Eq(b)
}

func (a Matrix) EqApprox(b Matrix) bool {
	if a.Rows() != b.Rows() || a.Cols() != b.Cols() {
		return false
	}
	for i := 0; i < a.Rows(); i++ {
		for j := 0; j < a.Cols(); j++ {
			if math.Abs(a[i][j]-b[i][j]) > ErrorTolerance {
				return false
			}
		}
	}
	return true
}

func (a Matrix) NeqApprox(b Matrix) bool {
	return !a.EqApprox(b)
}
