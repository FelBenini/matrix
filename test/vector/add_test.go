package vector_testing

import (
	"github.com/FelBenini/matrix/srcs/vector"
	"math"
	"testing"
)

func floatsAlmostEqual(a, b float32) bool {
	const eps = 1e-5
	return math.Abs(float64(a-b)) < eps
}

func TestAdd(t *testing.T) {
	u := vector.From([]float32{2, 3})
	v := vector.From([]float32{5, 7})

	u.Add(v)

	expected := []float32{7, 10}
	for i, val := range expected {
		if !floatsAlmostEqual(u.Data[i], val) {
			t.Fatalf("Add failed at index %d: got %v, want %v", i, u.Data[i], val)
		}
	}
}
