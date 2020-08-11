package regexop

type waitItem struct {
	P *stateSet
	a symbol
}

func _theSmallestOf(ss1, ss2 *stateSet) *stateSet {
	if ss1.size() < ss2.size() {
		return ss1
	}
	return ss2
}

func isRefinedBy(d *dfa, R *stateSet, Pa waitItem) (isRf bool, Rp, Rpp *stateSet) {
	Rp = newStateSet()
	Rpp = newStateSet()

	for q := range R.m {
		if Pa.P.find(d.trans.getSplit(q, Pa.a)) {
			Rp.insert(q)
		} else {
			Rpp.insert(q)
		}
	}

	if Rp.size() == 0 || Rpp.size() == 0 {
		return false, nil, nil
	}

	return true, Rp, Rpp
}

func updateWait(wait []waitItem, B, Bp, Bpp *stateSet) {
	waitSize := len(wait)
	for i := 0; i < waitSize; i++ {
		if wait[i].P == B {
			wait[i].P = Bp                                // just overwrite
			wait = append(wait, waitItem{Bpp, wait[i].a}) // append
		}
	}
}

func (d *dfa) completeOmittedForMinimize() {
	for s := range d.states.m {
		for a := range d.alphbet.m {
			if d.trans.getSplit(s, a) == constDeadState {
				d.states.insert(constDeadState)
				d.trans.insertSplit(s, a, constDeadState)
			}
		}
	}
}

// minimize returns an equivalent DFA with minimized number of state
// Hopcroft’s algorithm for DFA minimization
func (d *dfa) minimizeHopcroft() *dfa {

	d.completeOmittedForMinimize()

	pi := map[*stateSet](struct{}){} // 2
	F := d.finals.copy()
	QF := d.states.difference(d.finals)
	pi[F] = struct{}{}
	pi[QF] = struct{}{}

	S := _theSmallestOf(F, QF) // 3
	wait := []waitItem{}       // 4

	for a := range d.alphbet.m { // 5
		wait = append(wait, waitItem{S, a}) // 6
	} // 7

	for len(wait) > 0 { // 8

		Sa := wait[0]   // 9
		wait = wait[1:] // 10

		piSlice := []*stateSet{}
		for B := range pi {
			piSlice = append(piSlice, B)
		}
		for _, B := range piSlice { // 11

			isRf, Bp, Bpp := isRefinedBy(d, B, Sa) // 12

			if !isRf {
				continue
			}

			delete(pi, B) // 13
			pi[Bp] = struct{}{}
			pi[Bpp] = struct{}{}

			// Update
			updateWait(wait, B, Bp, Bpp) // 14

			C := _theSmallestOf(Bp, Bpp) // 15
			for a := range d.alphbet.m { // 16
				wait = append(wait, waitItem{C, a}) // 17
			} // 18
		} //19
	} //20

	return constrcutMinimizeFrom(d, pi)
}

func constrcutMinimizeFrom(d *dfa, pi map[*stateSet](struct{})) *dfa {
	// fmt.Println("minimizeHopcroft partition")
	// for ss := range pi {
	// 	fmt.Println(ss)
	// }

	// construct minimized DFA by partition
	alphbet := d.alphbet.copy()
	states := newStateSet()
	var initial state
	finals := newStateSet()
	trans := newTransDFA()

	for ss := range pi {

		if ss.size() > 0 {

			representative := ss.stateSlice()[0]

			states.insert(representative)

			if ss.find(d.initial) {
				initial = representative
			}

			if !intersectionStateSet(ss, d.finals).empty() {
				finals.insert(representative)
			}

		}
	}

	// states.erase(constDeadState)

	for p, to := range d.trans.m {
		if states.find(p.s) {
			for ss := range pi {
				if ss.find(to) {
					representative := ss.stateSlice()[0]
					trans.insertSplit(p.s, p.c, representative)
				}
			}
		}
	}

	return &dfa{alphbet, states, initial, finals, trans}
}
