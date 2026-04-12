package lib

import "testing"

func TestBedtimeMessage(t *testing.T) {
	tests := []struct {
		hour    int
		wantOk  bool
	}{
		{hour: 22, wantOk: true},
		{hour: 23, wantOk: true},
		{hour: 0, wantOk: true},
		{hour: 1, wantOk: true},
		{hour: 3, wantOk: true},
		{hour: 4, wantOk: false},
		{hour: 12, wantOk: false},
		{hour: 21, wantOk: false},
	}

	for _, tt := range tests {
		msg, ok := BedtimeMessage(tt.hour)
		if ok != tt.wantOk {
			t.Errorf("BedtimeMessage(%d): ok = %v, want %v", tt.hour, ok, tt.wantOk)
		}
		if ok && msg == "" {
			t.Errorf("BedtimeMessage(%d): message is empty", tt.hour)
		}
		if !ok && msg != "" {
			t.Errorf("BedtimeMessage(%d): expected empty message, got %q", tt.hour, msg)
		}
	}
}
