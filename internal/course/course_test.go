package course

import (
	"strings"
	"testing"
)

func validCourse() Course {
	return Course{
		Title: "Paseo en Go",
		Modules: []Module{
			{Slug: "00-setup", Title: "Setup", Lessons: []Lesson{
				{Slug: "01-welcome", Title: "Welcome"},
				{Slug: "02-tools", Title: "Tools"},
			}},
			{Slug: "01-basics", Title: "Basics", Lessons: []Lesson{
				{Slug: "01-packages", Title: "Packages", Exercise: &Exercise{Path: "01-basics/hello", Goal: "Say hello"}},
			}},
		},
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*Course)
		wantErr string
	}{
		{name: "valid course", mutate: func(*Course) {}},
		{name: "missing course title", mutate: func(c *Course) { c.Title = "" }, wantErr: "course title"},
		{name: "no modules", mutate: func(c *Course) { c.Modules = nil }, wantErr: "no modules"},
		{name: "duplicate module slug", mutate: func(c *Course) { c.Modules[1].Slug = "00-setup" }, wantErr: `duplicate module "00-setup"`},
		{name: "module without lessons", mutate: func(c *Course) { c.Modules[1].Lessons = nil }, wantErr: "has no lessons"},
		{name: "duplicate lesson slug", mutate: func(c *Course) { c.Modules[0].Lessons[1].Slug = "01-welcome" }, wantErr: `duplicate lesson "00-setup/01-welcome"`},
		{name: "lesson without title", mutate: func(c *Course) { c.Modules[0].Lessons[0].Title = "" }, wantErr: "needs a title"},
		{name: "invalid slug", mutate: func(c *Course) { c.Modules[0].Slug = "Setup Go" }, wantErr: "invalid slug"},
		{name: "exercise without path", mutate: func(c *Course) { c.Modules[1].Lessons[0].Exercise.Path = "" }, wantErr: "exercise path"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := validCourse()
			tt.mutate(&c)
			err := c.Validate()
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("Validate() = %v, want nil", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("Validate() = %v, want error containing %q", err, tt.wantErr)
			}
		})
	}
}

func TestLessonsKeepsCourseOrder(t *testing.T) {
	got := validCourse().Lessons()

	want := []string{"00-setup/01-welcome", "00-setup/02-tools", "01-basics/01-packages"}
	if len(got) != len(want) {
		t.Fatalf("Lessons() returned %d refs, want %d", len(got), len(want))
	}
	for i, ref := range got {
		if ref.ID() != want[i] {
			t.Errorf("Lessons()[%d].ID() = %q, want %q", i, ref.ID(), want[i])
		}
	}
	if got[0].BodyPath() != "00-setup/01-welcome.html" {
		t.Errorf("BodyPath() = %q", got[0].BodyPath())
	}
}
