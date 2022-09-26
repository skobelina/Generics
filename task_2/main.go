package main

import "fmt"

type Map[K comparable, V any] struct {
	innerMap map[K]V
}

// Task: create Map[K, V] with methods Add(K, V), Delete(K), Update(K, V), Get(K)

func (m *Map[K, V]) Add(k K, v V) *Map[K, V] {
	m.innerMap[k] = v
	return m
}

func (m *Map[K, V]) Delete(k K) *Map[K, V] {
	delete(m.innerMap, k)
	return m
}

func (m *Map[K, V]) Update(k K, v V) *Map[K, V] {
	m.innerMap[k] = v
	return m
}

func (m *Map[K, V]) Get(k K) V {
	return m.innerMap[k]
}

func main() {
	m := Map[int, string]{innerMap: make(map[int]string)}
	m.Add(1, "one").Add(2, "two").Add(3, "three")
	fmt.Println(m)
	m.Delete(2)
	fmt.Println(m)
	m.Update(3, "threee")
	fmt.Println(m)
	fmt.Println(m.Get(3))
}
