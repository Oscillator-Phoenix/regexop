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
