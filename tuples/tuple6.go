package tuples

type tuple6s[A, B, C, D, E, F any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
}

func Tuple6[A, B, C, D, E, F any](a A, b B, c C, d D, e E, f F) tuple6s[A, B, C, D, E, F] {
	return tuple6s[A, B, C, D, E, F]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
	}
}

func (t *tuple6s[A, B, C, D, E, F]) Get1() A {
	return t.A
}

func (t *tuple6s[A, B, C, D, E, F]) Get2() B {
	return t.B
}

func (t tuple6s[A, B, C, D, E, F]) Get3() C {
	return t.C
}

func (t *tuple6s[A, B, C, D, E, F]) Get4() D {
	return t.D
}

func (t *tuple6s[A, B, C, D, E, F]) Get5() E {
	return t.E
}

func (t *tuple6s[A, B, C, D, E, F]) Get6() F {
	return t.F
}
