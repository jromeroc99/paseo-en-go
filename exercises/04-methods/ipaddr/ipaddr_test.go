package ipaddr

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		ip   IPAddr
		want string
	}{
		{ip: IPAddr{127, 0, 0, 1}, want: "127.0.0.1"},
		{ip: IPAddr{8, 8, 8, 8}, want: "8.8.8.8"},
		{ip: IPAddr{192, 168, 1, 254}, want: "192.168.1.254"},
		{ip: IPAddr{}, want: "0.0.0.0"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := fmt.Sprint(tt.ip); got != tt.want {
				t.Errorf("fmt.Sprint(IPAddr%v) = %q, want %q", [4]byte(tt.ip), got, tt.want)
			}
		})
	}
}
