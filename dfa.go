package regexop

type dfa struct {
	alphbet *symbolSet
	states  *stateSet
	initial state
	finals  *stateSet
	trans   *transDFA
}

func intersection(d1, d2 *dfa) dfa {

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

	return dfa{alphbet, states, initial, finals, trans}
}
