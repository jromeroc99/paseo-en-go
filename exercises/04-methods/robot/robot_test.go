package robot

import "testing"

func TestSpeak(t *testing.T) {
	r := &Robot{Name: "R2"}
	if got, want := r.Speak(), "I am R2"; got != want {
		t.Errorf(`(&Robot{Name: "R2"}).Speak() = %q, want %q`, got, want)
	}
}

func TestSay(t *testing.T) {
	var sleeping *Robot
	tests := []struct {
		name    string
		speaker Speaker
		want    string
	}{
		{name: "a robot", speaker: &Robot{Name: "Wall-E"}, want: "I am Wall-E"},
		{name: "a nil *Robot inside the interface", speaker: sleeping, want: "zzz"},
		{name: "a nil interface", speaker: nil, want: "nobody is there"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Say(tt.speaker); got != tt.want {
				t.Errorf("Say(%s) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
