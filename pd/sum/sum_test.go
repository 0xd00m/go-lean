package sum

import "testing"

func TestSum(t *testing.T) {
	num := [10]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}

	got := Sum(num)
	want := 30

	if got != want {
		t.Errorf("got %q want %q, which array %v", got, want, num)
	}

}
