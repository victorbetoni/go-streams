package test

import (
	"fmt"
	"testing"

	"github.com/victorbetoni/go-streams/comparators"
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

func TestJoinFunction(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{1, 10, 2, 10, 3, 10, 4, 10, 5, 10, 6}

	slice := streams.StreamOf[int](input).Join(10).ToSlice()

	compareSlices(*slice, expected, t)
}

func TestReduceFunction(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := 1 + 2 + 3 + 4 + 5 + 6

	reducted := streams.StreamOf[int](input).Reduce(0, func(e1, e2 int) int { return e1 + e2 })

	if expected != *reducted {
		t.Errorf("reducing gone wrong. expected %d and got %d", expected, *reducted)
	}
}

func TestAnyMatchFunction(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := true
	filter := func(val int) bool {
		return val == 3
	}

	result := streams.StreamOf[int](input).AnyMatch(filter)

	if expected != result {
		t.Errorf("expected %v and got %v", expected, result)
	}
}

func TestReversedFunction(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{6, 5, 4, 3, 2, 1}

	result := streams.StreamOf[int](input).Reversed()

	compareSlices(result.Current, expected, t)
}

func TestSortedFunction(t *testing.T) {
	input := []int{7, 3, 6, 2, 0, 5}
	expected := []int{0, 2, 3, 5, 6, 7}

	result := streams.StreamOf[int](input).Sorted(comparators.CompareInt)

	compareSlices(result.Current, expected, t)
}

func TestSkipFunction(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{3, 4, 5, 6}

	result := streams.StreamOf[int](input).Skip(2)

	compareSlices(result.Current, expected, t)
}

func TestLimitFunction(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{1, 2, 3}

	result := streams.StreamOf[int](input).Limit(3)

	compareSlices(result.Current, expected, t)
}
