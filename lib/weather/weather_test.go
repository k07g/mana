package weather

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestToday(t *testing.T) {
	tests := []struct {
		name        string
		body        string
		statusCode  int
		wantErr     bool
		wantContain string
	}{
		{
			name: "正常系",
			body: `{
				"current_condition": [{
					"temp_C": "20",
					"lang_ja": [{"value": "晴れ"}]
				}]
			}`,
			statusCode:  http.StatusOK,
			wantErr:     false,
			wantContain: "晴れ",
		},
		{
			name: "気温が含まれる",
			body: `{
				"current_condition": [{
					"temp_C": "15",
					"lang_ja": [{"value": "曇り"}]
				}]
			}`,
			statusCode:  http.StatusOK,
			wantErr:     false,
			wantContain: "15°C",
		},
		{
			name: "current_conditionが空",
			body: `{
				"current_condition": []
			}`,
			statusCode: http.StatusOK,
			wantErr:    true,
		},
		{
			name: "lang_jaが空",
			body: `{
				"current_condition": [{
					"temp_C": "20",
					"lang_ja": []
				}]
			}`,
			statusCode: http.StatusOK,
			wantErr:    true,
		},
		{
			name:       "不正なJSON",
			body:       `invalid json`,
			statusCode: http.StatusOK,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.body))
			}))
			defer srv.Close()

			origURL := apiURL
			apiURL = srv.URL
			defer func() { apiURL = origURL }()

			msg, err := Today(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Today() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if tt.wantContain != "" && !strings.Contains(msg, tt.wantContain) {
				t.Errorf("Today() = %q, want to contain %q", msg, tt.wantContain)
			}
		})
	}
}
