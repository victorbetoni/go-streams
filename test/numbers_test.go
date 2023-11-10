package test

import (
	"fmt"
	"testing"

	"github.com/victorbetoni/go-streams/streams"
)

func TestMapFunction(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []string{"1", "2", "3", "4", "5"}

	slice := streams.Map[int, string](streams.StreamOf[int](input), func(i int) string {
		return fmt.Sprintf("%d", i)
	}).ToSlice()

	for i, val := range *slice {
		if val != expected[i] {
			t.Errorf("error in position %d. Expected %s and got %s", i, val, expected[i])
		}
	}
}

func TestFilterFunction(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{4, 5}

	slice := streams.StreamOf[int](input).Filter(func(e int) bool {
		return e > 3
	}).ToSlice()

	compareSlices(*slice, expected, t)
}
