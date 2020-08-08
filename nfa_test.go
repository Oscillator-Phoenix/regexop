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

func exampleNFA1() *nfa {
	// (a|b)*abb
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

	n := nfa{alphbet, states, initial, finals, trans}

	return &n
}

func exampleNFA2() *nfa {
	// (a|b)*abb
	states := newStateSet()
	alphbet := newSymbolSet()
	initial := state(0)
	finals := newStateSet()
	trans := newTransNFA()

	states.insert(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	alphbet.insert('a', 'b') // without epsilon
	finals.insert(10)

	trans.insertSplitSingle(0, constEpsilon, 1, 7)
	trans.insertSplitSingle(1, constEpsilon, 2, 4)
	trans.insertSplitSingle(2, 'a', 3)
	trans.insertSplitSingle(3, constEpsilon, 6)
	trans.insertSplitSingle(4, 'b', 5)
	trans.insertSplitSingle(5, constEpsilon, 6)
	trans.insertSplitSingle(6, constEpsilon, 1, 7)
	trans.insertSplitSingle(7, 'a', 8)
	trans.insertSplitSingle(8, 'b', 9)
	trans.insertSplitSingle(9, 'b', 10)

	n := nfa{alphbet, states, initial, finals, trans}

	return &n
}

func TestNFA2DFA1(t *testing.T) {
	d := exampleNFA1().toDFA()
	fmt.Println("\n\nDFA1")
	fmt.Println(d)
}

func TestNFA2DFA2(t *testing.T) {
	d := exampleNFA2().toDFA()
	fmt.Println("\n\nDFA1")
	fmt.Println(d)
}

func TestNFA2DFA3(t *testing.T) {
	var p parser
	regex := "(a|b)*abb"
	d := p.regexToDFA(regex)
	fmt.Println(d)
}
