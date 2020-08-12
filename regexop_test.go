package regexop

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"testing"
	"time"
)

// auto random integration testing in this package

// Note：
// This testing uses Go standard package `regexp` to generate answers of test cases.
// However, there are some differences between the regular expression grammars of this package and those in package `regexp`.
// And, the useful solution to solve it is to wrap the whole regular expression with brackets.

var (
	_randSeed int64 = 0

	constNumSample         int = 500
	constRandomStringScale int = 20

	symbols = []symbol{'a', 'b'}

	// required:
	// `a` and `b` shoud occur at the same time in a regex, such as `a+` and `b?` are invalid
	regexes = []string{
		`((a|b)*abb)`,
		`(a*b*)`,
		`(ab(a|b)*)`,
		`(a?b)`,
		`(a+b)`,
		`((bbb(a?)(b*))+)`,
		`((a*)b)`,
		`(a(a|b)*)`,
		`(b(a|b)*)`,
		`(ab+)`,
		`((a+)|(b+))`,
		`(a|b)(a|b)(a|b)(a|b)(a|b)`,
		`(a|b)*a(a|b)(a|b)(a|b)`,
		`((a|b)(a|b))`,
		`(aa(a|b))`,
		`((a|b)*aabb)`,
		`a(a|b)*b`,
		`b(a|b)*a`,
		`aaaaaa(b)*`,
		`(((((a|b)+))))`,
		`(a*|b?)`,
	}
)

func randSeed() int64 {
	if _randSeed == 0 {
		_randSeed = time.Now().Unix()
	}
	return _randSeed
}

func randomStrings(symbols []symbol, numSample int, scale int) []string {
	rand.Seed(randSeed())
	ret := []string{}
	for i := 0; i < numSample; i++ {
		var b strings.Builder
		strLen := rand.Intn(2 * scale)
		for i := 0; i < strLen; i++ {
			r := symbols[rand.Intn(len(symbols))]
			b.WriteRune(rune(r))
		}
		ret = append(ret, b.String())
	}
	return ret
}

func TestGoStdRegexp(t *testing.T) {
	re := regexp.MustCompile("(a|b)*abb")
	fmt.Println(re.MatchString("ab"))       // false
	fmt.Println(re.MatchString("abab"))     // false
	fmt.Println(re.MatchString("abb"))      // true
	fmt.Println(re.MatchString("ababaabb")) // true
}

func TestRegex2NFA(t *testing.T) {

	var p parser

	for i, regex := range regexes {

		n := p.regexToNFA(regex)

		re := regexp.MustCompile("^" + regex + "$") // ......

		tests := randomStrings(symbols, constNumSample, constRandomStringScale)

		for j, str := range tests {
			predict := n.accept(str)
			answer := re.MatchString(str)
			if predict != answer {
				t.Logf("answer = %v, predict = %v\n", answer, predict)
				t.Logf("failed at regexes[%d] = '%s' with samples[%d] '%s' \n", i, regex, j, str)
				t.FailNow()
			}
		}

		t.Logf("passed %d/%d\n", i+1, len(regexes))
	}

}

func TestRegex2DFA(t *testing.T) {

	var p parser

	for i, regex := range regexes {

		d := p.regexToDFA(regex)

		re := regexp.MustCompile("^" + regex + "$") // ......

		tests := randomStrings(symbols, constNumSample, constRandomStringScale)

		for j, str := range tests {
			predict := d.accept(str)
			answer := re.MatchString(str)
			if predict != answer {
				t.Logf("answer = %v, predict = %v\n", answer, predict)
				t.Logf("failed at regexes[%d] = '%s' with samples[%d] '%s' \n", i, regex, j, str)
				t.FailNow()
			}
		}

		t.Logf("passed %d/%d\n", i+1, len(regexes))
	}

}

func TestRegex2DFAandDFAComplement1(t *testing.T) {
	var p parser

	for i, regex := range regexes {

		d := p.regexToDFA(regex).complement()

		re := regexp.MustCompile("^" + regex + "$") // ......

		tests := randomStrings(symbols, constNumSample, constRandomStringScale)

		for j, str := range tests {
			predict := !d.accept(str)
			answer := re.MatchString(str)
			if predict != answer {
				t.Logf("answer = %v, predict = %v\n", answer, predict)
				t.Logf("failed at regexes[%d] = '%s' with samples[%d] '%s' \n", i, regex, j, str)
				t.FailNow()
			}
		}

		t.Logf("passed %d/%d\n", i+1, len(regexes))
	}
}

func TestRegex2DFAandDFAComplement2(t *testing.T) {
	var p parser

	for i, regex := range regexes {

		d := p.regexToDFA(regex).complement().complement()

		re := regexp.MustCompile("^" + regex + "$") // ......

		tests := randomStrings(symbols, constNumSample, constRandomStringScale)

		for j, str := range tests {
			predict := d.accept(str)
			answer := re.MatchString(str)
			if predict != answer {
				t.Logf("answer = %v, predict = %v\n", answer, predict)
				t.Logf("failed at regexes[%d] = '%s' with samples[%d] '%s' \n", i, regex, j, str)
				t.FailNow()
			}
		}

		t.Logf("passed %d/%d\n", i+1, len(regexes))
	}
}

func TestRegex2DFAandDFAIntersection(t *testing.T) {
	var p parser

	for i, regex1 := range regexes {

		re1 := regexp.MustCompile("^" + regex1 + "$")
		d1 := p.regexToDFA(regex1)

		for j, regex2 := range regexes {

			re2 := regexp.MustCompile("^" + regex2 + "$")
			d2 := p.regexToDFA(regex2)

			interDFA := intersectionTwoDFA(d1, d2)

			tests := randomStrings(symbols, constNumSample, constRandomStringScale)

			for k, str := range tests {
				predict := interDFA.accept(str)
				answer1 := re1.MatchString(str)
				answer2 := re2.MatchString(str)
				answer := answer1 && answer2 // intersection
				if predict != answer {
					t.Logf("answer = %v, predict = %v\n", answer, predict)
					t.Logf("failed at regexes[%d] = '%s' , regexs[%d] = %s, with samples[%d] '%s' \n", i, regex1, j, regex2, k, str)
					t.FailNow()
				}
			}
		}

		t.Logf("passed %d/%d\n", i+1, len(regexes))
	}

}

func TestRegex2DFAandDFADifference(t *testing.T) {
	var p parser

	for i, regex1 := range regexes {

		re1 := regexp.MustCompile("^" + regex1 + "$")
		d1 := p.regexToDFA(regex1)

		for j, regex2 := range regexes {

			re2 := regexp.MustCompile("^" + regex2 + "$")
			d2 := p.regexToDFA(regex2)

			diffDFA := d1.difference(d2)

			tests := randomStrings(symbols, constNumSample, constRandomStringScale)

			for k, str := range tests {
				predict := diffDFA.accept(str)
				answer1 := re1.MatchString(str)
				answer2 := re2.MatchString(str)
				answer := answer1 && !answer2

				if predict != answer {
					t.Logf("answer = %v, predict = %v\n", answer, predict)
					t.Logf("failed at regexes[%d] = '%s' , regexs[%d] = %s, with samples[%d] '%s' \n", i, regex1, j, regex2, k, str)
					t.FailNow()
				}
			}
		}

		t.Logf("passed %d/%d\n", i+1, len(regexes))
	}
}

func TestRegex2DFAandDFAUnion(t *testing.T) {
	var p parser

	for i, regex1 := range regexes {

		re1 := regexp.MustCompile("^" + regex1 + "$")
		d1 := p.regexToDFA(regex1)

		for j, regex2 := range regexes {

			re2 := regexp.MustCompile("^" + regex2 + "$")
			d2 := p.regexToDFA(regex2)

			unDFA := unionTwoDFA(d1, d2)

			tests := randomStrings(symbols, constNumSample, constRandomStringScale)

			for k, str := range tests {
				predict := unDFA.accept(str)
				answer1 := re1.MatchString(str)
				answer2 := re2.MatchString(str)
				answer := answer1 || answer2 // union

				if predict != answer {
					t.Logf("answer = %v, predict = %v\n", answer, predict)
					t.Logf("failed at regexes[%d] = '%s' , regexs[%d] = %s, with samples[%d] '%s' \n", i, regex1, j, regex2, k, str)
					t.FailNow()
				}
			}
		}

		t.Logf("passed %d/%d\n", i+1, len(regexes))
	}
}

func TestRegex2DFAMinimize1(t *testing.T) {

	var p parser

	for i, regex := range regexes {

		d := p.regexToDFA(regex)
		md := d.minimize()
		re := regexp.MustCompile("^" + regex + "$") // ......

		tests := randomStrings(symbols, constNumSample, constRandomStringScale)

		for j, str := range tests {
			predict1 := d.accept(str)
			predict2 := md.accept(str)
			answer := re.MatchString(str)
			if !(predict1 == predict2 && predict2 == answer) {
				t.Logf("answer = %v, predict1 = %v, predict2 = %v\n", answer, predict1, predict2)
				t.Logf("failed at regexes[%d] = '%s' with samples[%d] '%s' \n", i, regex, j, str)
				t.FailNow()
			}
		}

		t.Logf("passed %d/%d\n", i+1, len(regexes))
	}

}

func TestRegex2DFAMinimize2(t *testing.T) {

	var p parser

	for i, regex := range regexes {
		md := p.regexToDFA(regex).minimize()
		re := regexp.MustCompile("^" + regex + "$") // ......

		tests := randomStrings(symbols, constNumSample, constRandomStringScale)

		for j, str := range tests {
			answer := re.MatchString(str)
			predict1 := answer // omit
			predict2 := md.accept(str)
			if !(predict1 == predict2 && predict2 == answer) {
				t.Logf("answer = %v, predict1 = %v, predict2 = %v\n", answer, predict1, predict2)
				t.Logf("failed at regexes[%d] = '%s' with samples[%d] '%s' \n", i, regex, j, str)
				t.FailNow()
			}
		}

		t.Logf("passed %d/%d\n", i+1, len(regexes))
	}

}

func TestRegexOpExamples(t *testing.T) {
	var uRegex string = UnionRegex("a", "b", "c")
	// uRegex = "a|b|c"

	var s bool = IsSubsetOf("a*b", "a*")
	// s = false

	var e bool = IsEqualTo("a+", "aa*")
	// e = true

	if !IsEqualTo(uRegex, "a|b|c") {
		t.Fail()
	}
	if !(s == false) {
		t.Fail()
	}
	if !(e == true) {
		t.Fail()
	}
}

func TestDFA2Regex1(t *testing.T) {

	var p parser

	for i, regex := range regexes {
		md := p.regexToDFA(regex).minimize()
		d2rd := regexp.MustCompile("^" + "(" + md.toRegex() + ")" + "$")
		re := regexp.MustCompile("^" + regex + "$") // ......

		tests := randomStrings(symbols, constNumSample, constRandomStringScale)

		for j, str := range tests {
			answer := re.MatchString(str)
			predict1 := md.accept(str) // omit
			predict2 := d2rd.MatchString(str)
			if !(predict1 == predict2 && predict2 == answer) {
				t.Logf("answer = %v, predict1 = %v, predict2 = %v\n", answer, predict1, predict2)
				t.Logf("failed at regexes[%d] = '%s' with samples[%d] '%s' \n", i, regex, j, str)
				t.FailNow()
			}
		}

		t.Logf("passed %d/%d\n", i+1, len(regexes))
	}

}

func TestDFA2Regex2(t *testing.T) {

	var p parser

	for i, regex := range regexes {
		md := p.regexToDFA(regex).minimize()
		d2rd := p.regexToDFA(md.toRegex())

		// if !IsEqualTo(regex, md.toRegex()) {
		// 	t.Logf("IsEqualTo failed at regexes[%d] = '%s'\n", i, regex)
		// 	t.FailNow()
		// }

		re := regexp.MustCompile("^" + regex + "$") // ......

		tests := randomStrings(symbols, constNumSample, constRandomStringScale)

		for j, str := range tests {
			answer := re.MatchString(str)
			predict1 := md.accept(str) // omit
			predict2 := d2rd.accept(str)
			if !(predict1 == predict2 && predict2 == answer) {
				t.Logf("answer = %v, predict1 = %v, predict2 = %v\n", answer, predict1, predict2)
				t.Logf("failed at regexes[%d] = '%s' with samples[%d] '%s' \n", i, regex, j, str)
				t.FailNow()
			}
		}

		fmt.Printf("passed %d/%d\n", i+1, len(regexes))
	}

}
