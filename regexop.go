package regexop

// UnionRegex returns the union of multiple regular expressions
func UnionRegex(regexes ...string) string {
	if len(regexes) < 1 {
		return ""
	}

	ret := regexes[0]
	for i := 1; i < len(regexes); i++ {
		ret = unionTwoRegex(ret, regexes[i])
	}

	return ret
}

// IntersectionRegex returns the intersection of multiple regular expressions
func IntersectionRegex(regexes ...string) string {
	if len(regexes) < 1 {
		return ""
	}

	ret := regexes[0]
	for i := 1; i < len(regexes); i++ {
		ret = intersectionTwoRegex(ret, regexes[i])
	}

	return ret
}

// DiffRegex returns a new regular expression means `regex1 - regex2`
func DiffRegex(regex1, regex2 string) string {
	var p parser
	d1 := p.regexToDFA(regex1)
	d2 := p.regexToDFA(regex2)
	d := d1.difference(d2)
	return d.toRegex()
}

// ComplementRegex returns a new regular expression means `clousureOfAlphabet - regex` i.e. the complement of regex
func ComplementRegex(regex string) string {
	var p parser
	d := p.regexToDFA(regex)
	c := d.complement()
	return c.toRegex()
}

// IsSubsetOf returns weather the first regular expression is the subset of the rest regular expressions
func IsSubsetOf(regexes ...string) bool {
	if len(regexes) < 2 {
		return false
	}

	target := regexes[0]
	for i := 1; i < len(regexes); i++ {
		if isSubsetTwoRegex(target, regexes[i]) {
			return true
		}
	}

	return false
}

// IsEqualTo returns weather the regex1 is euqal to regex2, i.e. both accept the same language
func IsEqualTo(regex1, regex2 string) bool {
	return IsSubsetOf(regex1, regex2) && IsSubsetOf(regex2, regex1)
}

// ---------------------------------------------------------------------------------------

func unionTwoRegex(regex1, regex2 string) string {
	var p parser
	d1 := p.regexToDFA(regex1)
	d2 := p.regexToDFA(regex2)
	u := unionTwoDFA(d1, d2)
	return u.toRegex()
}

func intersectionTwoRegex(regex1, regex2 string) string {
	var p parser
	d1 := p.regexToDFA(regex1)
	d2 := p.regexToDFA(regex2)
	i := intersectionTwoDFA(d1, d2)
	return i.toRegex()
}

func isSubsetTwoRegex(regex1, regex2 string) bool {
	var p parser
	d1 := p.regexToDFA(regex1)
	d2 := p.regexToDFA(regex2)
	return d1.isSubsetOf(d2)
}
