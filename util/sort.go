package util

import (
	"sort"
)

type sortableWrapper[T any] struct {
	vs []T
	less func(a, b T) bool
}

func (w *sortableWrapper[T]) Less(i, j int) bool {
	return w.less(w.vs[i], w.vs[j])
}

func (w *sortableWrapper[T]) Len() int {
	return len(w.vs)
}

func (w *sortableWrapper[T]) Swap(i, j int) {
	w.vs[i], w.vs[j] = w.vs[j], w.vs[i]
}

func Sort[T any](vs []T, less func(a, b T) bool) {
	sort.Sort(&sortableWrapper[T]{
		vs: vs,
		less: less,
	})
}