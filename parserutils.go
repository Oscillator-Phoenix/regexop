package regexop

type nfaStack struct {
	s []*nfa
}

type runeStack struct {
	s []rune
}

func newNFAStack() *nfaStack {
	var s nfaStack
	s.s = make([]*nfa, 0)
	return &s
}

func (s *nfaStack) empty() bool {
	return len(s.s) == 0
}

func (s *nfaStack) size() int {
	return len(s.s)
}

func (s *nfaStack) top() *nfa {
	if s.empty() {
		panic("empty stack")
	}
	return s.s[len(s.s)-1]
}

func (s *nfaStack) push(n *nfa) {
	s.s = append(s.s, n)
}

func (s *nfaStack) pop() {
	if s.empty() {
		panic("empty stack")
	}
	s.s = s.s[:(len(s.s) - 1)]
}

func newRuneStack() *runeStack {
	var s runeStack
	s.s = make([]rune, 0)
	return &s
}

func (s *runeStack) empty() bool {
	return len(s.s) == 0
}

func (s *runeStack) size() int {
	return len(s.s)
}

func (s *runeStack) top() rune {
	if s.empty() {
		panic("empty stack")
	}
	return s.s[len(s.s)-1]
}

func (s *runeStack) push(n rune) {
	s.s = append(s.s, n)
}

func (s *runeStack) pop() {
	if s.empty() {
		panic("empty stack")
	}
	s.s = s.s[:(len(s.s) - 1)]
}
