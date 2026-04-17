package news

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
		wantErr     bool
		wantContain string
	}{
		{
			name: "正常系",
			body: `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel>
    <item><title>ニュース1</title></item>
    <item><title>ニュース2</title></item>
    <item><title>ニュース3</title></item>
    <item><title>ニュース4</title></item>
  </channel>
</rss>`,
			wantErr:     false,
			wantContain: "ニュース1",
		},
		{
			name: "3件までしか返さない",
			body: `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel>
    <item><title>ニュース1</title></item>
    <item><title>ニュース2</title></item>
    <item><title>ニュース3</title></item>
    <item><title>ニュース4</title></item>
  </channel>
</rss>`,
			wantErr:     false,
			wantContain: "ニュース3",
		},
		{
			name: "itemが空",
			body: `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel></channel>
</rss>`,
			wantErr: true,
		},
		{
			name:    "不正なXML",
			body:    `invalid xml`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(tt.body))
			}))
			defer srv.Close()

			origURL := rssURL
			rssURL = srv.URL
			defer func() { rssURL = origURL }()

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

func TestTodayLimit(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
  <channel>
    <item><title>ニュース1</title></item>
    <item><title>ニュース2</title></item>
    <item><title>ニュース3</title></item>
    <item><title>ニュース4</title></item>
  </channel>
</rss>`))
	}))
	defer srv.Close()

	origURL := rssURL
	rssURL = srv.URL
	defer func() { rssURL = origURL }()

	msg, err := Today(context.Background())
	if err != nil {
		t.Fatalf("Today() error = %v", err)
	}
	if strings.Contains(msg, "ニュース4") {
		t.Errorf("Today() should not contain 4th item, got %q", msg)
	}
}
