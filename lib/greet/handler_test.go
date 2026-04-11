package greet

import "testing"

func TestGreetHandler(t *testing.T) {
	tests := []struct {
		input  string
		wantOk bool
	}{
		{"おはよ", true},
		{"おはよう", true},
		{"こんにちは", true},
		{"こんばんは", true},
		{"おやすみ", true},
		{"さようなら", true},
		{"バイバイ", true},
		{"", false},
		{"hello", false},
		{"1 + 1", false},
	}

	for _, tt := range tests {
		message, ok := GreetHandler(tt.input)
		if ok != tt.wantOk {
			t.Errorf("GreetHandler(%q): ok = %v, want %v", tt.input, ok, tt.wantOk)
		}
		if ok && message == "" {
			t.Errorf("GreetHandler(%q): message is empty", tt.input)
		}
		if !ok && message != "" {
			t.Errorf("GreetHandler(%q): expected empty message, got %q", tt.input, message)
		}
	}
}
