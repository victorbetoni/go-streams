package main

import (
	"fmt"

	"github.com/victorbetoni/go-streams/function"
	"github.com/victorbetoni/go-streams/streams"
)

func main() {
	a := []int64{1, 2, 3, 4, 5, 6}
	fmt.Printf("%d", streams.StreamOf[int64](a).Join(10).MapToInt64(function.Identity[int64]()).Sum())
}
