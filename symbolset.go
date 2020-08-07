package regexop

import "strings"

const (
	constEpsilon symbol = symbol(rune(0))
)

type symbol rune

type symbolSet struct {
	m map[symbol]struct{}
}

func newSymbolSet() *symbolSet {
	var ss symbolSet
	ss.m = make(map[symbol]struct{})
	return &ss
}

func newSymbolSetWithEpsilon() *symbolSet {
	ss := newSymbolSet()
	ss.insert(constEpsilon) // with constEpsilon
	return ss
}

func (ss *symbolSet) insert(s ...symbol) {
	for _, _symbol := range s {
		ss.m[_symbol] = struct{}{}
	}
}

func (ss *symbolSet) erase(s symbol) {
	delete(ss.m, s)
}

func (ss *symbolSet) symbolSlice() []symbol {
	ret := make([]symbol, 0, len(ss.m))
	for s := range ss.m {
		ret = append(ret, s)
	}
	return ret
}

func (ss symbolSet) String() string {
	var b strings.Builder

	b.WriteString("[ ")

	for s := range ss.m {
		if s == constEpsilon {
			b.WriteString("<esp>")
		} else {
			b.WriteRune(rune(s))
		}
		b.WriteString(" ")
	}

	b.WriteString("]")

	return b.String()
}

func (ss *symbolSet) find(s symbol) bool {
	_, find := ss.m[s]
	return find
}

func (ss *symbolSet) copy() *symbolSet {
	c := newSymbolSet()
	for s := range ss.m {
		c.insert(s)
	}
	return c
}

func unionSymbolSet(sss ...*symbolSet) *symbolSet {
	if len(sss) == 0 {
		return nil
	}

	union := newSymbolSet()
	for _, ss := range sss {
		for s := range ss.m {
			union.insert(s)
		}
	}
	return union
}

func intersectionSymbolSet(sss ...*symbolSet) *symbolSet {
	if len(sss) == 0 {
		return nil
	}

	intersection := sss[0].copy()

	for i := 1; i < len(sss); i++ {
		for s := range intersection.m {
			if !sss[i].find(s) {
				intersection.erase(s)
			}
		}
	}

	return intersection
}
