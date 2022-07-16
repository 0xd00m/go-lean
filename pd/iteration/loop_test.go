package loop

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	got := Loop("a", 7)
	want := "aaaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func ExampleLoop() {
	result := Loop("a", 7)
	fmt.Println(result)
	// Output: aaaaaa
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop("a", 7)
	}

}
