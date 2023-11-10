package streams

type Float interface {
	float64 | float32
}

type Integer interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

type Complex interface {
	complex64 | complex128
}

type Number interface {
	Float | Integer | Complex
}

type NumberStream[E Number] struct {
	Stream[E]
}

func (n *NumberStream[E]) Sum() E {
	var identity E
	switch any(identity).(type) {
	case float64, float32:
		return *n.Reduce(0.0, func(e1, e2 E) E { return e1 + e2 })
	case complex64, complex128:
		return *n.Reduce(complex(0, 0), func(e1, e2 E) E { return e1 + e2 })
	default:
		return *n.Reduce(0, func(e1, e2 E) E { return e1 + e2 })
	}
}

func (n *Stream[E]) MapToFloat32(mapper func(E) float32) *NumberStream[float32] {
	return mapToNumber[E, float32](n, mapper)
}

func (n *Stream[E]) MapToFloat64(mapper func(E) float64) *NumberStream[float64] {
	return mapToNumber[E, float64](n, mapper)
}

func (n *Stream[E]) MapToInt(mapper func(E) int) *NumberStream[int] {
	return mapToNumber[E, int](n, mapper)
}

func (n *Stream[E]) MapToInt8(mapper func(E) int8) *NumberStream[int8] {
	return mapToNumber[E, int8](n, mapper)
}

func (n *Stream[E]) MapToInt16(mapper func(E) int16) *NumberStream[int16] {
	return mapToNumber[E, int16](n, mapper)
}

func (n *Stream[E]) MapToInt32(mapper func(E) int32) *NumberStream[int32] {
	return mapToNumber[E, int32](n, mapper)
}

func (n *Stream[E]) MapToInt64(mapper func(E) int64) *NumberStream[int64] {
	return mapToNumber[E, int64](n, mapper)
}

func (n *Stream[E]) MapToUint(mapper func(E) uint) *NumberStream[uint] {
	return mapToNumber[E, uint](n, mapper)
}

func (n *Stream[E]) MapToUint8(mapper func(E) uint8) *NumberStream[uint8] {
	return mapToNumber[E, uint8](n, mapper)
}

func (n *Stream[E]) MapToUint16(mapper func(E) uint16) *NumberStream[uint16] {
	return mapToNumber[E, uint16](n, mapper)
}

func (n *Stream[E]) MapToUint32(mapper func(E) uint32) *NumberStream[uint32] {
	return mapToNumber[E, uint32](n, mapper)
}

func (n *Stream[E]) MapToUint64(mapper func(E) uint64) *NumberStream[uint64] {
	return mapToNumber[E, uint64](n, mapper)
}

func (n *Stream[E]) MapToComplex64(mapper func(E) complex64) *NumberStream[complex64] {
	return mapToNumber[E, complex64](n, mapper)
}

func (n *Stream[E]) MapToUintptr(mapper func(E) uintptr) *NumberStream[uintptr] {
	return mapToNumber[E, uintptr](n, mapper)
}

func mapToNumber[E any, T Number](stream *Stream[E], mapper func(E) T) *NumberStream[T] {
	numbers := make([]T, len(stream.Current))
	for i, val := range stream.Current {
		numbers[i] = mapper(val)
	}
	return &NumberStream[T]{
		Stream: Stream[T]{Current: numbers},
	}
}
