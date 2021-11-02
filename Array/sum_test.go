package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("make the sums of some slices", func(t *testing.T) {

		got := SumAll([]int{1, 2, 3}, []int{1, 2})
		want := []int{6, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
