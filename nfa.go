package regexop

import (
	"fmt"
	"strings"
)

// all todo

type nfa struct {
	alphbet *symbolSet
	states  *stateSet
	initial state
	finals  *stateSet
	trans   *transNFA
}

func (n *nfa) toDFA() *dfa {
	return nil
}

func (n *nfa) epsilonClosure(states *stateSet) *stateSet {
	return nil
}

func (n *nfa) move(states *stateSet, a symbol) *stateSet {
	return nil
}

func (n nfa) accept(ss []symbol) bool {
	return false
}

func (n nfa) String() string {
	var b strings.Builder

	b.WriteString("state set:\n")
	b.WriteString(fmt.Sprint(n.states))
	b.WriteString("\n")

	b.WriteString("symbol set:\n")
	b.WriteString(fmt.Sprint(n.alphbet))
	b.WriteString("\n")

	b.WriteString("initial state:\n")
	b.WriteString(fmt.Sprint(n.initial))
	b.WriteString("\n")

	b.WriteString("final state set:\n")
	b.WriteString(fmt.Sprint(n.finals))
	b.WriteString("\n")

	b.WriteString("trans of NFA:\n")
	b.WriteString(fmt.Sprint(n.trans))
	b.WriteString("\n")

	return b.String()
}
