package VMware

import (
	"net/url"
	"testing"
)

func TestClient_Metascan(t *testing.T) {
	u, err := url.Parse("https://192.168.59.4/sdk")
	if err != nil {
		t.Fatalf("failed to parse url")
	}

	tests := []struct {
		name string
		tags []string

		wantErr bool
	}{
		{"", []string{"windows", "PRTG"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewClient(u, "prtg@heynes.local", ".l3tm31n")
			if err != nil {
				t.Fatal("cant get client")
			}
			gotRtnMap := NewTagMap()
			if err := c.Metascan(tt.tags, gotRtnMap); (err != nil) != tt.wantErr {
				t.Errorf("Metascan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
