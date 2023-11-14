package tuples

type tuple10s[A, B, C, D, E, F, G, H, I, J any] struct {
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
	J J
}

func Tuple10[A, B, C, D, E, F, G, H, I, J any](a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) tuple10s[A, B, C, D, E, F, G, H, I, J] {
	return tuple10s[A, B, C, D, E, F, G, H, I, J]{
		A: a,
		B: b,
		C: c,
		D: d,
		E: e,
		F: f,
		G: g,
		H: h,
		I: i,
		J: j,
	}
}

func (t *tuple10s[A, B, C, D, E, F, G, H, I, J]) Get1() A {
	return t.A
}

func (t *tuple10s[A, B, C, D, E, F, G, H, I, J]) Get2() B {
	return t.B
}

func (t *tuple10s[A, B, C, D, E, F, G, H, I, J]) Get3() C {
	return t.C
}

func (t *tuple10s[A, B, C, D, E, F, G, H, I, J]) Get4() D {
	return t.D
}

func (t *tuple10s[A, B, C, D, E, F, G, H, I, J]) Get5() E {
	return t.E
}

func (t *tuple10s[A, B, C, D, E, F, G, H, I, J]) Get6() F {
	return t.F
}

func (t *tuple10s[A, B, C, D, E, F, G, H, I, J]) Get7() G {
	return t.G
}

func (t *tuple10s[A, B, C, D, E, F, G, H, I, J]) Get8() H {
	return t.H
}

func (t *tuple10s[A, B, C, D, E, F, G, H, I, J]) Get9() I {
	return t.I
}

func (t *tuple10s[A, B, C, D, E, F, G, H, I, J]) Get10() J {
	return t.J
}
