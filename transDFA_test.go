package regexop

import (
	"fmt"
	"testing"
)

func TestTransDFAInsert(t *testing.T) {
	// (0, a) -> 1
	// (1, a) -> 2
	// (2, b) -> 3

	td := newTransDFA()
	td.insertSplit(0, 'a', 0)
	td.insertSplit(0, 'a', 1)
	td.insertSplit(1, 'a', 2)
	td.insertSplit(2, 'b', 3)
	fmt.Println(td)

}
