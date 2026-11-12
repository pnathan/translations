package pack

import (
	"regexp"
	"strings"
	"unicode"
)

// Rules extracts from a work's decisions.md the rulings one unit needs, so
// that a model stage reads a few dozen lines instead of the whole file.
//
// The file is split into sections ("## ") and, inside a section, into
// entries (a line starting "- " with the lines that follow it up to the next
// such line). The preamble and the sections in alwaysSections are kept
// whole. An entry of any other section is kept when its head (the text up to
// the first ": " or " → ", and a bold title if it has one) has no Hebrew, or
// when a Hebrew word of the head occurs in the unit's text. The sections in
// skipSections are dropped. The full file stays the authority.
func Rules(decisions, unitText string) string {
	words := unitWords(unitText)
	var out strings.Builder
	var sec *ruleSection
	flush := func() {
		if sec == nil {
			return
		}
		body := sec.render(words)
		if strings.TrimSpace(body) != "" {
			out.WriteString(body)
		}
	}
	sec = &ruleSection{always: true}
	for _, line := range strings.Split(decisions, "\n") {
		if h, ok := strings.CutPrefix(line, "## "); ok {
			flush()
			sec = &ruleSection{heading: line, always: hasPrefixAny(h, alwaysSections), skip: hasPrefixAny(h, skipSections)}
			continue
		}
		sec.add(line)
	}
	flush()
	return out.String()
}

var (
	alwaysSections = []string{"Output format", "Style", "Translator's checklist", "Owner's answers"}
	skipSections   = []string{"Pending consistency"}
)

func hasPrefixAny(s string, ps []string) bool {
	for _, p := range ps {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}

type ruleSection struct {
	heading string
	always  bool
	skip    bool
	intro   []string
	entries [][]string
}

func (s *ruleSection) add(line string) {
	if strings.HasPrefix(line, "- ") || len(s.entries) == 0 && strings.HasPrefix(line, "-") {
		s.entries = append(s.entries, []string{line})
		return
	}
	if len(s.entries) == 0 {
		s.intro = append(s.intro, line)
		return
	}
	last := len(s.entries) - 1
	s.entries[last] = append(s.entries[last], line)
}

func (s *ruleSection) render(words map[string]bool) string {
	if s.skip {
		return ""
	}
	var b strings.Builder
	if s.always {
		if s.heading != "" {
			b.WriteString(s.heading + "\n")
		}
		b.WriteString(strings.Join(s.intro, "\n"))
		for _, e := range s.entries {
			b.WriteString("\n" + strings.Join(e, "\n"))
		}
		b.WriteString("\n")
		return b.String()
	}
	kept := 0
	for _, e := range s.entries {
		if !keep(e[0], words) {
			continue
		}
		if kept == 0 && s.heading != "" {
			b.WriteString(s.heading + "\n\n")
		}
		kept++
		b.WriteString(strings.Join(e, "\n") + "\n")
	}
	return b.String()
}

var reBold = regexp.MustCompile(`\*\*([^*]+)\*\*`)

// head is the part of an entry's first line that names what it rules on.
func head(first string) string {
	first = strings.TrimPrefix(first, "- ")
	if m := reBold.FindStringSubmatch(first); m != nil {
		return m[1]
	}
	for _, sep := range []string{": ", " → ", " = ", " (", ", "} {
		if i := strings.Index(first, sep); i > 0 {
			return first[:i]
		}
	}
	return first
}

func keep(first string, words map[string]bool) bool {
	hw := hebrewKeys(head(first))
	if len(hw) == 0 {
		return true
	}
	for _, k := range hw {
		if matches(k, words) {
			return true
		}
	}
	return false
}

func isHeb(r rune) bool { return unicode.Is(unicode.Hebrew, r) && !(r >= 0x0591 && r <= 0x05C7) }

// hebrewKeys lists the Hebrew words of s with quotes, geresh and gershayim
// removed ("רשב"י" and "רשב״י" both give "רשבי").
func hebrewKeys(s string) []string {
	var out []string
	var cur []rune
	for _, r := range s + " " {
		switch {
		case isHeb(r) && r != '׳' && r != '״':
			cur = append(cur, r)
		case r == '"' || r == '\'' || r == '׳' || r == '״' || r == '’' || r == '”':
			// inside an abbreviation: skip
		default:
			if len(cur) >= 2 {
				out = append(out, string(cur))
			}
			cur = nil
		}
	}
	return out
}

func unitWords(text string) map[string]bool {
	m := map[string]bool{}
	for _, w := range hebrewKeys(text) {
		m[w] = true
	}
	return m
}

const prefixes = "והבכלמש"

// matches reports whether key occurs in the unit: as a word, with up to two
// prefix letters before it, or (for keys of three letters or more) inside a word.
func matches(key string, words map[string]bool) bool {
	if words[key] {
		return true
	}
	r := []rune(key)
	for w := range words {
		wr := []rune(w)
		if len(r) >= 3 && strings.Contains(w, key) {
			return true
		}
		for n := 1; n <= 2 && n < len(wr); n++ {
			if string(wr[n:]) == key && strings.ContainsRune(prefixes, wr[0]) && (n == 1 || strings.ContainsRune(prefixes, wr[1])) {
				return true
			}
		}
	}
	return false
}
