package tuples

type tuple8s[A, B, C, D, E, F, G, H any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
}

func Tuple8[A, B, C, D, E, F, G, H any](a A, b B, c C, d D, e E, f F, g G, h H) tuple8s[A, B, C, D, E, F, G, H] {
	return tuple8s[A, B, C, D, E, F, G, H]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
		H: h,
	}
}

func (t *tuple8s[A, B, C, D, E, F, G, H]) Get1() A {
	return t.A
}

func (t *tuple8s[A, B, C, D, E, F, G, H]) Get2() B {
	return t.B
}

func (t *tuple8s[A, B, C, D, E, F, G, H]) Get3() C {
	return t.C
}

func (t *tuple8s[A, B, C, D, E, F, G, H]) Get4() D {
	return t.D
}

func (t *tuple8s[A, B, C, D, E, F, G, H]) Get5() E {
	return t.E
}

func (t *tuple8s[A, B, C, D, E, F, G, H]) Get6() F {
	return t.F
}

func (t *tuple8s[A, B, C, D, E, F, G, H]) Get7() G {
	return t.G
}

func (t *tuple8s[A, B, C, D, E, F, G, H]) Get8() H {
	return t.H
}
