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

// IsSubsetOf return weather the first regular expression is the subset of the rest regular expressions
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

func unionTwoRegex(regex1, regex2 string) string {
	return ""
}

func intersectionTwoRegex(regex1, regex2 string) string {
	return ""
}

func isSubsetTwoRegex(regex1, regex2 string) bool {
	return false
}
