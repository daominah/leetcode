package p0010_regular_expression_matching

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	for i, c := range []struct {
		s    string
		p    string
		want bool
	}{
		{s: "aa", p: "a", want: false},
		{s: "aa", p: "a*", want: true},
		{s: "ab", p: ".*", want: true},
	} {
		got := isMatch(c.s, c.p)
		if got != c.want {
			t.Errorf("error isMatch i %v got %v, want %v", i, got, c.want)
		}
	}
}
