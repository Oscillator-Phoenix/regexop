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

func TestParserRegexToNFA(t *testing.T) {
	var p parser
	p.regexToNFA("abc你我他()|*?+")
}
