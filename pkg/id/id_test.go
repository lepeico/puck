package id

import "testing"

func TestNew(t *testing.T) {
	m := make(map[string]bool)
	for x := 1; x < 32; x++ {
		s := New(12)
		if m[s.String()] {
			t.Errorf("New returned duplicated ID %s", s)
		}
		m[s.String()] = true
		if s.String() != FromString(s.String()).String() {
			t.Errorf("FromString(New.String()) returned %q which does not equal to New", s)
		}
	}
}
