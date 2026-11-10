package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	type tests map[string]struct {
		input    string
		expected []string
	}

	cases := tests{
		"simple":    {input: " hello world ", expected: []string{"hello", "world"}},
		"no spaces": {input: "hello", expected: []string{"hello"}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(tc.input)
			diff := cmp.Diff(actual, tc.expected)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
