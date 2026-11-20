package content

import (
	"io/fs"
	"regexp"
	"testing"

	"paseo-en-go/web"
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

// animationTag matches the opening tag of an animation, e.g. <pointer-demo.
var animationTag = regexp.MustCompile(`<([a-z]+-demo)\b`)

// The page loads animations/<tag>.js for each tag it finds, so a tag
// without its file would leave the lesson showing only the text fallback.
func TestEveryAnimationHasComponent(t *testing.T) {
	for _, ref := range Course.Lessons() {
		body, err := fs.ReadFile(FS, ref.BodyPath())
		if err != nil {
			continue // reported by TestEveryLessonHasBody
		}
		for _, match := range animationTag.FindAllSubmatch(body, -1) {
			file := "static/js/animations/" + string(match[1]) + ".js"
			if _, err := fs.Stat(web.FS, file); err != nil {
				t.Errorf("lesson %q uses <%s>, but web/%s does not exist", ref.ID(), match[1], file)
			}
		}
	}
}
