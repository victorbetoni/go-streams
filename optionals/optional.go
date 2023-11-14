package optionals

type Optional[E any] struct {
	Wrapped *E
}

func Of[E any](val E) *Optional[E] {
	return &Optional[E]{
		Wrapped: &val,
	}
}

func (o *Optional[E]) GetOrElse(val E) *E {
	if o.Wrapped == nil {
		return &val
	}
	return o.Wrapped
}

func (o *Optional[E]) GetOrPanic(err error) *E {
	if o.Wrapped == nil {
		panic(err)
	}
	return o.Wrapped

}

func (o *Optional[E]) GetOrElseDo(do func()) *E {
	if o.Wrapped == nil {
		do()
	}
	return o.Wrapped

}

func (o *Optional[E]) IfPresent(do func(*E)) {
	if o.Wrapped != nil {
		do(o.Wrapped)
	}
}

func (o *Optional[E]) IfAbsent(do func()) {
	if o.Wrapped == nil {
		do()
	}
}
