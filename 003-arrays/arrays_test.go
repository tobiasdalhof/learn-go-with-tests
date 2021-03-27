package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3, 4, 5})
	want := 15
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{10, 10, 10}, []int{20, 20, 20}, []int{30, 30, 30})
		want := []int{20, 40, 60}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{20}, []int{20, 20})
		want := []int{0, 0, 20}
		checkSums(t, got, want)
	})
}
