package regexop

import (
	"fmt"
	"sort"
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

func (ss *stateSet) size() int {
	return len(ss.m)
}

func (ss *stateSet) empty() bool {
	return ss.size() == 0
}

func (ss *stateSet) insert(s ...state) {
	for _, _state := range s {
		ss.m[_state] = struct{}{}
	}
}

func (ss *stateSet) erase(s state) {
	delete(ss.m, s)
}

func (ss *stateSet) getOne() state {
	for _state := range ss.m {
		ret := _state
		return ret
	}
	panic("will not reach here: getOne() from empty state set")
}

// stateSlice returns sorted state slice
func (ss *stateSet) stateSlice() []state {
	is := ss.intSlice()
	ret := make([]state, len(is))
	for i := 0; i < len(is); i++ {
		ret[i] = state(is[i])
	}
	return ret
}

// intSlice returns sorted int slice
func (ss *stateSet) intSlice() []int {
	ret := make([]int, 0, len(ss.m))
	for s := range ss.m {
		ret = append(ret, int(s))
	}
	sort.Ints(ret) // sort
	return ret
}

func (ss stateSet) String() string {
	return fmt.Sprint(ss.intSlice())
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

func (ss *stateSet) difference(ss2 *stateSet) *stateSet {
	diff := ss.copy()
	for s := range ss2.m {
		diff.erase(s)
	}
	return diff
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
