package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("should handle a slice ", func(t *testing.T) {
		var sliceArray = []int{2, 3, 4, 5}
		got := Sum(sliceArray)
		want := 14
		if got != want {
			t.Errorf("Expected the sum of %v to be %d  but got %d", sliceArray, want, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{2, 3, 4, 5}, []int{23, 5})
	want := []int{14, 28}

	if !reflect.DeepEqual(got, want) {
		t.Error("Not correct")
	}
}
func TestSumAllTails(t *testing.T) {

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
