package iteration

import "testing"
import "fmt"

func ExampleRepeat() {
	fmt.Println(Repeat("b", 5))
	//Output:bbbbb
}

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("Got something  %s instead of %s", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
