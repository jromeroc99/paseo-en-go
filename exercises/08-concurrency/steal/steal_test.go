package steal

import (
	"slices"
	"testing"
)

func TestSteal(t *testing.T) {
	tests := []struct {
		name       string
		thief      []string
		victim     []string
		wantThief  []string
		wantVictim []string
	}{
		{
			name:       "an idle thief takes half of an even queue",
			thief:      []string{},
			victim:     []string{"a", "b", "c", "d"},
			wantThief:  []string{"a", "b"},
			wantVictim: []string{"c", "d"},
		},
		{
			name:       "an odd queue rounds up",
			thief:      []string{},
			victim:     []string{"a", "b", "c", "d", "e"},
			wantThief:  []string{"a", "b", "c"},
			wantVictim: []string{"d", "e"},
		},
		{
			name:       "a single task is stolen",
			thief:      []string{},
			victim:     []string{"a"},
			wantThief:  []string{"a"},
			wantVictim: []string{},
		},
		{
			name:       "an empty victim changes nothing",
			thief:      []string{},
			victim:     []string{},
			wantThief:  []string{},
			wantVictim: []string{},
		},
		{
			name:       "stolen tasks go after the thief's own",
			thief:      []string{"x"},
			victim:     []string{"a", "b"},
			wantThief:  []string{"x", "a"},
			wantVictim: []string{"b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thief := &P{Queue: slices.Clone(tt.thief)}
			victim := &P{Queue: slices.Clone(tt.victim)}
			Steal(thief, victim)
			if !slices.Equal(thief.Queue, tt.wantThief) {
				t.Errorf("Steal(thief %v, victim %v): thief.Queue = %v, want %v",
					tt.thief, tt.victim, thief.Queue, tt.wantThief)
			}
			if !slices.Equal(victim.Queue, tt.wantVictim) {
				t.Errorf("Steal(thief %v, victim %v): victim.Queue = %v, want %v",
					tt.thief, tt.victim, victim.Queue, tt.wantVictim)
			}
		})
	}
}
