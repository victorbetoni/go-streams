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

func Sum[E Number](stream *NumberStream[E]) E {
	var identity E
	switch any(identity).(type) {
	case float64, float32:
		return *Reduce[E](&stream.Stream, 0.0, func(e1, e2 E) E { return e1 + e2 })
	case complex64, complex128:
		return *Reduce[E](&stream.Stream, complex(0, 0), func(e1, e2 E) E { return e1 + e2 })
	default:
		return *Reduce[E](&stream.Stream, 0, func(e1, e2 E) E { return e1 + e2 })
	}
}

func MapToFloat32[T any](stream *Stream[T], predicate func(T) float32) *NumberStream[float32] {
	return mapToNumber[T, float32](stream, predicate)
}

func MapToFloat64[T any](stream *Stream[T], predicate func(T) float64) *NumberStream[float64] {
	return mapToNumber[T, float64](stream, predicate)
}

func MapToInt[T any](stream *Stream[T], predicate func(T) int) *NumberStream[int] {
	return mapToNumber[T, int](stream, predicate)
}

func MapToInt8[T any](stream *Stream[T], predicate func(T) int8) *NumberStream[int8] {
	return mapToNumber[T, int8](stream, predicate)
}

func MapToInt16[T any](stream *Stream[T], predicate func(T) int16) *NumberStream[int16] {
	return mapToNumber[T, int16](stream, predicate)
}

func MapToInt32[T any](stream *Stream[T], predicate func(T) int32) *NumberStream[int32] {
	return mapToNumber[T, int32](stream, predicate)
}

func MapToInt64[T any](stream *Stream[T], predicate func(T) int64) *NumberStream[int64] {
	return mapToNumber[T, int64](stream, predicate)
}

func MapToUint[T any](stream *Stream[T], predicate func(T) uint) *NumberStream[uint] {
	return mapToNumber[T, uint](stream, predicate)
}

func MapToUint8[T any](stream *Stream[T], predicate func(T) uint8) *NumberStream[uint8] {
	return mapToNumber[T, uint8](stream, predicate)
}

func MapToUint16[T any](stream *Stream[T], predicate func(T) uint16) *NumberStream[uint16] {
	return mapToNumber[T, uint16](stream, predicate)
}

func MapToUint32[T any](stream *Stream[T], predicate func(T) uint32) *NumberStream[uint32] {
	return mapToNumber[T, uint32](stream, predicate)
}

func MapToUint64[T any](stream *Stream[T], predicate func(T) uint64) *NumberStream[uint64] {
	return mapToNumber[T, uint64](stream, predicate)
}

func MapToComplex64[T any](stream *Stream[T], predicate func(T) complex64) *NumberStream[complex64] {
	return mapToNumber[T, complex64](stream, predicate)
}

func MapToUintptr[T any](stream *Stream[T], predicate func(T) uintptr) *NumberStream[uintptr] {
	return mapToNumber[T, uintptr](stream, predicate)
}

func mapToNumber[E any, T Number](stream *Stream[E], mapper func(E) T) *NumberStream[T] {
	numbers := make([]T, len(*stream.Current))
	for i, val := range *stream.Current {
		numbers[i] = mapper(val)
	}
	return &NumberStream[T]{
		Stream: Stream[T]{Current: &numbers},
	}
}
