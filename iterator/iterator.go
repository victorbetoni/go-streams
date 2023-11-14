package iterator

import "github.com/victorbetoni/go-streams/streams"

type Iterator[E any] struct {
	collection []E
	Current    *E
	Index      int
}

func SliceIterator[E any](slice *[]E) *Iterator[E] {
	return StreamIterator[E](streams.StreamOf[E](*slice...))
}

func StreamIterator[E any](stream *streams.Stream[E]) *Iterator[E] {
	return &Iterator[E]{
		collection: stream.Current,
		Current:    nil,
		Index:      -1,
	}
}

func (i *Iterator[E]) HasNext() bool {
	return i.Index+1 < len(i.collection)
}

func (i *Iterator[E]) HasPrevious() bool {
	return i.Index-1 >= 0
}

func (i *Iterator[E]) Next() bool {
	if !i.HasNext() {
		return false
	}
	i.Index++
	i.Current = &i.collection[i.Index]
	return true
}

func (i *Iterator[E]) Previous() bool {
	if !i.HasPrevious() {
		return false
	}
	i.Index--
	i.Current = &i.collection[i.Index]
	return true
}

func (i *Iterator[E]) ForEachRemaning(do func(*E)) {
	for i.Next() {
		do(i.Current)
	}
}
