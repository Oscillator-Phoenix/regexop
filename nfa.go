package regexop

import (
	"fmt"
	"strings"
)

type nfa struct {
	alphbet *symbolSet
	states  *stateSet
	initial state
	finals  *stateSet
	trans   *transNFA
}

func (n *nfa) toDFA() *dfa {

	alphbet := n.alphbet.copy()
	alphbet.erase(constEpsilon)

	states := newStateSet()

	finals := newStateSet()

	trans := newTransDFA()

	stateSet2StateContext := map[string]state{}
	stateSet2State := func(ss *stateSet) state {
		str := ss.String() // injective: stateSet --> string
		if _, isPresent := stateSet2StateContext[str]; !isPresent {
			stateSet2StateContext[str] = newState()
		}
		return stateSet2StateContext[str]
	}

	s0 := newStateSet()
	s0.insert(n.initial)
	ecs0 := n.epsilonClosure(s0)
	initial := stateSet2State(ecs0) // initial
	states.insert(initial)

	unMaskQueue := [](*stateSet){}
	unMaskQueue = append(unMaskQueue, ecs0) // push

	for len(unMaskQueue) > 0 {

		T := unMaskQueue[0] // front of queue
		_T := stateSet2State(T)
		unMaskQueue = unMaskQueue[1:] // pop front

		// construct finals
		if !intersectionStateSet(n.finals, T).empty() {
			finals.insert(_T)
		}

		for a := range alphbet.m {

			U := n.epsilonClosure(n.move(T, a))
			_U := stateSet2State(U)
			if U.size() == 0 {
				continue
			}

			if !states.find(_U) {
				states.insert(_U)
				unMaskQueue = append(unMaskQueue, U) // push
			}

			trans.insertSplit(_T, a, _U)
		}
	}

	// printing for debug
	// fmt.Println("stateSet to state context")
	// for str, _state := range stateSet2StateContext {
	// 	fmt.Println("epsilon closure ", str, " -> ", _state)
	// }

	return &dfa{alphbet, states, initial, finals, trans}
}

func (n *nfa) epsilonClosure(T *stateSet) *stateSet {

	stack := []state{}
	for s := range T.m {
		stack = append(stack, s) // push
	}

	closure := T.copy()

	for len(stack) > 0 {

		top := stack[len(stack)-1]   // top
		stack = stack[:len(stack)-1] // pop

		for p, tos := range n.trans.m {
			if p.s == top && p.c == constEpsilon {
				for _state := range tos.m {
					if !closure.find(_state) {
						closure.insert(_state)
						stack = append(stack, _state) // push
					}
				}
			}
		}
	}

	return closure
}

func (n *nfa) move(T *stateSet, a symbol) *stateSet {

	next := newStateSet()
	for p, tos := range n.trans.m {
		if T.find(p.s) && p.c == a {
			for _state := range tos.m {
				next.insert(_state)
			}
		}
	}

	return next
}

func (n *nfa) accept(input string) bool {
	rs := []rune(input)
	start := newStateSet()
	start.insert(n.initial)
	ss := n.epsilonClosure(start)

	for _, a := range rs {
		ss = n.epsilonClosure(n.move(ss, symbol(a)))
	}

	if !intersectionStateSet(ss, n.finals).empty() {
		return true
	}
	return false
}

func (n nfa) String() string {
	var b strings.Builder

	b.WriteString("state set:\n")
	b.WriteString(fmt.Sprint(n.states))
	b.WriteString("\n")

	b.WriteString("symbol set:\n")
	b.WriteString(fmt.Sprint(n.alphbet))
	b.WriteString("\n")

	b.WriteString("initial state:\n")
	b.WriteString(fmt.Sprint(n.initial))
	b.WriteString("\n")

	b.WriteString("final state set:\n")
	b.WriteString(fmt.Sprint(n.finals))
	b.WriteString("\n")

	b.WriteString("trans of NFA:\n")
	b.WriteString(fmt.Sprint(n.trans))
	b.WriteString("\n")

	return b.String()
}
