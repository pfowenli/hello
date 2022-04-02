package morestrings

import "testing"

type Case struct {
	in   string
	want string
}

func TestReverseRunes(t *testing.T) {
	cases := []Case{
		{"Hello, Go", "oG ,olleH"},
		{"Hello, World", "dlroW ,olleH"},
		{"Hallo, Werlt", "tlreW ,ollaH"},
		{"", ""},
	}

	for _, c := range cases {
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
