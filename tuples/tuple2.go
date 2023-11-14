package tuples

type Tuple2S[A, B any] struct {
	A A
	B B
}

func Tuple2[A, B any](a A, b B) Tuple2S[A, B] {
	return Tuple2S[A, B]{
		A: a,
		B: b,
	}
}

func (t *Tuple2S[A, B]) Get1() A {
	return t.A
}

func (t *Tuple2S[A, B]) Get2() B {
	return t.B
}
