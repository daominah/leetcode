package p0010_regular_expression_matching

import (
	"fmt"
	"regexp"
)

func isMatch(s string, p string) bool {
	// haha, beats 100% submissions runtime
	pattern, err := regexp.Compile("^" + p + "$")
	if err != nil {
		fmt.Println("invalid pattern")
		return false
	}
	return pattern.MatchString(s)
}
