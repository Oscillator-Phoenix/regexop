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
	m map[transGNFAPair]string
}

func newTransGNFA() *transGNFA {
	var tg transGNFA
	tg.m = make(map[transGNFAPair]string)
	return &tg
}

func (tg *transGNFA) getRegex(from, to state) string {
	p := transGNFAPair{from, to}
	if re, isPresent := tg.m[p]; isPresent {
		return re
	}
	return ""
}

func (tg *transGNFA) setRegex(from, to state, regex string) {
	p := transGNFAPair{from, to}
	tg.m[p] = regex

}

func (tg transGNFA) String() string {
	var b strings.Builder
	for p, re := range tg.m {
		if re == string(constEpsilon) {
			b.WriteString(fmt.Sprintf("(%d, <esp>) -> %d\n", p.from, p.to))
		} else {
			b.WriteString(fmt.Sprintf("(%d, %s) -> %d\n", p.from, re, p.to))
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
	df  *stateSet // states - finals
	dif *stateSet // states - ( {initial} `union` finals )
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
		return g.trans.getRegex(g.initial, g.final)
	}

	// pop one
	rip := g.dif.stateSlice()[0]
	g.states.erase(rip)
	g.di.erase(rip)
	g.df.erase(rip)
	g.dif.erase(rip)
	fmt.Println("pop", rip)

	for qi := range g.df.m {
		for qj := range g.di.m {

			if qj == qi {
				continue
			}

			r1 := g.trans.getRegex(qi, rip)
			r2 := g.trans.getRegex(rip, rip)
			r3 := g.trans.getRegex(rip, qj)
			r4 := g.trans.getRegex(qi, qj)

			if r1 == "" || r3 == "" {
				continue
			}

			newRegex := ""

			if r1 != "" {
				newRegex += r1
			}

			if r2 != "" {
				newRegex += ("(" + r2 + ")" + "*")
			}

			if r3 != "" {
				newRegex += r3
			}

			if r4 != "" {
				if newRegex != "" {
					newRegex = "(" + newRegex + ")" + "|" + r4
				} else {
					newRegex += r4
				}
			}

			if newRegex != "" {
				g.trans.setRegex(qi, qj, newRegex)
				fmt.Printf("setRegex (%d, %d) : %s\n", qi, qj, newRegex)
			}
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
	trans.setRegex(initial, d.initial, "<e>")
	for df := range d.finals.m {
		trans.setRegex(df, final, "<e>")
	}
	for p, to := range d.trans.m {
		trans.setRegex(p.s, to, string(p.c))
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
