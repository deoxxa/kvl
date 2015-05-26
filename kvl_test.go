package kvl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tcase struct {
	s string
	r Record
}

var goodcases = []tcase{
	{"", nil},
	{"a=b", Record{
		{"a", "b"},
	}},
	{"a=b c=d", Record{
		{"a", "b"},
		{"c", "d"},
	}},
	{"a=1 a=2", Record{
		{"a", "1"},
		{"a", "2"},
	}},
}

var badcases = []tcase{
	{"", Record{
		{"", ""},
	}},
	{"a=b", Record{
		{"a", ""},
	}},
	{"a=b", Record{
		{"", "b"},
	}},
	{"a=b", Record{
		{"b", "a"},
	}},
	{"a=b c=d", Record{
		{"c", "d"},
		{"a", "b"},
	}},
	{"a=1 a=2", Record{
		{"a", "2"},
		{"a", "1"},
	}},
}

func TestParse(t *testing.T) {
	a := assert.New(t)

	for _, c := range goodcases {
		a.Equal(c.r, Parse(c.s), c.s)
	}

	for _, c := range badcases {
		a.NotEqual(c.r, Parse(c.s), c.s)
	}
}

func TestString(t *testing.T) {
	a := assert.New(t)

	for _, c := range goodcases {
		a.Equal(c.s, c.r.String(), c.s)
	}

	for _, c := range badcases {
		a.NotEqual(c.s, c.r.String(), c.s)
	}
}

func TestGet(t *testing.T) {
	a := assert.New(t)

	r := Record{
		{"a", "b"},
		{"a", "c"},
		{"x", "y"},
	}

	a.Equal("b", r.Get("a"))
	a.Equal([]string{"b", "c"}, r.GetAll("a"))
	a.Equal("y", r.Get("x"))
	a.Equal([]string{"y"}, r.GetAll("x"))
	a.Equal("", r.Get("z"))
}