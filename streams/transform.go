package streams

import (
	"math"

	"github.com/victorbetoni/go-streams/tuples"
)

func Map[E, V any](stream *Stream[E], mapper func(E) V) *Stream[V] {
	slice := make([]V, 0)
	for _, val := range stream.Current {
		slice = append(slice, mapper(val))
	}
	return StreamOf[V](slice)
}

func FlatMap[E, V any](stream *Stream[E], mapper func(E) []V) *Stream[V] {
	slice := make([]V, 0)
	for _, val := range stream.Current {
		slice = append(slice, mapper(val)...)
	}
	return StreamOf[V](slice)
}

func Zip[E, V, T any](s1 *Stream[E], s2 *Stream[V], zipper func(E, V) T) *Stream[T] {
	max := int(math.Max(float64(s1.Lenght()), float64(s2.Lenght())))
	s := make([]T, 0)
	for i := 0; i < max; i++ {
		var v1 E
		var v2 V
		if s1.Lenght()-1 >= i {
			v1 = s1.Current[i]
		}
		if s2.Lenght()-1 >= i {
			v2 = s2.Current[i]
		}
		s = append(s, zipper(v1, v2))
	}
	return StreamOf[T](s)
}

func ZipToTuples[E, V any](s1 *Stream[E], s2 *Stream[V]) *Stream[tuples.Tuple2S[E, V]] {
	max := int(math.Max(float64(s1.Lenght()), float64(s2.Lenght())))
	s := make([]tuples.Tuple2S[E, V], 0)
	for i := 0; i < max; i++ {
		var v1 E
		var v2 V
		if s1.Lenght()-1 >= i {
			v1 = s1.Current[i]
		}
		if s2.Lenght()-1 >= i {
			v2 = s2.Current[i]
		}
		s = append(s, tuples.Tuple2[E, V](v1, v2))
	}
	return StreamOf[tuples.Tuple2S[E, V]](s)
}
