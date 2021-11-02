package main

import "testing"

func TestRepeated(t *testing.T) {
	repeated := Repeate("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q', but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeate("a", 1)
	}
}
