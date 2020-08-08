package regexop

import (
	"fmt"
	"strings"
)

type dfa struct {
	alphbet *symbolSet
	states  *stateSet
	initial state
	finals  *stateSet
	trans   *transDFA
}

func (d *dfa) toRegex() string {
	// todo
	return ""
}

//intersection retruns a DFA which accpets intersection of d1 and d2
func intersection(d1, d2 *dfa) *dfa {

	ctx := newCartesianContext()

	alphbet := intersectionSymbolSet(d1.alphbet, d2.alphbet)
	states := ctx.cartesianStateSet(d1.states, d2.states)
	initial := ctx.cartesianState(d1.initial, d2.initial)
	finals := ctx.cartesianStateSet(d1.finals, d2.finals)
	trans := newTransDFA()

	for p1, to1 := range d1.trans.m {
		for p2, to2 := range d2.trans.m {
			if p1.c == p2.c {
				p := transPair{
					s: ctx.cartesianState(p1.s, p1.s),
					c: p1.c,
				}
				to := ctx.cartesianState(to1, to2)
				trans.insert(p, to)
			}
		}
	}

	return &dfa{alphbet, states, initial, finals, trans}
}

// difference retruns a DFA which accpets (d1 - d2)
func difference(d1, d2 *dfa) *dfa {
	return nil
}

// complement returns the complement dfa on closure of symbol set
func (d *dfa) complement() *dfa {
	return nil
}

// isEmpty returns weather the language the dfa accpeted is empty
func (d *dfa) isEmpty() bool {
	return false
}

func (d *dfa) accept(input string) bool {
	rs := []rune(input)
	s := d.initial

	for _, a := range rs {
		s = d.trans.getSplit(s, symbol(a))
	}

	if d.finals.find(s) {
		return true
	}
	return false
}

// minimize returns an equivalent DFA with minimized number of state
func (d *dfa) minimize() *dfa {
	// todo
	return nil
}

func (d dfa) String() string {
	var b strings.Builder

	b.WriteString("state set:\n")
	b.WriteString(fmt.Sprint(d.states))
	b.WriteString("\n")

	b.WriteString("symbol set:\n")
	b.WriteString(fmt.Sprint(d.alphbet))
	b.WriteString("\n")

	b.WriteString("initial state:\n")
	b.WriteString(fmt.Sprint(d.initial))
	b.WriteString("\n")

	b.WriteString("final state set:\n")
	b.WriteString(fmt.Sprint(d.finals))
	b.WriteString("\n")

	b.WriteString("trans of DFA:\n")
	b.WriteString(fmt.Sprint(d.trans))
	b.WriteString("\n")

	return b.String()
}
