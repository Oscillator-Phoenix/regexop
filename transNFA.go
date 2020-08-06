package regexop

type transPair struct {
	s state
	c symbol
}

type transNFA struct {
	m map[transPair](*stateSet)
}

func newTransNFA() *transNFA {
	var tn transNFA
	tn.m = make(map[transPair](*stateSet))
	return &tn
}

func (tn *transNFA) insert(p transPair, tos *stateSet) {
	tn.m[p] = tos
}

func (tn *transNFA) insertSingle(p transPair, to state) {
	if tn.m[p] == nil {
		tn.m[p] = newStateSet()
	}
	tn.m[p].insert(to)
}

func (tn *transNFA) insertSplit(from state, c symbol, tos *stateSet) {
	p := transPair{from, c}
	tn.insert(p, tos)
}

func (tn *transNFA) insertSplitSingle(from state, c symbol, to state) {
	p := transPair{from, c}
	tn.insertSingle(p, to)
}

func (tn *transNFA) copy() *transNFA {
	c := newTransNFA()

	for p, tos := range tn.m {
		c.insert(p, tos)
	}

	return c
}

func (tn *transNFA) find(p transPair) bool {
	_, find := tn.m[p]
	return find
}

func (tn *transNFA) unionInsert(p transPair, tos *stateSet) {
	if tn.m[p] == nil {
		tn.m[p] = newStateSet()
	}
	tn.m[p] = unionStateSet(tn.m[p], tos)
}

func unionTwoTransNFA(tn1, tn2 *transNFA) *transNFA {
	u := tn1.copy()

	for p, tos := range tn2.m {
		u.unionInsert(p, tos)
	}

	return u
}
