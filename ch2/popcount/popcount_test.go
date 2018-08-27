package popcount

import (
	"testing"
)

const (
	value           = 184467440737095516
	expected        = 30
	InvalidValueErr = "Error. Actual: %d, expected: %d."
)

func TestPopcount1(t *testing.T) {
	amount := popCount1(value)
	if amount != expected {
		t.Errorf(InvalidValueErr, amount, expected)
	}
}

func TestPopcount2(t *testing.T) {
	amount := popCount2(value)
	if amount != expected {
		t.Errorf(InvalidValueErr, amount, expected)
	}
}

func TestPopcount3(t *testing.T) {
	amount := popCount3(value)
	if amount != expected {
		t.Errorf(InvalidValueErr, amount, expected)
	}
}

func TestPopcount4(t *testing.T) {
	amount := popCount4(value)
	if amount != expected {
		t.Errorf(InvalidValueErr, amount, expected)
	}
}

func BenchmarkPopcount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount1(value)
	}
}

func BenchmarkPopcount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount2(value)
	}
}

func BenchmarkPopcount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount3(value)
	}
}

func BenchmarkPopcount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount4(value)
	}
}
