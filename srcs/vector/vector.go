package vector

import (
	"fmt"
	"github.com/FelBenini/matrix/srcs/number"
)

type Number = number.Number

type Vector[T Number] struct {
	Data []T
}

func From[T Number](data []T) Vector[T] {
	return Vector[T]{
		Data: append([]T(nil), data...),
	}
}

func (v Vector[T]) Print() {
	count := len(v.Data)

	for i := range count {
		fmt.Print(v.Data[i])
		if i < len(v.Data) {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
