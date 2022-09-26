package main

import "fmt"

type List[T any] struct {
	elems []T
}

// Task: create chaining methods Add(T), Update(i, T), Pop(), Delete(i)

func (l *List[T]) Add(t T) *List[T] {
	l.elems = append(l.elems, t)
	return l
}

func (l *List[T]) Update(i int, t T) *List[T] {
	l.elems[i] = t
	return l
}

func (l *List[T]) Pop() T {
	element := l.elems[len(l.elems)-1]
	l.elems = append(l.elems[0:len(l.elems)-1], l.elems[len(l.elems):]...)
	return element
}

func (l *List[T]) Delete(i int) *List[T] {
	l.elems = append(l.elems[0:i], l.elems[i+1:]...)
	return l
}

func main() {
	list := List[int]{}

	list.Add(1).Add(2).Add(3).Add(4).Add(5)
	fmt.Println(list)
	list.Update(1, 20)
	fmt.Println(list)
	list.Pop()
	fmt.Println(list)
	list.Delete(1)
	fmt.Println(list)
}
