package generics

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	var source = []int{1, 2, 3, 4, 5}
	var result = Map(source, func(value int) int {
		return value * 2
	})
	if reflect.DeepEqual(result, []int{2, 4, 6, 8, 10}) == false {
		t.Error("Expected: [2, 4, 6, 8, 10], Actual: ", result)
	}
}

func TestMapIndexed(t *testing.T) {
	var source = []string{"Hello", "World"}
	var result = MapIndexed(source, func(value string, index int) int {
		return index
	})
	if reflect.DeepEqual(result, []int{0, 1}) == false {
		t.Error("Expected: [0, 1], Actual: ", result)
	}
}

func TestFilterInteger(t *testing.T) {
	var source = []int{1, 2, 3, 4, 5}
	var result = Filter(source, func(value int) bool {
		return value%2 == 0
	})
	if reflect.DeepEqual(result, []int{2, 4}) == false {
		t.Error("Expected: [2, 4], Actual: ", result)
	}
}

func TestFilterString(t *testing.T) {
	var source = []string{"a", "bb", "cc", "d", "ee"}
	var result = Filter(source, func(value string) bool {
		return len(value) == 1
	})
	if reflect.DeepEqual(result, []string{"a", "d"}) == false {
		t.Error("Expected: [a, d], Actual: ", result)
	}
}
