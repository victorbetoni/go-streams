package function

func Identity[T any]() func(T) T {
	return func(t T) T { return t }
}
