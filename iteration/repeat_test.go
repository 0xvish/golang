package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("repeat character different number of times", func(t *testing.T) {
		repeated := Repeat("v", 1)
		expected := "v"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat character specific number of times", func(t *testing.T) {
		repeated := Repeat("v", 7)
		expected := "vvvvvvv"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}

func ExampleRepeat() {
	repeated := Repeat("v", 7)
	fmt.Println(repeated)
	// Output: vvvvvvv
}

func BenchmarkRepeat (b *testing.B) {
	for i := 0; i < b.N ; i++ {
		Repeat("a", 10)
	}
}