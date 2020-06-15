package utils

import "testing"

var underscoreTestCases = []struct {
	expected string
	real     string
}{
	{real: "Ab", expected: "ab"},
	{real: "Abc", expected: "abc"},
	{real: "AbC", expected: "ab_c"},
}

func TestUnderscore(t *testing.T) {
	for _, item := range underscoreTestCases {
		if Underscore(item.real) != item.expected {
			t.Errorf("Error transforming %s\n", item.real)
		}
	}
}

var camelizeTestCases = []struct {
	expected string
	real     string
}{
	{real: "ab", expected: "Ab"},
	{real: "abc", expected: "Abc"},
	{real: "ab c", expected: "AbC"},
}

func TestCamelize(t *testing.T) {
	for _, item := range camelizeTestCases {
		if Camelize(item.real) != item.expected {
			t.Errorf("Error transforming %s\n", item.real)
		}
	}
}
