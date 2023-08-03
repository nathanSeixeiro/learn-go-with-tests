package integers

import (
	"fmt"
	"testing"
)

func Add(x, y int) int {
	return x + y
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expect := 4

	if sum != expect {
		t.Errorf("diferentes values:\n expected: %d \n got: %d ", expect, sum)
	}
}
