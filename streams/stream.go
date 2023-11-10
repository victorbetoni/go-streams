package streams

type Stream[E any] struct {
	Current *[]E
}

func StreamOf[E any](slice []E) *Stream[E] {
	return &Stream[E]{
		Current: &slice,
	}
}

func Map[T, E any](stream *Stream[T], mapper func(T) E) *Stream[E] {
	slice := make([]E, 0)
	for _, val := range *stream.Current {
		slice = append(slice, mapper(val))
	}
	return StreamOf[E](slice)
}

func Filter[T any](stream *Stream[T], filter func(T) bool) *Stream[T] {
	slice := make([]T, 0)
	for _, val := range *stream.Current {
		if filter(val) {
			slice = append(slice, val)
		}
	}
	return StreamOf[T](slice)
}

func Join[T any](stream *Stream[T], joined T) *Stream[T] {
	slice := make([]T, 0)
	for i, val := range *stream.Current {
		slice = append(slice, val)
		if i != len(*stream.Current)-1 {
			slice = append(slice, joined)
		}
	}
	return StreamOf[T](slice)
}

func Reduce[T any](stream *Stream[T], initial T, reductor func(T, T) T) *T {
	for _, val := range *stream.Current {
		initial = reductor(initial, val)
	}
	return &initial
}

func AnyMatch[T any](stream *Stream[T], predicate func(T) bool) bool {
	for _, val := range *stream.Current {
		if predicate(val) {
			return true
		}
	}
	return false
}

func NoneMatch[T any](stream *Stream[T], predicate func(T) bool) bool {
	return !AnyMatch[T](stream, predicate)
}

func ToMap[T comparable, V any](stream *Stream[T], mapper func(T) V) map[T]V {
	m := make(map[T]V, len(*stream.Current))
	for _, val := range *stream.Current {
		m[val] = mapper(val)
	}
	return m
}
