package content

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"path"
	"regexp"
	"strings"
	"testing"

	"paseo-en-go/internal/course"
)

// The budget of one card, a testable stand-in for "fits a 1280×720 screen".
const (
	maxCardWords    = 90 // prose words, headings and code blocks excluded
	maxExampleLines = 20 // a longer example gets a card of its own
	maxCardPrograms = 1  // {{code}} snippets and Go code blocks
	maxCardOutputs  = 1  // plain <pre> blocks: a program's output or an error
)

// codeAction matches a {{code "examples/x/main.go"}} action and its file.
var codeAction = regexp.MustCompile(`\{\{\s*code\s+"([^"]+)"\s*\}\}`)

// card is what the budget checks need to know about one slide.
type card struct {
	words      int      // prose words outside headings, code and animations
	programs   int      // Go code blocks, embedded or inline
	outputs    int      // plain <pre> blocks
	animations int      // <*-demo> elements
	others     []string // elements beside an <h2> and an animation
	examples   []string // files embedded with {{code}}
}

func TestEveryLessonIsSlides(t *testing.T) {
	for _, ref := range Course.Lessons() {
		body, err := fs.ReadFile(FS, ref.BodyPath())
		if err != nil {
			continue // reported by TestEveryLessonHasBody
		}
		if _, err := parseCards(body); err != nil {
			t.Errorf("lesson %q is not a deck: %v", ref.ID(), err)
		}
	}
}

func TestSlideBudget(t *testing.T) {
	for _, ref := range Course.Lessons() {
		body, err := fs.ReadFile(FS, ref.BodyPath())
		if err != nil {
			continue // reported by TestEveryLessonHasBody
		}
		cards, err := parseCards(body)
		if err != nil {
			continue // reported by TestEveryLessonIsSlides
		}
		for i, c := range cards {
			for _, problem := range budgetProblems(ref, c) {
				t.Errorf("lesson %q, card %d: %s", ref.ID(), i+1, problem)
			}
		}
	}
}

// budgetProblems lists every rule of the card budget that c breaks.
func budgetProblems(ref course.Ref, c card) []string {
	var problems []string
	if c.animations > 0 {
		if c.animations > 1 || len(c.others) > 0 || c.words > 0 {
			problems = append(problems, "an animation card holds only the animation and an optional <h2>; move the rest to another card")
		}
		return problems
	}
	if c.words > maxCardWords {
		problems = append(problems, fmt.Sprintf("%d words of prose, at most %d; split the card", c.words, maxCardWords))
	}
	if c.programs > maxCardPrograms {
		problems = append(problems, fmt.Sprintf("%d programs, at most %d; give each its own card", c.programs, maxCardPrograms))
	}
	if c.outputs > maxCardOutputs {
		problems = append(problems, fmt.Sprintf("%d output blocks, at most %d; give each its own card", c.outputs, maxCardOutputs))
	}
	for _, file := range c.examples {
		lines, err := countLines(path.Join(ref.Module.Slug, file))
		if err != nil {
			problems = append(problems, err.Error())
			continue
		}
		if lines > maxExampleLines && (c.words > 0 || c.outputs > 0) {
			problems = append(problems, fmt.Sprintf("%s has %d lines, over %d; it needs a card with only its heading, so move the prose and output to the next card", file, lines, maxExampleLines))
		}
	}
	return problems
}

func countLines(name string) (int, error) {
	src, err := fs.ReadFile(FS, name)
	if err != nil {
		return 0, fmt.Errorf("{{code}} file: %w", err)
	}
	return bytes.Count(src, []byte("\n")), nil
}

// parseCards reads a lesson body that must be made only of top-level
// <section class="slide"> elements, and describes each of them.
func parseCards(body []byte) ([]card, error) {
	dec := xml.NewDecoder(bytes.NewReader(body))
	dec.Strict = false
	dec.AutoClose = xml.HTMLAutoClose
	dec.Entity = xml.HTMLEntity

	var cards []card
	for {
		tok, err := dec.Token()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local != "section" || attr(tok, "class") != "slide" {
				return nil, fmt.Errorf(`top-level <%s>; wrap every card in <section class="slide">`, tok.Name.Local)
			}
			c, err := readCard(dec)
			if err != nil {
				return nil, fmt.Errorf("card %d: %w", len(cards)+1, err)
			}
			cards = append(cards, c)
		case xml.CharData:
			if text := strings.TrimSpace(string(tok)); text != "" {
				return nil, fmt.Errorf("text outside a card: %.40q", text)
			}
		}
	}
	if len(cards) == 0 {
		return nil, errors.New("no cards")
	}
	return cards, nil
}

// readCard consumes one <section class="slide"> up to its end tag.
func readCard(dec *xml.Decoder) (card, error) {
	var c card
	depth := 0     // nesting below the section
	skipping := "" // the heading, <pre> or animation whose text is not prose
	isGo := false  // the <pre> being read holds Go code
	for {
		tok, err := dec.Token()
		if err != nil {
			return c, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			name := tok.Name.Local
			if depth == 0 {
				switch {
				case isAnimation(name):
					c.animations++
				case name != "h2":
					c.others = append(c.others, name)
				}
			}
			if skipping == "" && (name == "h2" || name == "pre" || isAnimation(name)) {
				skipping, isGo = name, false
			}
			if skipping == "pre" && name == "code" && attr(tok, "class") == "language-go" {
				isGo = true
			}
			depth++
		case xml.EndElement:
			if depth == 0 {
				return c, nil // </section>
			}
			depth--
			if tok.Name.Local == skipping {
				if skipping == "pre" {
					if isGo {
						c.programs++
					} else {
						c.outputs++
					}
				}
				skipping = ""
			}
		case xml.CharData:
			text := string(tok)
			for _, m := range codeAction.FindAllStringSubmatch(text, -1) {
				c.examples = append(c.examples, m[1])
				c.programs++
			}
			if skipping == "" {
				c.words += len(strings.Fields(codeAction.ReplaceAllString(text, "")))
			}
		}
	}
}

// isAnimation reports whether an element is an animation component,
// such as <pointer-demo>.
func isAnimation(name string) bool {
	return strings.HasSuffix(name, "-demo")
}

// attr returns the value of an element's attribute, or "" if it has none.
func attr(el xml.StartElement, name string) string {
	for _, a := range el.Attr {
		if a.Name.Local == name {
			return a.Value
		}
	}
	return ""
}
