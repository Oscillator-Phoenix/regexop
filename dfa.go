package regexop

// all todo

type dfa struct {
	alphbet *symbolSet
	states  *stateSet
	initial state
	finals  *stateSet
	trans   *transDFA
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
