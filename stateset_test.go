package regexop

import (
	"fmt"
	"testing"
)

func TestStateSet(t *testing.T) {
	ss := newStateSet()
	fmt.Println(ss) // []
	ss.insert(1)
	ss.insert(2)
	fmt.Println(ss) // [1 2]
	ss.insert(3)
	ss.insert(4)
	fmt.Println(ss) // [1 2 3 4]
	ss.erase(1)
	ss.erase(3)
	ss.erase(5)
	fmt.Println(ss) // [2 4]
}

func TestUnionStateSet(t *testing.T) {
	ss1 := newStateSet()
	ss1.insert(1)
	ss1.insert(2)

	ss2 := newStateSet()
	ss2.insert(2)
	ss2.insert(3)

	u1 := unionStateSet(ss1, ss2)
	fmt.Println(u1)

	ss3 := newStateSet()
	ss3.insert(3)
	ss3.insert(4)

	u2 := unionStateSet(ss1, ss2, ss3)
	fmt.Println(u2)
}

func TestIntersectionStateSet(t *testing.T) {
	ss1 := newStateSet()
	ss1.insert(1)
	ss1.insert(2)
	ss1.insert(3)

	ss2 := newStateSet()
	ss2.insert(2)
	ss2.insert(3)
	ss2.insert(4)

	i1 := intersectionStateSet(ss1, ss2)
	fmt.Println(i1)

	ss3 := newStateSet()
	ss3.insert(3)
	ss3.insert(4)
	ss3.insert(5)

	i2 := intersectionStateSet(ss1, ss2, ss3)
	fmt.Println(i2)
}
