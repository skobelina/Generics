package main

import "fmt"

// Task: create Counter with method Next(). Counter expected to support all kinds of int and its inherit types
func main() {
	count := Counter[int64]{0}
	count.Next()
	fmt.Println("count:", count.value)

	countUint := Counter[uint64]{0}
	countUint.Next()
	fmt.Println("countUint:", countUint.value)

	countFloat := Counter[float64]{0}
	countFloat.Next()
	fmt.Println("countFloat:", countFloat.value)
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Counter[T Number] struct {
	value T
}

func (c *Counter[T]) Next() {
	c.value++
}
