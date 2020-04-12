package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeatCount := 10
	repeated := Repeat("a", repeatCount)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 500)
	}
}

func ExampleRepeat() {
	repeat := Repeat("K", 5)
	fmt.Println(repeat)
	// Output: KKKKK
}
