package regexop

import (
	"fmt"
	"testing"
)

func TestMinimizeHopcroft1(t *testing.T) {
	states := newStateSet()
	states.insert(1, 2)
	alphbet := newSymbolSet()
	alphbet.insert('a', 'b')
	initial := state(1)
	finals := newStateSet()
	finals.insert(1, 2)
	trans := newTransDFA()
	trans.insertSplit(1, 'a', 2)
	trans.insertSplit(1, 'b', 2)
	trans.insertSplit(2, 'a', 2)
	trans.insertSplit(2, 'b', 2)

	d := &dfa{alphbet, states, initial, finals, trans}
	fmt.Println("d")
	fmt.Println(d)

	md := d.minimizeHopcroft()
	fmt.Println("md")
	fmt.Println(md)
}

func TestMinimizeHopcroft2(t *testing.T) {
	states := newStateSet()
	states.insert(1, 2, 3, 4, 5)
	alphbet := newSymbolSet()
	alphbet.insert('a', 'b')
	initial := state(1)
	finals := newStateSet()
	finals.insert(5)
	trans := newTransDFA()

	trans.insertSplit(1, 'a', 2)
	trans.insertSplit(2, 'a', 2)
	trans.insertSplit(3, 'a', 2)
	trans.insertSplit(4, 'a', 2)
	trans.insertSplit(5, 'a', 2)

	trans.insertSplit(1, 'b', 3)
	trans.insertSplit(2, 'b', 4)
	trans.insertSplit(3, 'b', 3)
	trans.insertSplit(4, 'b', 5)
	trans.insertSplit(5, 'b', 3)

	d := &dfa{alphbet, states, initial, finals, trans}
	fmt.Println("d")
	fmt.Println(d)

	md := d.minimizeHopcroft()
	fmt.Println("md")
	fmt.Println(md)
}

func TestMinimizeHopcroft3(t *testing.T) {
	states := newStateSet()
	states.insert(0, 1, 2, 3, 4, 5)
	alphbet := newSymbolSet()
	alphbet.insert('a', 'b')
	initial := state(0)
	finals := newStateSet()
	finals.insert(2, 3, 4)
	trans := newTransDFA()

	trans.insertSplit(0, 'a', 1)
	trans.insertSplit(1, 'a', 0)
	trans.insertSplit(2, 'a', 4)
	trans.insertSplit(3, 'a', 4)
	trans.insertSplit(4, 'a', 4)
	trans.insertSplit(5, 'a', 5)

	trans.insertSplit(0, 'b', 2)
	trans.insertSplit(1, 'b', 3)
	trans.insertSplit(2, 'b', 5)
	trans.insertSplit(3, 'b', 5)
	trans.insertSplit(4, 'b', 5)
	trans.insertSplit(5, 'b', 5)

	d := &dfa{alphbet, states, initial, finals, trans}
	fmt.Println("d")
	fmt.Println(d)

	md := d.minimizeHopcroft()
	fmt.Println("md")
	fmt.Println(md)
}

func TestMinimizeHopcroft4(t *testing.T) {
	var p parser

	d := p.regexToDFA("(a|b)*")
	fmt.Println("DFA")
	fmt.Println(d)

	md := d.minimizeHopcroft()
	fmt.Println("minimized DFA")
	fmt.Println(md)
}

func TestMinimizeHopcroft5(t *testing.T) {
	var p parser

	d := p.regexToDFA("(a|b)*abb")
	fmt.Println("DFA")
	fmt.Println(d)

	md := d.minimizeHopcroft()
	fmt.Println("minimized DFA")
	fmt.Println(md)
}
