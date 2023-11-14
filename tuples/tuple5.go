package tuples

type tuple5s[A, B, C, D, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

func Tuple5[A, B, C, D, E any](a A, b B, c C, d D, e E) tuple5s[A, B, C, D, E] {
	return tuple5s[A, B, C, D, E]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
	}
}

func (t *tuple5s[A, B, C, D, E]) Get1() A {
	return t.A
}

func (t *tuple5s[A, B, C, D, E]) Get2() B {
	return t.B
}

func (t *tuple5s[A, B, C, D, E]) Get3() C {
	return t.C
}

func (t *tuple5s[A, B, C, D, E]) Get4() D {
	return t.D
}

func (t *tuple5s[A, B, C, D, E]) Get5() E {
	return t.E
}
