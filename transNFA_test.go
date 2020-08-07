package regexop

import (
	"fmt"
	"testing"
)

func TestTransNFA(t *testing.T) {
	tn := newTransNFA()

	tn.insertSplitSingle(0, 'a', 1)
	fmt.Println(tn)

	ss := newStateSet()
	ss.insert(1)
	ss.insert(2)
	ss.insert(3)
	tn.insertSplit(0, 'b', ss)
	fmt.Println(tn)

	ss2 := newStateSet()
	ss2.insert(3)
	ss2.insert(4)
	tn.unionInsert(transPair{0, 'b'}, ss2)
	fmt.Println(tn)
}

func TestTransNFACopy(t *testing.T) {
	tn := newTransNFA()

	tn.insertSplitSingle(0, 'a', 1)
	ss := newStateSet()
	ss.insert(1)
	ss.insert(2)
	ss.insert(3)
	tn.insertSplit(0, 'b', ss)

	fmt.Println(tn)

	tn2 := tn.copy()
	fmt.Println("copied: ", tn2)
}

func TestUnionTransNFA(t *testing.T) {
	tn := newTransNFA()
	tn.insertSplitSingle(0, 'a', 1, 99)
	ss := newStateSet()
	ss.insert(1)
	ss.insert(2)
	ss.insert(3)
	tn.insertSplit(0, 'b', ss)

	tn2 := newTransNFA()
	tn2.insertSplitSingle(0, 'a', 100)
	ss2 := newStateSet()
	ss2.insert(2)
	ss2.insert(300)
	ss2.insert(400)
	tn2.insertSplit(0, 'b', ss2)

	u := unionTwoTransNFA(tn, tn2)

	fmt.Println("tn1")
	fmt.Println(tn)
	fmt.Println("tn2")
	fmt.Println(tn2)
	fmt.Println("union")
	fmt.Println(u)
}

func TestUnionInsertSplit(t *testing.T) {
	tn := newTransNFA()
	tn.insertSplitSingle(0, 'a', 1, 99)
	ss := newStateSet()
	ss.insert(1)
	ss.insert(2)
	ss.insert(3)
	tn.insertSplit(0, 'b', ss)

	fmt.Println("tn1")
	fmt.Println(tn)

	tn.unionInsertSplit(0, 'a', 233, 666)
	tn.unionInsertSplit(0, 'b', 555)

	fmt.Println("tn2")
	fmt.Println(tn)

	tn2 := newTransNFA()
	tn2.unionInsertSplit(0, 'a', 233, 666)
	tn2.unionInsertSplit(0, 'b', 555)
	fmt.Println("tn3")
	fmt.Println(tn2)
}
