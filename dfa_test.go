package regexop

import (
	"fmt"
	"testing"
)

func TestDFAIntersection1(t *testing.T) {
	var p parser
	d1 := p.regexToDFA("a|b")
	d2 := p.regexToDFA("a")
	d3 := intersectionTwoDFA(d1, d2)
	fmt.Println(d3)
}

func TestDFAIntersection2(t *testing.T) {
	var p parser
	d1 := p.regexToDFA("a|b")
	d2 := p.regexToDFA("a|b")

	fmt.Println(d1)
	fmt.Println(d2)

	d3 := intersectionTwoDFA(d1, d2)

	fmt.Println(d3)
}

func TestDFAIntersection3(t *testing.T) {
	var p parser
	d1 := p.regexToDFA("(a|b)*")
	d2 := p.regexToDFA("a(a|b)*")

	fmt.Println(d1)
	fmt.Println(d2)

	d3 := intersectionTwoDFA(d1, d2)

	fmt.Println(d3)
}

func TestDFAComplement1(t *testing.T) {
	var p parser

	d1 := p.regexToDFA("(a|b)*")
	fmt.Println(d1)

	d2 := d1.complement()
	fmt.Println(d2)

	d3 := d2.complement()
	fmt.Println(d3)
}

func TestDFAComplement2(t *testing.T) {
	var p parser

	n1 := p.regexToNFA("a*b*")
	fmt.Println(n1)

	d1 := p.regexToDFA("a*b*")
	fmt.Println(d1)

	d2 := d1.complement()
	fmt.Println(d2)

	d3 := d2.complement()
	fmt.Println(d3)
}

func TestDFAComplement3(t *testing.T) {
	var p parser

	n1 := p.regexToNFA("(a?b)")
	fmt.Println(n1)

	d1 := p.regexToDFA("(a?b)")
	fmt.Println(d1)

	d2 := d1.complement()
	fmt.Println(d2)
}

func TestDFAIsEmpty(t *testing.T) {
	var p parser

	if p.regexToDFA("a").isEmpty() == true {
		t.Fail()
	}
	if p.regexToDFA("a").complement().isEmpty() == true {
		t.Fail()
	}
	if p.regexToDFA("a*b*").isEmpty() == true {
		t.Fail()
	}
	if p.regexToDFA("(a|b)*").isEmpty() == true {
		t.Fail()
	}
	if p.regexToDFA("(a|b)*").complement().isEmpty() == false {
		t.Fail()
	}
}

func TestDFAIsSubset(t *testing.T) {
	type strPair struct {
		s1 string
		s2 string
	}

	equal := []strPair{
		strPair{"a", "a"},
		strPair{"a+", "aa*"},
		strPair{"a|b", "b|a"},
		strPair{"a+|b+", "aa*|bb*"},
		strPair{"(a|b)(a|b)", "aa|ab|ba|bb"},
		strPair{"a+", "a*a"},
		strPair{"a+(aa)+", "(aa)+a+"},
		strPair{"b*a*b?a*", "b*a*ba*|b*a*"},
		strPair{"abcdefg", "abcdefg"},
	}

	less := []strPair{
		strPair{"a", "a|x"},
		strPair{"a+", "aa*|x"},
		strPair{"a|b", "b|a|x"},
		strPair{"a+|b+", "aa*|bb*|x"},
		strPair{"(a|b)(a|b)", "aa|ab|ba|bb|x"},
		strPair{"a+", "a*a|x"},
		strPair{"a+(aa)+", "(aa)+a+|x"},
	}

	greater := []strPair{
		strPair{"x|a", "a"},
		strPair{"x|a+", "aa*"},
		strPair{"x|a|b", "b|a"},
		strPair{"x|a+|b+", "aa*|bb*"},
		strPair{"x|(a|b)(a|b)", "aa|ab|ba|bb"},
		strPair{"x|a+", "a*a"},
		strPair{"x|a+(aa)+", "(aa)+a+"},
	}

	noseq := []strPair{
		strPair{"a", "b"},
		strPair{"aa", "ab"},
		strPair{"aa", "ba"},
		strPair{"aa", "bb"},
		strPair{"a*", "b*"},
		strPair{"a*b", "ab*"},
		strPair{"a+b*", "b+a*"},
		strPair{"xa", "ya"},
		strPair{"xa+", "yaa*"},
		strPair{"xa|b", "yb|a"},
		strPair{"xa+|b+", "yaa*|bb*"},
		strPair{"x(a|b)(a|b)", "yaa|ab|ba|bb"},
		strPair{"xa+", "ya*a"},
		strPair{"xa+(aa)+", "y(aa)+a+"},
	}

	testCase := func(s string, tests []strPair, answer1, answer2 bool) {
		fmt.Printf("%s case tesging\n", s)
		var p parser
		for i, pair := range tests {
			d1 := p.regexToDFA(pair.s1)
			d2 := p.regexToDFA(pair.s2)
			f1 := d1.isSubsetOf(d2)
			f2 := d2.isSubsetOf(d1)

			if !(f1 == answer1 && f2 == answer2) {
				t.Logf("expected  (%v, %v), got (%v, %v)\n", answer1, answer2, f1, f2)
				t.Logf("failed at %s[%d] = {'%s', '%s'}\n", s, i, pair.s1, pair.s2)

				// // for debug
				// if s == "less" {
				// 	fmt.Println("d1", d1)

				// 	fmt.Println("d2", d2)

				// 	d3 := d1.complement()
				// 	fmt.Println("-d1", d3)

				// 	d4 := d2.complement()
				// 	fmt.Println("-d2", d4)

				// 	d5 := intersectionTwoDFA(d1, d4)
				// 	fmt.Println("d1-d2", d5)

				// 	d6 := intersectionTwoDFA(d2, d3)
				// 	fmt.Println("d2-d1", d6)
				// }

				t.FailNow()
			}
			t.Logf("passed %d/%d", i+1, len(tests))
		}
	}

	testCase("equal", equal, true, true)
	testCase("less", less, true, false)
	testCase("greater", greater, false, true)
	testCase("noseq", noseq, false, false)
}
