package vector_testing

import (
	"github.com/FelBenini/matrix/srcs/vector"
	"testing"
)

func TestScl(t *testing.T) {
	u := vector.From([]float32{2, 3})
	u.Scl(2)

	expected := []float32{4, 6}
	for i, val := range expected {
		if !floatsAlmostEqual(u.Data[i], val) {
			t.Fatalf("Scl failed at index %d: got %v, want %v", i, u.Data[i], val)
		}
	}
}
