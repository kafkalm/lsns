package util

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](vs ...T) T {
	max := vs[0]
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return max
}

func Min[T constraints.Ordered](vs ...T) T {
	min := vs[0]
	for _, v := range vs {
		if v < min {
			min = v
		}
	}
	return min
}