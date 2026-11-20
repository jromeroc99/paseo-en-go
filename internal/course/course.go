// Package course models the structure of the tour: a course made of
// modules, each made of lessons, some of which point to a local exercise.
// It knows nothing about HTTP or templates.
package course

import (
	"errors"
	"fmt"
	"regexp"
)

// Course is the whole tour.
type Course struct {
	Title   string
	Modules []Module
}

// Module groups related lessons, e.g. "Pointers and structs".
type Module struct {
	Slug    string
	Title   string
	Summary string
	Lessons []Lesson
}

// Lesson is a single page of the tour.
type Lesson struct {
	Slug     string
	Title    string
	Exercise *Exercise // nil when the lesson has no exercise
}

// Exercise points to a folder inside the exercises/ module that the
// student solves locally and validates with `go test`.
type Exercise struct {
	Path string // relative to exercises/, e.g. "01-basics/hello"
	Goal string
}

// Ref identifies a lesson together with the module it belongs to.
type Ref struct {
	Module Module
	Lesson Lesson
}

// ID is the stable identifier of a lesson, e.g. "01-basics/01-packages".
// It is used in URLs and as the key for saved progress.
func (r Ref) ID() string {
	return r.Module.Slug + "/" + r.Lesson.Slug
}

// BodyPath is where the lesson body template lives inside the content FS.
func (r Ref) BodyPath() string {
	return r.ID() + ".html"
}

// Lessons returns every lesson of the course in reading order.
func (c Course) Lessons() []Ref {
	var refs []Ref
	for _, m := range c.Modules {
		for _, l := range m.Lessons {
			refs = append(refs, Ref{Module: m, Lesson: l})
		}
	}
	return refs
}

var slugPattern = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

// Validate reports the first structural problem found in the course.
func (c Course) Validate() error {
	if c.Title == "" {
		return errors.New("course title is required")
	}
	if len(c.Modules) == 0 {
		return errors.New("course has no modules")
	}

	modules := map[string]bool{}
	for _, m := range c.Modules {
		if err := validateModule(m); err != nil {
			return err
		}
		if modules[m.Slug] {
			return fmt.Errorf("duplicate module %q", m.Slug)
		}
		modules[m.Slug] = true
	}
	return nil
}

func validateModule(m Module) error {
	if !slugPattern.MatchString(m.Slug) {
		return fmt.Errorf("module %q: invalid slug", m.Slug)
	}
	if m.Title == "" {
		return fmt.Errorf("module %q needs a title", m.Slug)
	}
	if len(m.Lessons) == 0 {
		return fmt.Errorf("module %q has no lessons", m.Slug)
	}

	lessons := map[string]bool{}
	for _, l := range m.Lessons {
		id := Ref{Module: m, Lesson: l}.ID()
		if !slugPattern.MatchString(l.Slug) {
			return fmt.Errorf("lesson %q: invalid slug", id)
		}
		if l.Title == "" {
			return fmt.Errorf("lesson %q needs a title", id)
		}
		if l.Exercise != nil && l.Exercise.Path == "" {
			return fmt.Errorf("lesson %q: exercise path is required", id)
		}
		if lessons[l.Slug] {
			return fmt.Errorf("duplicate lesson %q", id)
		}
		lessons[l.Slug] = true
	}
	return nil
}
