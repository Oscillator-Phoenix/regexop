package regexop

import (
	"fmt"
	"testing"
)

func TestNFAString(t *testing.T) {
	states := newStateSet()
	alphbet := newSymbolSet()
	initial := state(0)
	finals := newStateSet()
	trans := newTransNFA()

	states.insert(0, 1, 2, 3)
	alphbet.insert('a', 'b')
	finals.insert(3)

	trans.insertSplitSingle(0, 'a', 0, 1)
	trans.insertSplitSingle(0, 'b', 0)
	trans.insertSplitSingle(1, 'b', 2)
	trans.insertSplitSingle(2, 'b', 3)

	// (a|b)*abb
	n := nfa{alphbet, states, initial, finals, trans}

	fmt.Println(n)
}
