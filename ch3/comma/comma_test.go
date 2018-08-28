package comma

import "testing"

func TestComma(t *testing.T) {
	cases := map[string]string{
		"123":     "123",
		"1234":    "1,234",
		"12345":   "12,345",
		"123456":  "123,456",
		"1234567": "1,234,567",
	}

	for in, expected := range cases {
		actual1 := comma(in)
		if actual1 != expected {
			t.Errorf("Comma: actual=%s, expected=%s", actual1, expected)
		}

		actual2 := commaNotRecursive(in)
		if actual2 != expected {
			t.Errorf("Comma not recursive: actual=%s, expected=%s", actual2, expected)
		}
	}
}
