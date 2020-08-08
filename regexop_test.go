package regexop

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"testing"
)

// auto random integration testing in this package

// Note：
// This testing uses Go standard package `regexp` to generate answers of test cases.
// However, there are some differences of regular expression grammars between the  in this package and those in the package `regexp`.
// And, the useful solution to solve it is to wrap the whole regular expression with brackets.

var (
	constNumSample         int = 100
	constRandomStringScale int = 20

	symbols = []symbol{'a', 'b'}

	regexes = []string{
		`(a|b)*abb`,
		`a*b*`,
		`ab(a|b)*`,
		`a?`,
		`a+`,
		`(bbb(a?)(b*))+`,
		`a*`,
		`a(a|b)*`,
		`b(a|b)*`,
		`ab+`,
		`((a+)|(b+))`,
		`(a|b)(a|b)(a|b)(a|b)(a|b)`,
		`(a|b)*a(a|b)(a|b)(a|b)(a|b)(a|b)(a|b)(a|b)(a|b)(a|b)`,
		`((a|b)(a|b))`,
	}
)

func randomStrings(symbols []symbol, numSample int, scale int) []string {
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
