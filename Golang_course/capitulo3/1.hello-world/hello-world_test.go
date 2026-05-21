// -*- coding: utf-8 -*-
// hello-world_test.go
// -----------------------------------------------------------------------------
//
// Started on <mar 17-02-2026 00:30:55.576444307 (1771284655)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Pruebas de la función greetings
package main

import "testing"

func TestGreetings(t *testing.T) {
	got := greetings()
	expected := []string{
		"Hello World!",
		"Hallo Welt!",
	}
	if len(got) != len(expected) {
		t.Fatalf("unexpected length: got %d, want %d", len(got), len(expected))
	}
	for i := range expected {
		if got[i] != expected[i] {
			t.Errorf("at index %d: got %q, want %q", i, got[i], expected[i])
		}
	}
}

// Local Variables:
// mode:go
// fill-column:80
// End:
