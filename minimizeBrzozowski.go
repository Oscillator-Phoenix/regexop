package regexop

func (d *dfa) reverseToNFA() *nfa {
	alphbet := d.alphbet.copy()

	initial := newState()

	states := d.states.copy()
	states.insert(initial)

	finals := newStateSet()
	finals.insert(d.initial) // make the initial state of DFA a final state of NFA

	trans := newTransNFA()
	trans.insertSplit(initial, constEpsilon, d.finals.copy())
	for p, to := range d.trans.m {
		trans.insertSplitSingle(to, p.c, p.s) // reversing all the transDFA
	}

	return &nfa{alphbet, states, initial, finals, trans}
}

func (d *dfa) reverseToDFA() *dfa {
	return d.reverseToNFA().toDFA()
}

// minimize returns an equivalent DFA with minimized number of state
// Brzozowski's algorithm for DFA minimization
func (d *dfa) minimizeBrzozowski() *dfa {
	// Bug: this implementation CANNOT returns a minimized DFA !!!
	// Specifically, it CANNOT reduce the prefix redundancy states !!!
	return d.reverseToDFA().reverseToDFA()
}
