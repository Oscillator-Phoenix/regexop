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

// toRegex transforms a DFA to an equivalent regular expression
// this implementation is from the book `Introduction to the Theory of Computation` Chapter 1.3.2
func (d *dfa) toRegex() string {
	re := d.minimize().toGNAF().toRegex()
	return prettyRegex(re)
}

// prettyRegex aims at reducing outermost pair of brackets
func prettyRegex(re string) string {
	// todo
	return re
}

// There are two usual method to minimize a DFA.
// One is called Brzozowski algorithm, the other is called Hopcroft algorithm.
func (d *dfa) minimize() *dfa {
	return d.minimizeHopcroft()
	// return  d.minimizeBrzozowski()  Bug!!!
}

// isSubset return weather d is the subset of d2
// `d1 is the subset of d2` is equal to (d1 - d2) is a empty set
func (d *dfa) isSubsetOf(d2 *dfa) bool {
	return d.difference(d2).isEmpty()
}

//intersection retruns a DFA which accpets intersection of d1 and d2
func intersectionTwoDFA(d1, d2 *dfa) *dfa {

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
					s: ctx.cartesianState(p1.s, p2.s),
					c: p1.c,
				}
				to := ctx.cartesianState(to1, to2)
				trans.insert(p, to)
			}
		}
	}

	// // printing for debug
	// fmt.Println("intersection two DFA Cartesian Context")
	// for p, s := range ctx.cartesianMap {
	// 	fmt.Printf("(%d, %d) <-> %d\n", p.x, p.y, s)
	// }

	return &dfa{alphbet, states, initial, finals, trans}
}

// difference retruns a DFA which accpets (d - d2)
func (d *dfa) difference(d2 *dfa) *dfa {
	// d1 - d2 <=> (d1) intersection (complement d2)

	// you should make sure:
	// the the complement operation on d2 will not lose that omitted trans information in d2,
	// which is important information in the intersection dfa of d1 and -d2.
	// So, you should enlarge the alphbet of d2.
	// In face, add irrelevant symbols about trans to alphbet will not effect the correctness of a DFA.
	// However, all symbols in alphbet effect the constructing the complement of DFA.
	d2.alphbet = unionSymbolSet(d.alphbet, d2.alphbet)

	return intersectionTwoDFA(d, d2.complement())
}

// complement returns the complement dfa on closure of symbol set
func (d *dfa) complement() *dfa {

	alphbet := d.alphbet.copy()

	states := d.states.copy()
	deadState := newState()
	states.insert(deadState)

	initial := d.initial

	finals := newStateSet()
	for s := range states.m {
		if !d.finals.find(s) {
			finals.insert(s)
		}
	}

	trans := newTransDFA()
	// initialize trans as dead trans
	for from := range states.m {
		for a := range alphbet.m {
			trans.insertSplit(from, a, deadState)
		}
	}
	// cover the dead trans
	for p, to := range d.trans.m {
		trans.insert(p, to)
	}

	return &dfa{alphbet, states, initial, finals, trans}
}

// isEmpty returns weather the language the dfa accpeted is empty
func (d *dfa) isEmpty() bool {
	passed := newStateSet() // set of passed state
	q := []state{}          // queue of state to pass

	q = append(q, d.initial) // push

	for len(q) > 0 {

		front := q[0]
		q = q[1:] // pop
		passed.insert(front)

		if d.finals.find(front) {
			return false //  path to finals, not empty
		}

		for a := range d.alphbet.m {
			if d.trans.findSplit(front, a) {
				to := d.trans.getSplit(front, a)
				if !passed.find(to) {
					q = append(q, to) // push
				}
			}
		}
	}

	return true // is empty
}

func unionTwoDFA(d1, d2 *dfa) *dfa {
	// the reason for enlarge alphbet is the same as mentioned in the comments of function `dfa.difference()`
	u := unionSymbolSet(d1.alphbet, d2.alphbet)
	d1.alphbet = u
	d2.alphbet = u
	return intersectionTwoDFA(d1.complement(), d2.complement()).complement()
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
