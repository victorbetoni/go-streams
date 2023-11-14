package tuples

type tuple3s[A, B, C any] struct {
	A A
	B B
	C C
}

func Tuple3[A, B, C any](a A, b B, c C) tuple3s[A, B, C] {
	return tuple3s[A, B, C]{
		A: a,
		B: b,
		C: c,
	}
}

func (t *tuple3s[A, B, C]) Get1() A {
	return t.A
}

func (t *tuple3s[A, B, C]) Get2() B {
	return t.B
}

func (t *tuple3s[A, B, C]) Get3() C {
	return t.C
}
