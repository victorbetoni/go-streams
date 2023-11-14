package streams

import (
	"math"

	"github.com/victorbetoni/go-streams/sort"
)

type Predicate func(any) bool
type Function func(any) any

type Stream[E any] struct {
	Current []E
}

func StreamOf[E any](slice ...E) *Stream[E] {
	return &Stream[E]{
		Current: slice,
	}
}

func Empty[E any]() *Stream[E] {
	s := make([]E, 0)
	return &Stream[E]{
		Current: s,
	}
}

func (s *Stream[E]) Lenght() int {
	return len(s.Current)
}

func (s *Stream[E]) Filter(filter func(E) bool) *Stream[E] {
	slice := make([]E, 0)
	for _, val := range s.Current {
		if filter(val) {
			slice = append(slice, val)
		}
	}
	return StreamOf[E](slice...)
}

func (s *Stream[E]) ForEach(consumer func(E)) {
	for _, val := range s.Current {
		consumer(val)
	}
}

func (s *Stream[E]) Join(joined E) *Stream[E] {
	slice := make([]E, 0)
	for i, val := range s.Current {
		slice = append(slice, val)
		if i != len(s.Current)-1 {
			slice = append(slice, joined)
		}
	}
	return StreamOf[E](slice...)
}

func (s *Stream[E]) Skip(n int) *Stream[E] {
	if len(s.Current) <= n {
		return Empty[E]()
	}
	return StreamOf[E](s.Current[n:]...)
}

func (s *Stream[E]) Limit(n int) *Stream[E] {
	if len(s.Current) <= n {
		return StreamOf[E](s.Current...)
	}
	return StreamOf[E](s.Current[:n+1]...)
}

func (s *Stream[E]) Sorted(comparator func(e1, e2 E) int) *Stream[E] {
	arr := sort.QuickSort[E](s.Current, comparator, 0, len(s.Current)-1)
	return StreamOf[E](arr...)
}

func (s *Stream[E]) Reversed() *Stream[E] {
	upper, lower := len(s.Current)-1, 0
	d := int(math.Floor(float64(len(s.Current)) / 2))
	for i := 0; i < d; i++ {
		s.Current[upper-i], s.Current[lower+i] = s.Current[lower+i], s.Current[upper-i]
	}
	return StreamOf[E](s.Current...)
}

func (s *Stream[E]) Reduce(identity E, reductor func(E, E) E) *E {
	for _, val := range s.Current {
		identity = reductor(identity, val)
	}
	return &identity
}

func (s *Stream[E]) AnyMatch(predicate func(E) bool) bool {
	for _, val := range s.Current {
		if predicate(val) {
			return true
		}
	}
	return false
}

func (s *Stream[E]) NoneMatch(predicate func(E) bool) bool {
	return !s.AnyMatch(predicate)
}

func (s *Stream[E]) AllMatch(predicate func(E) bool) bool {
	for _, val := range s.Current {
		if !predicate(val) {
			return false
		}
	}
	return true
}

func (s *Stream[E]) FindFirst() *E {
	if len(s.Current) < 1 {
		return nil
	}
	return &s.Current[0]
}

func ToMap[T comparable, V any](stream *Stream[T], mapper func(T) V) map[T]V {
	m := make(map[T]V, len(stream.Current))
	for _, val := range stream.Current {
		m[val] = mapper(val)
	}
	return m
}

func (s *Stream[E]) ToSlice() *[]E {
	return &s.Current
}
