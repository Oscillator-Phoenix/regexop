package regexop

import "fmt"

var _regeularSymbolSet *symbolSet = nil
var _lookBackNewNFASymbolSet *symbolSet = nil

func regeularSymbolSet() *symbolSet {

	if _regeularSymbolSet == nil {
		// _regeularSymbolSet init
		_regeularSymbolSet = newSymbolSet()
		_regeularSymbolSet.insert(symbol(rune('*')))
		_regeularSymbolSet.insert(symbol(rune('+')))
		_regeularSymbolSet.insert(symbol(rune('?')))
		_regeularSymbolSet.insert(symbol(rune('|')))
		_regeularSymbolSet.insert(symbol(rune('(')))
		_regeularSymbolSet.insert(symbol(rune(')')))
		// fmt.Println("regular symbol:", _regeularSymbolSet)
	}
	return _regeularSymbolSet
}

func isRegularSymbol(r rune) bool {
	return regeularSymbolSet().find(symbol(r))
}

func lookBackNewNFASymbolSet() *symbolSet {
	if _lookBackNewNFASymbolSet == nil {
		_lookBackNewNFASymbolSet = newSymbolSet()
		_lookBackNewNFASymbolSet.insert(constEpsilon)
		_lookBackNewNFASymbolSet.insert('*')
		_lookBackNewNFASymbolSet.insert('+')
		_lookBackNewNFASymbolSet.insert('?')
		_lookBackNewNFASymbolSet.insert(')')
		// fmt.Println("_lookBackNewNFASymbolSet:", _lookBackNewNFASymbolSet)
	}
	return _lookBackNewNFASymbolSet
}

func isLookBackNewNFASymbols(r rune) bool {
	return !isRegularSymbol(r) || lookBackNewNFASymbolSet().find(symbol(r))
}

type parser struct {
	// empty
}

func (p *parser) regexToDFA(regex string) *dfa {
	return p.regexToNFA(regex).toDFA()
}

// regexToNFA transforms a regular expression to an equivalent NFA
// the priority of regular operation: Star == Plus == Query > Con > Union
func (p *parser) regexToNFA(regex string) *nfa {
	nfaStack := newNFAStack()
	opStack := newRuneStack()

	re := []rune(regex)

	// for i := 0; i < len(re); i++ {
	// 	fmt.Println(string(re[i]), isRegularSymbol(re[i]), isLookBackNewNFASymbols(re[i]))
	// }

	conNFAOnStack := func() {
		nfa1 := nfaStack.top()
		nfaStack.pop()
		nfa2 := nfaStack.top()
		nfaStack.pop()
		nfaStack.push(p.conNFA(nfa2, nfa1))
		// fmt.Println("con")
	}

	unionNFAOnStack := func() {
		nfa1 := nfaStack.top()
		nfaStack.pop()
		nfa2 := nfaStack.top()
		nfaStack.pop()
		nfaStack.push(p.unionNFA(nfa2, nfa1))
		// fmt.Println("union")
	}

	for i := 0; i < len(re); i++ {
		// fmt.Println("No ", i, " re[i]= ", string(re[i]), " opStack: ", opStack, " nfaStckSize: ", nfaStack.size())

		if !isRegularSymbol(re[i]) {

			if i > 0 && isLookBackNewNFASymbols(re[i-1]) {
				if !opStack.empty() && opStack.top() == '#' {
					opStack.pop() // pop '#'
					conNFAOnStack()
				}
				opStack.push('#') // push '#'
			}

			nfaStack.push(p.symbolNFA(symbol(re[i])))
			// fmt.Println("symbol", string(re[i]))

		} else if re[i] == '|' { // union operation has the lowest priority

			for !opStack.empty() && opStack.top() != '(' {
				if opStack.top() == '|' {
					opStack.pop() // pop '|'
					unionNFAOnStack()
				} else if opStack.top() == '#' {
					opStack.pop() // pop '#'
					conNFAOnStack()
				}
			}

			opStack.push('|')

		} else if re[i] == '*' {

			nfa1 := nfaStack.top()
			nfaStack.pop()
			nfaStack.push(p.starNFA(nfa1))
			// fmt.Println("star")

		} else if re[i] == '+' {

			nfa1 := nfaStack.top()
			nfaStack.pop()
			nfaStack.push(p.plusNFA(nfa1))
			// fmt.Println("plus")

		} else if re[i] == '?' {

			nfa1 := nfaStack.top()
			nfaStack.pop()
			nfaStack.push(p.queryNFA(nfa1))
			// fmt.Println("query")

		} else if re[i] == '(' {

			if i > 0 && isLookBackNewNFASymbols(re[i-1]) {
				if !opStack.empty() && opStack.top() == '#' {
					opStack.pop() // pop '#'
					conNFAOnStack()
				}
				opStack.push('#')
			}

			opStack.push('(')

		} else if re[i] == ')' {

			for opStack.top() != '(' {
				if opStack.top() == '|' {
					opStack.pop() // pop '|'
					unionNFAOnStack()
				} else if opStack.top() == '#' {
					opStack.pop() // pop '#'
					conNFAOnStack()
				}
			}

			opStack.pop() // pop '('

		} else {
			panic("parse error") // todo: prettify
		}

	}

	for !opStack.empty() {

		if opStack.top() == '|' {
			opStack.pop() // pop '|'
			unionNFAOnStack()
		} else if opStack.top() == '#' {
			opStack.pop() // pop '#'
			conNFAOnStack()
		}

	}

	if nfaStack.size() != 1 {
		fmt.Println("nfaStack size: ", nfaStack.size())
		panic("parse error") // todo: prettify
	}

	return nfaStack.top()
}

func (p *parser) epsilonNFA() *nfa {
	return p.symbolNFA(constEpsilon)
}

func (p *parser) symbolNFA(s symbol) *nfa {
	var states *stateSet
	var alphbet *symbolSet
	var initial state
	var finals *stateSet
	var trans *transNFA

	s0 := newState()
	s1 := newState()

	states = newStateSet()
	states.insert(s0, s1)

	alphbet = newSymbolSet() // without epsilon
	alphbet.insert(s)

	initial = s0

	finals = newStateSet()
	finals.insert(s1)

	trans = newTransNFA()
	trans.insertSplitSingle(s0, s, s1)

	return &nfa{alphbet, states, initial, finals, trans}
}

func (p *parser) unionNFA(n1, n2 *nfa) *nfa {
	var states *stateSet
	var alphbet *symbolSet
	var initial state
	var finals *stateSet
	var trans *transNFA

	s0 := newState()
	s1 := newState()

	states = unionStateSet(n1.states, n2.states)
	states.insert(s0, s1)

	alphbet = unionSymbolSet(n1.alphbet, n2.alphbet)

	initial = s0

	finals = newStateSet()
	finals.insert(s1)

	trans = unionTwoTransNFA(n1.trans, n2.trans)
	trans.insertSplitSingle(s0, constEpsilon, n1.initial, n2.initial)
	tmpFinalsSlice := unionStateSet(n1.finals, n2.finals).stateSlice()
	for _, f := range tmpFinalsSlice {
		trans.insertSplitSingle(f, constEpsilon, s1)
	}

	return &nfa{alphbet, states, initial, finals, trans}
}

// conNFA returns a new NFA accpeting "n1#n2" where operation "#" is concatenation
func (p *parser) conNFA(n1, n2 *nfa) *nfa {
	var states *stateSet
	var alphbet *symbolSet
	var initial state
	var finals *stateSet
	var trans *transNFA

	states = unionStateSet(n1.states, n2.states)

	alphbet = unionSymbolSet(n1.alphbet, n2.alphbet)

	initial = n1.initial

	finals = n2.finals.copy()

	trans = unionTwoTransNFA(n1.trans, n2.trans)
	tmpFinalsSlice := n1.finals.stateSlice()
	for _, f := range tmpFinalsSlice {
		trans.unionInsertSplit(f, constEpsilon, n2.initial)
	}

	return &nfa{alphbet, states, initial, finals, trans}
}

func (p *parser) starNFA(n *nfa) *nfa {
	var states *stateSet
	var alphbet *symbolSet
	var initial state
	var finals *stateSet
	var trans *transNFA

	s0 := newState()
	s1 := newState()

	states = n.states.copy()
	states.insert(s0, s1)

	alphbet = n.alphbet.copy()

	initial = s0

	finals = newStateSet()
	finals.insert(s1)

	trans = n.trans.copy()
	trans.unionInsertSplit(s0, constEpsilon, n.initial, s1)
	tmpFinalsSlice := n.finals.stateSlice()
	for _, f := range tmpFinalsSlice {
		trans.unionInsertSplit(f, constEpsilon, n.initial, s1)
	}

	return &nfa{alphbet, states, initial, finals, trans}
}

func (p *parser) plusNFA(n *nfa) *nfa {
	var states *stateSet
	var alphbet *symbolSet
	var initial state
	var finals *stateSet
	var trans *transNFA

	s0 := newState()
	s1 := newState()

	states = n.states.copy()
	states.insert(s0, s1)

	alphbet = n.alphbet.copy()

	initial = s0

	finals = newStateSet()
	finals.insert(s1)

	trans = n.trans.copy()
	trans.unionInsertSplit(s0, constEpsilon, n.initial) // just initial state of nfa
	tmpFinalsSlice := n.finals.stateSlice()
	for _, f := range tmpFinalsSlice {
		trans.unionInsertSplit(f, constEpsilon, n.initial, s1)
	}

	return &nfa{alphbet, states, initial, finals, trans}
}

func (p *parser) queryNFA(n *nfa) *nfa {
	var states *stateSet
	var alphbet *symbolSet
	var initial state
	var finals *stateSet
	var trans *transNFA

	s0 := newState()
	s1 := newState()

	states = n.states.copy()
	states.insert(s0, s1)

	alphbet = n.alphbet.copy()

	initial = s0

	finals = newStateSet()
	finals.insert(s1)

	trans = n.trans.copy()
	trans.unionInsertSplit(s0, constEpsilon, n.initial, s1)
	tmpFinalsSlice := n.finals.stateSlice()
	for _, f := range tmpFinalsSlice {
		trans.unionInsertSplit(f, constEpsilon, s1) // just s1
	}

	return &nfa{alphbet, states, initial, finals, trans}
}
