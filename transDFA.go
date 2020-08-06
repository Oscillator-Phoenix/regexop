package regexop

type transDFA struct {
	m map[transPair]state
}

func newTransDFA() *transDFA {
	var td transDFA
	td.m = make(map[transPair]state)
	return &td
}

func (td *transDFA) insert(p transPair, s state) {
	td.m[p] = s
}
