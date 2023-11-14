package tuples

type tuple7s[A, B, C, D, E, F, G any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
}

func Tuple7[A, B, C, D, E, F, G any](a A, b B, c C, d D, e E, f F, g G) tuple7s[A, B, C, D, E, F, G] {
	return tuple7s[A, B, C, D, E, F, G]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
	}
}

func (t *tuple7s[A, B, C, D, E, F, G]) Get1() A {
	return t.A
}

func (t *tuple7s[A, B, C, D, E, F, G]) Get2() B {
	return t.B
}

func (t *tuple7s[A, B, C, D, E, F, G]) Get3() C {
	return t.C
}

func (t *tuple7s[A, B, C, D, E, F, G]) Get4() D {
	return t.D
}

func (t *tuple7s[A, B, C, D, E, F, G]) Get5() E {
	return t.E
}

func (t *tuple7s[A, B, C, D, E, F, G]) Get6() F {
	return t.F
}

func (t *tuple7s[A, B, C, D, E, F, G]) Get7() G {
	return t.G
}
