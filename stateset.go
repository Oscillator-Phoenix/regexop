package regexop

import (
	"fmt"
)

type state int

type stateSet struct {
	m map[state]struct{}
}

func newStateSet() *stateSet {
	var ss stateSet
	ss.m = make(map[state]struct{})
	return &ss
}

func (ss *stateSet) insert(s state) {
	ss.m[s] = struct{}{}
}

func (ss *stateSet) erase(s state) {
	delete(ss.m, s)
}

func (ss *stateSet) stateSlice() []state {
	ret := make([]state, 0, len(ss.m))
	for s := range ss.m {
		ret = append(ret, s)
	}
	return ret
}

func (ss stateSet) String() string {
	return fmt.Sprint(ss.stateSlice())
}

func (ss *stateSet) find(s state) bool {
	_, find := ss.m[s]
	return find
}

func (ss *stateSet) copy() *stateSet {
	c := newStateSet()
	for s := range ss.m {
		c.insert(s)
	}
	return c
}

func unionStateSet(sss ...*stateSet) *stateSet {
	if len(sss) == 0 {
		return nil
	}

	union := newStateSet()
	for _, ss := range sss {
		for s := range ss.m {
			union.insert(s)
		}
	}
	return union
}

func intersectionStateSet(sss ...*stateSet) *stateSet {
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
