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
		re := g.trans.getRegex(g.initial, g.final)
		return string(re)
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

			if r1 == nil || r3 == nil {
				continue
			}

			newRegex := ""

			if r1 != nil {
				newRegex += string(r1)
			}

			if r2 != nil {
				newRegex += ("(" + string(r2) + ")" + "*")
			}

			if r3 != nil {
				newRegex += string(r3)
			}

			if r4 != nil {
				if newRegex != "" {
					newRegex = "(" + newRegex + ")" + "|" + string(r4)
				} else {
					newRegex += string(r4)
				}
			}

			g.trans.setRegex(qi, qj, []rune(newRegex))
			fmt.Printf("setRegex (%d, %d) : %s\n", qi, qj, newRegex)
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
