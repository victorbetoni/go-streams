package tuples

type tuple9s[A, B, C, D, E, F, G, H, I any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
}

func Tuple9[A, B, C, D, E, F, G, H, I any](a A, b B, c C, d D, e E, f F, g G, h H, i I) tuple9s[A, B, C, D, E, F, G, H, I] {
	return tuple9s[A, B, C, D, E, F, G, H, I]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
		H: h,
		I: i,
	}
}

func (t *tuple9s[A, B, C, D, E, F, G, H, I]) Get1() A {
	return t.A
}

func (t *tuple9s[A, B, C, D, E, F, G, H, I]) Get2() B {
	return t.B
}

func (t *tuple9s[A, B, C, D, E, F, G, H, I]) Get3() C {
	return t.C
}

func (t *tuple9s[A, B, C, D, E, F, G, H, I]) Get4() D {
	return t.D
}

func (t *tuple9s[A, B, C, D, E, F, G, H, I]) Get5() E {
	return t.E
}

func (t *tuple9s[A, B, C, D, E, F, G, H, I]) Get6() F {
	return t.F
}

func (t *tuple9s[A, B, C, D, E, F, G, H, I]) Get7() G {
	return t.G
}

func (t *tuple9s[A, B, C, D, E, F, G, H, I]) Get8() H {
	return t.H
}

func (t *tuple9s[A, B, C, D, E, F, G, H, I]) Get9() I {
	return t.I
}
