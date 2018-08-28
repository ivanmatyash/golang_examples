package anagram

import "testing"

func TestAnagram(t *testing.T) {
	cases := []struct {
		in1, in2 string
		expected bool
	}{
		{"a", "aa", false},
		{"aaa", "aab", false},
		{"aaa", "aaa", true},
		{"", "", true},
		{"abc", "cba", true},
		{"123", "321", true},
		{"привет", "тевирп", true},
		{"тест", "аааа", false},
	}

	for _, c := range cases {
		actual := isAnagram(c.in1, c.in2)
		if actual != c.expected {
			t.Errorf("IsAnagram: actual=%t, expected=%t, s1=%s, s2=%s.", actual, c.expected, c.in1, c.in2)
		}
	}
}
