package regexop

import (
	"fmt"
	"strings"
)

type transDFA struct {
	m map[transPair]state
}

func newTransDFA() *transDFA {
	var td transDFA
	td.m = make(map[transPair]state)
	return &td
}

func (td *transDFA) insert(p transPair, to state) {
	td.m[p] = to
}

func (td *transDFA) insertSplit(from state, a symbol, to state) {
	td.insert(transPair{from, a}, to)
}

func (td *transDFA) getSplit(from state, a symbol) state {
	if s, isPresent := td.m[transPair{from, a}]; isPresent {
		return s
	}
	return constDeadState
}

func (td transDFA) String() string {
	var b strings.Builder
	for p, to := range td.m {
		b.WriteString(fmt.Sprintf("(%d, %s) -> %d", p.s, string(p.c), to))
		b.WriteString("\n")
	}
	return b.String()
}
