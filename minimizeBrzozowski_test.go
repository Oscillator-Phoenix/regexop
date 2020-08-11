package regexop

import (
	"fmt"
	"testing"
)

func TestDFAReverse1(t *testing.T) {
	var p parser

	d := p.regexToDFA("abb")
	fmt.Println("DFA")
	fmt.Println(d)

	rn := d.reverseToNFA()
	fmt.Println("reverse NFA")
	fmt.Println(rn)

	rd := d.reverseToDFA()
	fmt.Println("reverse DFA")
	fmt.Println(rd)

	md := d.reverseToDFA().reverseToDFA()
	fmt.Println("minimized DFA")
	fmt.Println(md)
}

func TestDFAReverse2(t *testing.T) {
	var p parser

	d := p.regexToDFA("a(a|b)*")
	fmt.Println("DFA")
	fmt.Println(d)

	rn := d.reverseToNFA()
	fmt.Println("reverse NFA")
	fmt.Println(rn)

	rd := d.reverseToDFA()
	fmt.Println("reverse DFA")
	fmt.Println(rd)

	md := d.reverseToDFA().reverseToDFA()
	fmt.Println("minimized DFA")
	fmt.Println(md)
}

func TestDFAReverse3(t *testing.T) {
	var p parser

	d := p.regexToDFA("(a|b)*")
	fmt.Println("DFA")
	fmt.Println(d)

	rn := d.reverseToNFA()
	fmt.Println("reverse NFA")
	fmt.Println(rn)

	rd := d.reverseToDFA()
	fmt.Println("reverse DFA")
	fmt.Println(rd)

	md := d.reverseToDFA().reverseToDFA()
	fmt.Println("minimized DFA")
	fmt.Println(md)
}

func TestDFAReverse4(t *testing.T) {
	var p parser

	d := p.regexToDFA("(a|b)*(a|b)*a")
	fmt.Println("DFA")
	fmt.Println(d)

	rn := d.reverseToNFA()
	fmt.Println("reverse NFA")
	fmt.Println(rn)

	rd := d.reverseToDFA()
	fmt.Println("reverse DFA")
	fmt.Println(rd)

	md := d.reverseToDFA().reverseToDFA()
	fmt.Println("minimized DFA")
	fmt.Println(md)
}

func TestDFAReverseToNFA1(t *testing.T) {
	var p parser

	d := p.regexToDFA("a|b")
	fmt.Println("DFA")
	fmt.Println(d)

	rn := d.reverseToNFA()
	fmt.Println("reverse NFA")
	fmt.Println(rn)
}

func TestDFAReverseToDFA1(t *testing.T) {
	var p parser

	d := p.regexToDFA("a|b")
	fmt.Println("DFA")
	fmt.Println(d)

	rd := d.reverseToDFA()
	fmt.Println("reverse DFA")
	fmt.Println(rd)
}

func TestDFAReverseToDFA2(t *testing.T) {
	var p parser

	d := p.regexToDFA("ab")
	fmt.Println("DFA")
	fmt.Println(d)

	rd := d.reverseToDFA()
	fmt.Println("reverse DFA")
	fmt.Println(rd)
}

func TestDFAReverseToDFA3(t *testing.T) {
	var p parser

	d := p.regexToDFA("(a|b)*abb")
	fmt.Println("DFA")
	fmt.Println(d)

	rd := d.reverseToDFA()
	fmt.Println("reverse DFA")
	fmt.Println(rd)
}
