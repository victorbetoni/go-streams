package tuples

type tuple4s[A, B, C, D any] struct {
	A A
	B B
	C C
	D D
}

func Tuple4[A, B, C, D any](a A, b B, c C, d D) tuple4s[A, B, C, D] {
	return tuple4s[A, B, C, D]{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

func (t *tuple4s[A, B, C, D]) Get1() A {
	return t.A
}

func (t *tuple4s[A, B, C, D]) Get2() B {
	return t.B
}

func (t *tuple4s[A, B, C, D]) Get3() C {
	return t.C
}

func (t *tuple4s[A, B, C, D]) Get4() D {
	return t.D
}
