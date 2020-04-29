package dependencyinjection

import "fmt"
import "io"

// So hello to a fellow human
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
