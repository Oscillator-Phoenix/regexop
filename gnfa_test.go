package regexop

import (
	"fmt"
	"testing"
)

func TestTransGNFA(t *testing.T) {
	tg := newTransGNFA()
	fmt.Println(tg)

	tg.setRegex(0, 1, []rune("a"))
	tg.setRegex(1, 2, []rune("b"))
	tg.setRegex(3, 4, []rune("b*"))

	fmt.Println(tg)

}

func TestGNFA1(t *testing.T) {
	var p parser

	d := p.regexToDFA("a|b")
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

}

func TestGNFAToRegex1(t *testing.T) {
	var p parser

	d := p.regexToDFA("a|b")
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	fmt.Println(g.toRegex())
}

func TestGNFAToRegex2(t *testing.T) {
	var p parser

	d := p.regexToDFA("a")
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	fmt.Println(g.toRegex())
}

func TestGNFAToRegex3(t *testing.T) {
	var p parser

	d := p.regexToDFA("ab")
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	fmt.Println(g.toRegex())
}

func TestGNFAToRegex4(t *testing.T) {
	var p parser

	d := p.regexToDFA("a|b")
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	fmt.Println(g.toRegex())
}

func TestGNFAToRegex5(t *testing.T) {
	var p parser

	d := p.regexToDFA("a|b|c")
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	fmt.Println(g.toRegex())
}

func TestGNFAToRegex6(t *testing.T) {
	var p parser

	d := p.regexToDFA("a*").minimize()
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	fmt.Println(g.toRegex())
}

func TestCase1(t *testing.T) {
	str := ""
	rs := []rune(str)
	fmt.Println("str", str)
	fmt.Println("rs", rs)
	fmt.Println("rs len", len(rs))
}

func TestGNFAToRegex7(t *testing.T) {
	var p parser

	d := p.regexToDFA("(a|b)*").minimize()
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	fmt.Println(g.toRegex())
}

func TestGNFAToRegex8(t *testing.T) {
	var p parser

	d := p.regexToDFA("(a|b)*abb").minimize()
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	fmt.Println(g.toRegex())
}

func TestGNFAToRegex9(t *testing.T) {
	var p parser

	d := p.regexToDFA("a(a|b)*b").minimize()
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	fmt.Println(g.toRegex())
}

func TestGNFAToRegex10(t *testing.T) {
	var p parser

	d := p.regexToDFA("a|b|c|d").minimize()
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	fmt.Println(prettyRegex(g.toRegex()))
}

func TestGNFAToRegex11(t *testing.T) {
	var p parser

	test := "a*|b?|c+|(d)"
	d := p.regexToDFA(test).minimize()
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	re := g.toRegex()
	fmt.Println(re)
	fmt.Println(IsEqualTo(re, test))
}

func TestGNFAToRegex12(t *testing.T) {
	var p parser

	test := "a*|bb|c+|(d)"
	d := p.regexToDFA(test).minimize()
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	re := g.toRegex()
	fmt.Println(re)
	fmt.Println(IsEqualTo(re, test))
}

func TestGNFAToRegex13(t *testing.T) {
	var p parser

	test := "(a|b)*abb"

	d := p.regexToDFA(test).minimize()
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	re := g.toRegex()
	fmt.Println(re)
	fmt.Println(IsEqualTo(re, test))
}

func TestGNFAToRegex14(t *testing.T) {
	var p parser

	test := "(bbb(a?)(b*))+"

	d := p.regexToDFA(test).minimize()
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	re := g.toRegex()
	fmt.Println(re)
	fmt.Println(IsEqualTo(re, test))
}

func TestGNFAToRegex15(t *testing.T) {
	var p parser

	test := "(a|b)*a(a|b)(a|b)(a|b)"

	d := p.regexToDFA(test).minimize()
	fmt.Println(d)

	g := d.toGNAF()
	fmt.Println(g)

	re := g.toRegex()
	fmt.Println(re)
	fmt.Println(len(re))
	// fmt.Println(IsEqualTo(re, test))  // this cost long time!!!
}
