package strf

import "testing"

func TestFormat(t *testing.T) {
	got := Format("", 1, 2, 3)
	if got != "" {
		t.Error(got)
	}

	got = Format("haha", 1, 2, 3)
	if got != "haha" {
		t.Error(got)
	}

	got = Format("ha{0}ha{1}{0}", 1, 2, 3)
	if got != "ha1ha21" {
		t.Error(got)
	}

	got = Format("{{{0}}}}{}", 1, 2, 3)
	if got != "{{1}}}{}" {
		t.Error(got)
	}

	got = Format("🦊{1}🐱中文{2}", 1, 2, 3)
	if got != "🦊2🐱中文3" {
		t.Error(got)
	}
}

func TestFormatCore(t *testing.T) {
	got := FormatCore("{}{}")
	if got != "%!v(MISSING)%!v(MISSING)" {
		t.Error(got)
	}

	got = FormatCore("", 1, 2, 3)
	if got != "%!(EXTRA int=1, int=2, int=3)" {
		t.Error(got)
	}

	got = FormatCore("haha", 1, 2, 3)
	if got != "haha%!(EXTRA int=1, int=2, int=3)" {
		t.Error(got)
	}

	got = FormatCore("ha{}ha{}{}", 1, 2, 3)
	if got != "ha1ha23" {
		t.Error(got)
	}

	got = FormatCore("{{{0}}}}{}", 1)
	if got != "{{{0}}}}1" {
		t.Error(got)
	}

	got = FormatCore("🦊{}🐱中文{}", 1, 2)
	if got != "🦊1🐱中文2" {
		t.Error(got)
	}
}
