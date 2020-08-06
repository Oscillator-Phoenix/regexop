package regexop

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
	return ""
}
