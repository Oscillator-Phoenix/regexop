package regexop

type nfa struct {
	alphbet *symbolSet
	states  *stateSet
	initial state
	finals  *stateSet
	trans   *transNFA
}
