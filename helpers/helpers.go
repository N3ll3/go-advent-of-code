package helpers

import (
	"fmt"

	"github.com/kkrypt0nn/swiss-army-knife/aoc"
)

func main() {
    str := "hello world"
    substr := "world"
    fmt.Println(aoc.ContainsExactly(str, substr))
    // Output: false
    str = "world"
    substr = "dlrow"
    fmt.Println(aoc.ContainsExactly(str, substr))
    // Output: true
}