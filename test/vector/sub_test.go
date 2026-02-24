package vector_testing

import (
	"github.com/FelBenini/matrix/srcs/vector"
	"testing"
)

func TestSub(t *testing.T) {
	u := vector.From([]float32{2, 3})
	v := vector.From([]float32{5, 7})

	u.Sub(v)

	expected := []float32{-3, -4}
	for i, val := range expected {
		if !floatsAlmostEqual(u.Data[i], val) {
			t.Fatalf("Sub failed at index %d: got %v, want %v", i, u.Data[i], val)
		}
	}
}
