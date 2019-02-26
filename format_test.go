package strf

import "testing"

func TestFormat(t *testing.T) {
	got := Format("{0}{1}")
	if got != "{0}{1}" {
		t.Error()
	}

	got = Format("", 1, 2, 3)
	if got != "" {
		t.Error()
	}

	got = Format("haha", 1, 2, 3)
	if got != "haha" {
		t.Error()
	}

	got = Format("ha{0}ha{1}{0}", 1, 2, 3)
	if got != "ha1ha21" {
		t.Error()
	}

	got = Format("{{{0}}}}{}", 1, 2, 3)
	if got != "{{1}}}{}" {
		t.Error()
	}

	got = Format("🦊{1}🐱中文{2}", 1, 2, 3)
	if got != "🦊2🐱中文3" {
		t.Error()
	}
}
