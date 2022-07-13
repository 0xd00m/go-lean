package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := add(2, 2)
	want := 4

	if sum != want {
		t.Errorf("want %d got %d", want, sum)
	}
}

func ExampleAdd() {
	sum := add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
