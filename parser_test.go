package regexop

import (
	"fmt"
	"testing"
)

func TestParserSymbolNFA(t *testing.T) {
	var p parser
	fmt.Println(p.symbolNFA('a'))
	fmt.Println(p.symbolNFA('b'))
	fmt.Println(p.epsilonNFA())
}

func TestParserUnionNFA(t *testing.T) {
	var p parser
	n1 := p.symbolNFA('a')
	n2 := p.symbolNFA('b')
	u := p.unionNFA(n1, n2)
	fmt.Println(u)
}

func TestParserConNFA(t *testing.T) {
	var p parser
	n1 := p.symbolNFA('a')
	n2 := p.symbolNFA('b')
	n3 := p.unionNFA(n1, n2)
	n4 := p.symbolNFA('c')
	n5 := p.conNFA(n4, n3)
	fmt.Println(n5)
}

func TestParserStarNFA1(t *testing.T) {
	var p parser
	n1 := p.symbolNFA('a')
	n2 := p.starNFA(n1)
	fmt.Println(n2)
}

func TestParserStarNFA2(t *testing.T) {
	var p parser
	n1 := p.symbolNFA('a')
	n2 := p.symbolNFA('b')
	n3 := p.unionNFA(n1, n2)
	n4 := p.symbolNFA('c')
	n5 := p.conNFA(n4, n3)
	n6 := p.starNFA(n5)
	fmt.Println(n6)
}

func TestParserPlusNFA1(t *testing.T) {
	var p parser
	n1 := p.symbolNFA('a')
	n2 := p.plusNFA(n1)
	fmt.Println(n2)
}

func TestParserPlusNFA2(t *testing.T) {
	var p parser
	n1 := p.symbolNFA('a')
	n2 := p.symbolNFA('b')
	n3 := p.unionNFA(n1, n2)
	n4 := p.symbolNFA('c')
	n5 := p.conNFA(n4, n3)
	n6 := p.starNFA(n5)
	n7 := p.plusNFA(n6)
	fmt.Println(n7)
}

func TestParserQueryNFA1(t *testing.T) {
	var p parser
	n1 := p.symbolNFA('a')
	n2 := p.queryNFA(n1)
	fmt.Println(n2)
}

func TestParserQueryNFA2(t *testing.T) {
	var p parser
	n1 := p.symbolNFA('a')
	n2 := p.symbolNFA('b')
	n3 := p.unionNFA(n1, n2)
	n4 := p.symbolNFA('c')
	n5 := p.conNFA(n4, n3)
	n6 := p.starNFA(n5)
	n7 := p.plusNFA(n6)
	n8 := p.queryNFA(n7)
	fmt.Println(n8)
}

func TestParserRegexToNFA_(t *testing.T) {
	var p parser
	p.regexToNFA("abc你我他()|*?+") // illegal regualr expression
}

func TestParserRegexToNFA0(t *testing.T) {
	var p parser

	n0 := p.regexToNFA("(aa|b)*abb")
	fmt.Println("n0")
	fmt.Println(n0)
}

func TestParserRegexToNFA1(t *testing.T) {
	var p parser

	n1 := p.regexToNFA("(a|b)*abb")
	fmt.Println("n1")
	fmt.Println(n1)
}

func TestParserRegexToNFA2(t *testing.T) {
	var p parser

	n2 := p.regexToNFA("a")
	fmt.Println("n2")
	fmt.Println(n2)
}
