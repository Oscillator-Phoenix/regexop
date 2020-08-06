package regexop

var stateCount int = -1 // stateCount is a global state Increaser

func newState() state {
	stateCount++
	return state(stateCount)
}

type cartesianStatePair struct {
	x state
	y state
}

// cartesianContext keeps the cartesian map (state_1, state_2) --> state_3
type cartesianContext struct {
	cartesianMap map[cartesianStatePair]state
}

func newCartesianContext() *cartesianContext {
	var c cartesianContext
	c.cartesianMap = make(map[cartesianStatePair]state)
	return &c
}

// cartesianState is an idempotent function
func (c *cartesianContext) cartesianState(s1, s2 state) state {
	if _, isPresent := c.cartesianMap[cartesianStatePair{s1, s2}]; !isPresent {
		c.cartesianMap[cartesianStatePair{s1, s2}] = newState()
	}
	return c.cartesianMap[cartesianStatePair{s1, s2}]
}

func (c *cartesianContext) cartesianStateSet(ss1, ss2 *stateSet) *stateSet {
	ss := newStateSet()
	for x := range ss1.m {
		for y := range ss2.m {
			ss.insert(c.cartesianState(x, y))
		}
	}
	return ss
}
