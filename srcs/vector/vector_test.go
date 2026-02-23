package vector

import (
	"math"
	"testing"
)

func floatsAlmostEqual(a, b float32) bool {
	const eps = 1e-5
	return math.Abs(float64(a-b)) < eps
}

func TestAdd(t *testing.T) {
	u := From([]float32{2, 3})
	v := From([]float32{5, 7})

	u.Add(v)

	expected := []float32{7, 10}
	for i, val := range expected {
		if !floatsAlmostEqual(u.Data[i], val) {
			t.Fatalf("Add failed at index %d: got %v, want %v", i, u.Data[i], val)
		}
	}
}
func TestSub(t *testing.T) {
	u := From([]float32{2, 3})
	v := From([]float32{5, 7})

	u.Sub(v)

	expected := []float32{-3, -4}
	for i, val := range expected {
		if !floatsAlmostEqual(u.Data[i], val) {
			t.Fatalf("Sub failed at index %d: got %v, want %v", i, u.Data[i], val)
		}
	}
}

func TestScl(t *testing.T) {
	u := From([]float32{2, 3})
	u.Scl(2)

	expected := []float32{4, 6}
	for i, val := range expected {
		if !floatsAlmostEqual(u.Data[i], val) {
			t.Fatalf("Scl failed at index %d: got %v, want %v", i, u.Data[i], val)
		}
	}
}
