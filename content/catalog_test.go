package content

import (
	"io/fs"
	"testing"
)

func TestCourseIsValid(t *testing.T) {
	if err := Course.Validate(); err != nil {
		t.Fatalf("Course.Validate() = %v", err)
	}
}

func TestEveryLessonHasBody(t *testing.T) {
	for _, ref := range Course.Lessons() {
		if _, err := fs.Stat(FS, ref.BodyPath()); err != nil {
			t.Errorf("lesson %q: missing body %q", ref.ID(), ref.BodyPath())
		}
	}
}
