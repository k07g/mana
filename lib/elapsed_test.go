package lib

import (
	"strings"
	"testing"
	"time"
)

func TestElapsedMessage(t *testing.T) {
	tests := []struct {
		elapsed  time.Duration
		wantPart string
	}{
		{elapsed: 30 * time.Second, wantPart: "秒"},
		{elapsed: 90 * time.Second, wantPart: "分"},
		{elapsed: 65 * time.Minute, wantPart: "時間"},
	}

	for _, tt := range tests {
		startTime := time.Now().Add(-tt.elapsed)
		msg := ElapsedMessage(startTime)
		if !strings.Contains(msg, tt.wantPart) {
			t.Errorf("ElapsedMessage(%v): got %q, want to contain %q", tt.elapsed, msg, tt.wantPart)
		}
	}
}
