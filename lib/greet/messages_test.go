package greet

import (
	"slices"
	"testing"
)

func TestGoodMorning(t *testing.T) {
	for range 20 {
		msg := goodMorning()
		if !slices.Contains(goodMorningMessages, msg) {
			t.Errorf("goodMorning() = %q, not in goodMorningMessages", msg)
		}
	}
}

func TestGoodNight(t *testing.T) {
	for range 20 {
		msg := goodNight()
		if !slices.Contains(goodNightMessages, msg) {
			t.Errorf("goodNight() = %q, not in goodNightMessages", msg)
		}
	}
}

func TestGoodBye(t *testing.T) {
	if got := goodBye(); got == "" {
		t.Error("goodBye() returned empty string")
	}
}
