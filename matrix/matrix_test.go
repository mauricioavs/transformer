package matrix

import (
	"testing"
)

func TestRows(t *testing.T) {
	m := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	if got := m.Rows(); got != 2 {
		t.Errorf("Rows() = %d; want 2", got)
	}

	var empty Matrix
	if got := empty.Rows(); got != 0 {
		t.Errorf("Rows() on empty = %d; want 0", got)
	}
}

func TestCols(t *testing.T) {
	m := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	if got := m.Cols(); got != 3 {
		t.Errorf("Cols() = %d; want 3", got)
	}

	var empty Matrix
	if got := empty.Cols(); got != 0 {
		t.Errorf("Cols() on empty = %d; want 0", got)
	}
}

func TestDims(t *testing.T) {
	m := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	rows, cols := m.Dims()
	if rows != 2 || cols != 3 {
		t.Errorf("Dims() = (%d, %d); want (2, 3)", rows, cols)
	}

	var empty Matrix
	rows, cols = empty.Dims()
	if rows != 0 || cols != 0 {
		t.Errorf("Dims() on empty = (%d, %d); want (0, 0)", rows, cols)
	}
}

func TestAdd(t *testing.T) {
	a := Matrix{
		{1, 2},
		{3, 4},
	}
	b := Matrix{
		{5, 6},
		{7, 8},
	}
	want := Matrix{
		{6, 8},
		{10, 12},
	}

	got := a.Add(b)
	if len(got) != len(want) {
		t.Fatalf("Add result rows = %d; want %d", len(got), len(want))
	}
	for i := range got {
		if len(got[i]) != len(want[i]) {
			t.Fatalf("Add result cols row %d = %d; want %d", i, len(got[i]), len(want[i]))
		}
		for j := range got[i] {
			if got[i][j] != want[i][j] {
				t.Errorf("Add result[%d][%d] = %f; want %f", i, j, got[i][j], want[i][j])
			}
		}
	}
}

func TestAddDimensionMismatch(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Add did not panic on dimension mismatch")
		}
	}()

	a := Matrix{
		{1, 2},
	}
	b := Matrix{
		{1, 2},
		{3, 4},
	}

	_ = a.Add(b) // must panic
}

func TestAddEmptyMatrices(t *testing.T) {
	a := Matrix{}
	b := Matrix{}

	got := a.Add(b)
	if len(got) != 0 {
		t.Errorf("Add of empty matrices result length = %d; want 0", len(got))
	}
}
