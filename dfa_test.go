package regexop

import (
	"fmt"
	"testing"
)

func TestDFAIntersection1(t *testing.T) {
	var p parser
	d1 := p.regexToDFA("a|b")
	d2 := p.regexToDFA("a")
	d3 := intersectionTwoDFA(d1, d2)
	fmt.Println(d3)
}

func TestDFAIntersection2(t *testing.T) {
	var p parser
	d1 := p.regexToDFA("a|b")
	d2 := p.regexToDFA("a|b")

	fmt.Println(d1)
	fmt.Println(d2)

	d3 := intersectionTwoDFA(d1, d2)

	fmt.Println(d3)
}

func TestDFAIntersection3(t *testing.T) {
	var p parser
	d1 := p.regexToDFA("(a|b)*")
	d2 := p.regexToDFA("a(a|b)*")

	fmt.Println(d1)
	fmt.Println(d2)

	d3 := intersectionTwoDFA(d1, d2)

	fmt.Println(d3)
}

func TestDFAComplement1(t *testing.T) {
	var p parser

	d1 := p.regexToDFA("(a|b)*")
	fmt.Println(d1)

	d2 := d1.complement()
	fmt.Println(d2)

	d3 := d2.complement()
	fmt.Println(d3)
}

func TestDFAComplement2(t *testing.T) {
	var p parser

	n1 := p.regexToNFA("a*b*")
	fmt.Println(n1)

	d1 := p.regexToDFA("a*b*")
	fmt.Println(d1)

	d2 := d1.complement()
	fmt.Println(d2)

	d3 := d2.complement()
	fmt.Println(d3)
}

func TestDFAComplement3(t *testing.T) {
	var p parser

	n1 := p.regexToNFA("(a?b)")
	fmt.Println(n1)

	d1 := p.regexToDFA("(a?b)")
	fmt.Println(d1)

	d2 := d1.complement()
	fmt.Println(d2)
}
