package regexop

import (
	"fmt"
	"strings"
)

type transGNFAPair struct {
	from state
	to   state
}

type transGNFA struct {
	// (from, to) -> regex
	m map[transGNFAPair]([]rune)
}

func newTransGNFA() *transGNFA {
	var tg transGNFA
	tg.m = make(map[transGNFAPair]([]rune))
	return &tg
}

func (tg *transGNFA) getRegex(from, to state) []rune {
	p := transGNFAPair{from, to}
	if re, isPresent := tg.m[p]; isPresent {
		return re
	}
	return nil
}

func (tg *transGNFA) setRegex(from, to state, regex []rune) {
	if regex == nil {
		panic("error at transGNFA.setRegex(): wrong regex")
	}
	p := transGNFAPair{from, to}
	tg.m[p] = regex
}

func (tg transGNFA) String() string {
	var b strings.Builder
	for p, re := range tg.m {
		if len(re) == 0 {
			b.WriteString(fmt.Sprintf("(%d, <esp>) -> %d\n", p.from, p.to))
		} else {
			b.WriteString(fmt.Sprintf("(%d, %s) -> %d\n", p.from, string(re), p.to))
		}
	}
	return b.String()
}

type gnfa struct {
	alphbet *symbolSet
	states  *stateSet
	initial state
	final   state
	trans   *transGNFA
	// help data for converting to regex
	di  *stateSet // states - {initial}
	df  *stateSet // states - {final}
	dif *stateSet // states - ( {initial} `union` {final} )
}

func newGFA(alphbet *symbolSet, states *stateSet, initial state, final state, trans *transGNFA) *gnfa {
	var g gnfa

	g.alphbet = alphbet
	g.states = states
	g.initial = initial
	g.final = final
	g.trans = trans

	g.di = g.states.copy()
	g.di.erase(g.initial)

	g.df = g.states.copy()
	g.df.erase(g.final)

	g.dif = g.states.copy()
	g.dif.erase(g.initial)
	g.dif.erase(g.final)

	return &g
}

func (g *gnfa) toRegex() string {

	if g.dif.size() == 0 {
		re := g.trans.getRegex(g.initial, g.final)
		return string(re)
	}

	// pop one
	rip := g.dif.getOne()
	g.states.erase(rip)
	g.di.erase(rip)
	g.df.erase(rip)
	g.dif.erase(rip)
	// fmt.Println("pop", rip)

	for qi := range g.df.m {
		for qj := range g.di.m {

			r1 := g.trans.getRegex(qi, rip)
			r2 := g.trans.getRegex(rip, rip)
			r3 := g.trans.getRegex(rip, qj)
			r4 := g.trans.getRegex(qi, qj)

			if r1 == nil || r3 == nil {
				continue
			}

			// newRegex := ""
			var newRegex strings.Builder

			if r1 != nil {
				if len(r1) > 0 { // is not epsilon
					if len(r1) == 1 {
						// newRegex += string(r1)
						newRegex.WriteString(string(r1))
					} else {
						// newRegex += "(" + string(r1) + ")"
						newRegex.WriteString("(")
						newRegex.WriteString(string(r1))
						newRegex.WriteString(")")
					}
				}
			}

			if r2 != nil {
				if len(r2) > 0 { // is not epsilon
					if len(r2) == 1 {
						// newRegex += (string(r2) + "*")
						newRegex.WriteString(string(r2))
						newRegex.WriteString("*")
					} else {
						// newRegex += ("(" + string(r2) + ")" + "*")
						newRegex.WriteString("(")
						newRegex.WriteString(string(r2))
						newRegex.WriteString(")")
						newRegex.WriteString("*")
					}
				}
			}

			if r3 != nil {
				if len(r3) > 0 {
					if len(r3) == 1 {
						// newRegex += string(r3)
						newRegex.WriteString(string(r3))
					} else {
						// newRegex += ("(" + string(r3) + ")")
						newRegex.WriteString("(")
						newRegex.WriteString(string(r3))
						newRegex.WriteString(")")
					}
				}
			}

			// if newRegex != "" {
			// 	newRegex = "(" + newRegex + ")"
			// }
			if newRegex.Len() > 0 {
				tmpStr := newRegex.String()
				newRegex.Reset()
				newRegex.WriteString("(")
				newRegex.WriteString(tmpStr)
				newRegex.WriteString(")")
				// fmt.Println("tmpStr", tmpStr)
				// fmt.Println("newRegex", newRegex.String())
			}

			if r4 != nil {
				if len(r4) == 0 { // is epsilon
					// newRegex = newRegex + "?"
					newRegex.WriteString("?")
				} else {
					if len(r4) == 1 {
						// newRegex = newRegex + "|" + string(r4)
						newRegex.WriteString("|")
						newRegex.WriteString(string(r4))
					} else {
						// newRegex = newRegex + "|" + "(" + string(r4) + ")"
						newRegex.WriteString("|")
						newRegex.WriteString("(")
						newRegex.WriteString(string(r4))
						newRegex.WriteString(")")
					}
				}
			}

			// g.trans.setRegex(qi, qj, []rune(newRegex))
			g.trans.setRegex(qi, qj, []rune(newRegex.String()))
			// fmt.Printf("setRegex (%d, %d) : %s\n", qi, qj, newRegex)
		}
	}

	return g.toRegex() // recursive call after reducing a state
}

func (d *dfa) toGNAF() *gnfa {
	alphbet := d.alphbet.copy()

	initial := newState()

	final := newState()

	states := d.states.copy()
	states.insert(initial, final)

	trans := newTransGNFA()

	// []rune{} represents the edge with symbol `epsilon`
	// nil represents no edge

	trans.setRegex(initial, d.initial, []rune{})

	for df := range d.finals.m {
		trans.setRegex(df, final, []rune{})
	}

	for p, to := range d.trans.m {
		// 使用 | 操作合并两端相同地箭头， 使得一个状态到另一个状态只有一个箭头
		from := p.s
		re := trans.getRegex(from, to)
		if re == nil {
			trans.setRegex(from, to, []rune{rune(p.c)})
		} else {
			re = append(re, rune('|'))
			re = append(re, rune(p.c))
			trans.setRegex(from, to, re)
		}
	}

	return newGFA(alphbet, states, initial, final, trans)
}

func (g gnfa) String() string {
	var b strings.Builder

	b.WriteString("state set:\n")
	b.WriteString(fmt.Sprint(g.states))
	b.WriteString("\n")

	b.WriteString("symbol set:\n")
	b.WriteString(fmt.Sprint(g.alphbet))
	b.WriteString("\n")

	b.WriteString("initial state:\n")
	b.WriteString(fmt.Sprint(g.initial))
	b.WriteString("\n")

	b.WriteString("final state:\n")
	b.WriteString(fmt.Sprint(g.final))
	b.WriteString("\n")

	b.WriteString("trans of GNFA:\n")
	b.WriteString(fmt.Sprint(g.trans))
	b.WriteString("\n")

	return b.String()
}
