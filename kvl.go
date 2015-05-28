// Package kvl provides a key/value list
package kvl

import (
	"strings"
)

// Pair implements a single key/value pair
type Pair [2]string

// Record implements a list of key/value pairs
type Record []Pair

// Get retrieves a value from the list as a string. It will return an empty
// string if no record with that key exists.
func (r Record) Get(key string) string {
	for _, p := range r {
		if p[0] == key {
			return p[1]
		}
	}

	return ""
}

// GetAll retrieves all the values for a given key from the list.
func (r Record) GetAll(key string) []string {
	var l []string

	for _, p := range r {
		if p[0] == key {
			l = append(l, p[1])
		}
	}

	return l
}

// Add adds a key/value pair to the list.
func (r Record) Add(key, value string) Record {
	return append(r, Pair{key, value})
}

// Set replaces the value of the first existing pair with the key specified, or,
// if a pair with that key is not found, inserts a new pair into the list.
func (r Record) Set(key, value string) Record {
	r2 := make(Record, len(r))
	copy(r2, r)

	for i, p := range r2 {
		if p[0] == key {
			r2[i][1] = value

			return r2
		}
	}

	return r2.Add(key, value)
}

// String renders the list as a string, preserving insertion order.
func (r Record) String() string {
	s := ""

	for i, p := range r {
		if i != 0 {
			s += " "
		}

		s += p[0]

		if p[1] != "" {
			s += "=" + p[1]
		}
	}

	return s
}

// Parse parses a string into key/value pairs.
func Parse(s string) Record {
	if s == "" {
		return nil
	}

	var r Record

	for _, s := range strings.Split(s, " ") {
		if len(s) == 0 {
			continue
		}

		p := strings.SplitN(s, "=", 2)

		if len(p) == 1 {
			r = append(r, Pair{p[0], ""})
		} else {
			r = append(r, Pair{p[0], p[1]})
		}
	}

	return r
}
