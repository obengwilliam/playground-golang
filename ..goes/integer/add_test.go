package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	fmt.Println(Add(2, 3))
	// Output: 5
}
func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if want != got {
		t.Errorf("What i got %d, is not what i want %d", got, want)
	}
}
