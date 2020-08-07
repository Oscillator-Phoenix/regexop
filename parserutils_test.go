package regexop

import (
	"fmt"
	"testing"
)

func TestNFAStack(t *testing.T) {
	s := newNFAStack()
	n1 := &nfa{}
	n2 := &nfa{}
	n3 := &nfa{}

	s.push(n1)
	s.push(n2)
	fmt.Println(s)

	s.pop()
	s.push(n3)
	fmt.Println(s)

	s.pop()
	fmt.Println(s)

	s.pop()
	fmt.Println(s)

	if !s.empty() {
		t.Fail()
	}
}

func TestRuneStack(t *testing.T) {
	s := newRuneStack()
	n1 := 'a'
	n2 := 'b'
	n3 := 'c'

	s.push(n1)
	s.push(n2)
	fmt.Println(s)

	s.pop()
	s.push(n3)
	fmt.Println(s)

	s.pop()
	fmt.Println(s)

	s.pop()
	fmt.Println(s)

	if !s.empty() {
		t.Fail()
	}
}
