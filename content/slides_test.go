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
	"strconv"
	"strings"
	"testing"

	"paseo-en-go/internal/course"
)

// The budget of one card, a testable stand-in for "fits a 1280×720 screen".
const (
	maxCardWords    = 90 // prose words, headings and code blocks excluded
	maxCardPrograms = 1  // {{code}} snippets and Go code blocks
	maxCardOutputs  = 1  // plain <pre> blocks: a program's output or an error
	// A card with a heading has room for 17 lines of code at 1280×720; a
	// longer program is shown across cards with {{codeLines}}.
	maxSnippetLines = 17
	// A program and its output share a card only while both are this short.
	maxLinesWithOutput = 13
)

// codeAction matches {{code "examples/x/main.go"}} and
// {{codeLines "examples/x/main.go" 1 17}}, with the file and the line range.
var codeAction = regexp.MustCompile(`\{\{\s*code(?:Lines)?\s+"([^"]+)"(?:\s+(\d+)\s+(\d+))?\s*\}\}`)

// snippet is a program embedded with {{code}} or {{codeLines}}; first and
// last are 0 when the whole file is shown.
type snippet struct {
	file        string
	first, last int
}

// card is what the budget checks need to know about one slide.
type card struct {
	words       int       // prose words outside headings, code and animations
	programs    int       // Go code blocks, embedded or inline
	outputs     int       // plain <pre> blocks
	outputLines int       // lines inside those blocks
	animations  int       // <*-demo> elements
	others      []string  // elements beside an <h2> and an animation
	examples    []snippet // programs embedded with {{code}} or {{codeLines}}
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
	for _, ex := range c.examples {
		lines, err := countLines(path.Join(ref.Module.Slug, ex.file))
		if err != nil {
			problems = append(problems, err.Error())
			continue
		}
		if ex.last > 0 {
			lines = ex.last - ex.first + 1
		}
		switch {
		case lines > maxSnippetLines:
			problems = append(problems, fmt.Sprintf("%s shows %d lines, over %d; show it across cards with {{codeLines}}", ex.file, lines, maxSnippetLines))
		case lines+c.outputLines > maxLinesWithOutput && c.outputs > 0:
			problems = append(problems, fmt.Sprintf("%s and its output take %d lines, over %d; move the output to the next card", ex.file, lines+c.outputLines, maxLinesWithOutput))
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

// newSnippet reads a codeAction match.
func newSnippet(m []string) snippet {
	first, _ := strconv.Atoi(m[2])
	last, _ := strconv.Atoi(m[3])
	return snippet{file: m[1], first: first, last: last}
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
	preLines := 0  // lines of text in the <pre> being read
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
				skipping, isGo, preLines = name, false, 1
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
						c.outputLines += preLines
					}
				}
				skipping = ""
			}
		case xml.CharData:
			text := string(tok)
			for _, m := range codeAction.FindAllStringSubmatch(text, -1) {
				c.examples = append(c.examples, newSnippet(m))
				c.programs++
			}
			if skipping == "pre" {
				preLines += strings.Count(text, "\n")
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
