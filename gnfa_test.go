package regexop

import (
	"fmt"
	"testing"
)

func TestTransGNFA(t *testing.T) {
	tg := newTransGNFA()
	fmt.Println(tg)

	tg.setRegex(0, 1, "a")
	tg.setRegex(1, 2, "b")
	tg.setRegex(3, 4, "b*")

	fmt.Println(tg)

	if tg.getRegex(0, 1) != "a" {
		t.FailNow()
	}
	if tg.getRegex(1, 2) != "b" {
		t.FailNow()
	}
	if tg.getRegex(2, 3) != "" {
		t.FailNow()
	}
	if tg.getRegex(3, 4) != "b*" {
		t.FailNow()
	}
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
