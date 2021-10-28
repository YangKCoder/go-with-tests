package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertSum := func(sum int, expected int, t testing.TB) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}
	t.Run("2+2=4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertSum(sum, expected, t)
	})

	t.Run("-2+2=0", func(t *testing.T) {
		sum := Add(-2, 2)
		expected := 0
		assertSum(sum, expected, t)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
