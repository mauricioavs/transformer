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

func TestScale(t *testing.T) {
	a := Matrix{
		{1, -2},
		{3, 0},
	}
	factor := 2.0
	expected := Matrix{
		{2, -4},
		{6, 0},
	}

	result := a.Scale(factor)
	if result.Neq(expected) {
		t.Errorf("Scale() failed. Expected %v, got %v", expected, result)
	}
}

func TestTranspose(t *testing.T) {
	a := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	expected := Matrix{
		{1, 4},
		{2, 5},
		{3, 6},
	}

	result := a.T()
	if result.Neq(expected) {
		t.Errorf("Transpose() failed. Expected %v, got %v", expected, result)
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

func TestAdd_DimensionMismatch(t *testing.T) {
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

	a.Add(b) // must panic
}

func TestAdd_EmptyMatrices(t *testing.T) {
	a := Matrix{}
	b := Matrix{}

	got := a.Add(b)
	if len(got) != 0 {
		t.Errorf("Add of empty matrices result length = %d; want 0", len(got))
	}
}

func TestMul(t *testing.T) {
	a := Matrix{
		{1, 2},
		{3, 4},
	}
	b := Matrix{
		{2, 0},
		{1, 2},
	}
	expected := Matrix{
		{4, 4},
		{10, 8},
	}

	result := a.Mul(b)
	if result.Neq(expected) {
		t.Errorf("Mul() failed. Expected %v, got %v", expected, result)
	}
}

func TestMul_DimensionMismatch(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on multiplication dimension mismatch")
		}
	}()

	a := Matrix{
		{1, 2},
	}
	b := Matrix{
		{1, 2},
	}
	_ = a.Mul(b)
}

func TestEq(t *testing.T) {
	a := Matrix{
		{1, 2},
		{3, 4},
	}
	b := Matrix{
		{1, 2},
		{3, 4},
	}
	c := Matrix{
		{1, 2},
		{4, 3},
	}

	if !a.Eq(b) {
		t.Errorf("Eq failed: matrices should be equal")
	}

	if a.Eq(c) {
		t.Errorf("Eq failed: matrices should NOT be equal")
	}
}

func TestNeq(t *testing.T) {
	a := Matrix{
		{1, 2},
	}
	b := Matrix{
		{1, 2},
	}
	c := Matrix{
		{2, 1},
	}

	if a.Neq(b) {
		t.Errorf("Neq failed: matrices should be equal (Neq=false)")
	}

	if !a.Neq(c) {
		t.Errorf("Neq failed: matrices should NOT be equal (Neq=true)")
	}
}

func TestEqApprox(t *testing.T) {
	// Matrices with difference less than tolerance (should be equal)
	a := Matrix{
		{1.0, 2.0 + 0.9e-8}, // difference 0.9e-8 < 1e-8 tolerance
		{3.0, 4.0},
	}
	b := Matrix{
		{1.0, 2.0},
		{3.0, 4.0},
	}
	// Matrix with difference just above tolerance (should fail equality)
	c := Matrix{
		{1.0, 2.0 + 1.1e-8}, // difference 1.1e-8 > 1e-8 tolerance
		{3.0, 4.0},
	}

	oldTol := ErrorTolerance
	ErrorTolerance = 1e-8
	defer func() { ErrorTolerance = oldTol }()

	if !a.EqApprox(b) {
		t.Errorf("EqApprox failed: matrices should be approximately equal")
	}

	if c.EqApprox(b) {
		t.Errorf("EqApprox failed: matrices should NOT be approximately equal")
	}

	// Test with different dimensions (should be considered different)
	d := Matrix{{1, 2}}
	e := Matrix{{1, 2}, {0, 0}}
	if d.EqApprox(e) {
		t.Errorf("EqApprox failed: matrices with different dimensions should NOT be equal")
	}
}

func TestNeqApprox(t *testing.T) {
	a := Matrix{
		{1.0, 2.0},
	}
	// Difference less than tolerance -> should be considered equal -> NeqApprox false
	b := Matrix{
		{1.0, 2.0 + 0.9e-6}, // difference 0.9e-6 < 1e-6
	}
	// Difference greater than tolerance -> should be considered different -> NeqApprox true
	c := Matrix{
		{1.0, 2.0 + 1.1e-6}, // difference 1.1e-6 > 1e-6
	}

	oldTol := ErrorTolerance
	ErrorTolerance = 1e-6
	defer func() { ErrorTolerance = oldTol }()

	if a.NeqApprox(b) {
		t.Errorf("NeqApprox failed: matrices should be approximately equal (NeqApprox=false)")
	}

	if !a.NeqApprox(c) {
		t.Errorf("NeqApprox failed: matrices should NOT be approximately equal (NeqApprox=true)")
	}
}
