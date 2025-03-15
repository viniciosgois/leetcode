package binarysearch

import (
	"testing"
)

func TestMaximumCountOfPositiveIntegerAndNegativeInteger(t *testing.T) {
	nums := []int{-462, -333, -221, -144, 143, 273, 331, 443, 666}

	t.Run("find", func(t *testing.T) {
		got := MaximumCountOfPositiveIntegerAndNegativeInteger(nums)
		want := 5

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
