package streams

type Predicate func(any) bool
type Function func(any) any

type Stream[E any] struct {
	Current []E
}

func StreamOf[E any](slice []E) *Stream[E] {
	return &Stream[E]{
		Current: slice,
	}
}

func Map[E, V any](stream *Stream[E], mapper func(E) V) *Stream[V] {
	slice := make([]V, 0)
	for _, val := range stream.Current {
		slice = append(slice, mapper(val))
	}
	return StreamOf[V](slice)
}

func (s *Stream[E]) Filter(filter func(E) bool) *Stream[E] {
	slice := make([]E, 0)
	for _, val := range s.Current {
		if filter(val) {
			slice = append(slice, val)
		}
	}
	return StreamOf[E](slice)
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
	return StreamOf[E](slice)
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
