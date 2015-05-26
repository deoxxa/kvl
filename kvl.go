package kvl

import (
	"strings"
)

type Record [][2]string

func (r Record) Get(s string) string {
	for _, p := range r {
		if p[0] == s {
			return p[1]
		}
	}

	return ""
}

func (r Record) GetAll(s string) []string {
	var l []string

	for _, p := range r {
		if p[0] == s {
			l = append(l, p[1])
		}
	}

	return l
}

func (r Record) String() string {
	s := ""

	for i, p := range r {
		if i != 0 {
			s += " "
		}

		s += p[0] + "=" + p[1]
	}

	return s
}

func Parse(s string) Record {
	if s == "" {
		return nil
	}

	var r Record

	for _, s := range strings.Split(s, " ") {
		p := strings.SplitN(s, "=", 2)
		r = append(r, [2]string{p[0], p[1]})
	}

	return r
}
