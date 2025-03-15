package binarysearch

import "testing"

func TestSearchInsert(t *testing.T) {
	t.Run("with mapped index", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6}
		got := SearchInsert(nums, 5)
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("with unmapped index", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6}
		got := SearchInsert(nums, 7)
		want := 6

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
