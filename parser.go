package regexop

type parser struct {
}

func (p *parser) regexToDFA(regex string) *dfa {
	return nil
}

func (p *parser) regexToNFA(regex string) *nfa {
	return nil
}

func (p *parser) epsilonNFA() *nfa {
	return nil
}

func (p *parser) symbolNFA() *nfa {
	return nil
}

func (p *parser) unionNFA(n1, n2 *nfa) *nfa {
	return nil
}

func (p *parser) conNFA(n1, n2 *nfa) *nfa {
	return nil
}

func (p *parser) starNFA(n *nfa) *nfa {
	return nil
}

func (p *parser) plusNFA(n *nfa) *nfa {
	return nil
}

func (p *parser) queryNFA(n *nfa) *nfa {
	return nil
}
