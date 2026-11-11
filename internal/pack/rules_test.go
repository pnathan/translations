package pack

import (
	"strings"
	"testing"
)

const sample = `# T

Preamble kept.

## Output format for this work (prose)

Always kept.

## Terms

- **ספירה / ספירות**: fixed.
  - sub entry
- **אין סוף**: fixed.
- A general rule with no Hebrew head: kept.
- רשב"י → "Rabbi Shimon": kept when present.

## Pending consistency pass

- 2:18 something.
`

func TestRules(t *testing.T) {
	got := Rules(sample, "והספירות העליונות ורשב״י")
	for _, want := range []string{"Preamble kept", "Always kept", "ספירה / ספירות", "sub entry", "A general rule", "Rabbi Shimon"} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in\n%s", want, got)
		}
	}
	for _, bad := range []string{"אין סוף", "2:18"} {
		if strings.Contains(got, bad) {
			t.Errorf("unexpected %q", bad)
		}
	}
	if !strings.Contains(Rules(sample, "אין סוף"), "אין סוף") {
		t.Error("אין סוף not kept")
	}
}
