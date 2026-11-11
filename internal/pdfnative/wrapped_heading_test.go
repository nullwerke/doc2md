package pdfnative

import "testing"

// A heading the page wrapped is one heading: a hyphen-split word joins
// without the hyphen, a lowercase continuation with a space; a real next
// heading stays.
func TestMergeWrappedHeadings(t *testing.T) {
	in := []block{
		{kind: kindHeading, level: 1, text: "8 Organisation und Vorgaben zum Problem- und Ände-"},
		{kind: kindHeading, level: 1, text: "rungsmanagement"},
		{kind: kindHeading, level: 2, text: "8.1 Verlauf des Änderungsverfahrens"},
		{kind: kindHeading, level: 1, text: "9 Organisation und Vorgaben zum"},
		{kind: kindHeading, level: 1, text: "konfigurationsbezogenen Management"},
		{kind: kindHeading, level: 1, text: "10 Anforderungen"},
	}
	got := mergeWrappedHeadings(in)
	want := []string{"8 Organisation und Vorgaben zum Problem- und Änderungsmanagement", "8.1 Verlauf des Änderungsverfahrens",
		"9 Organisation und Vorgaben zum konfigurationsbezogenen Management", "10 Anforderungen"}
	if len(got) != len(want) {
		t.Fatalf("%d Überschriften, want %d: %v", len(got), len(want), got)
	}
	for i, w := range want {
		if got[i].text != w {
			t.Errorf("%d: %q, want %q", i, got[i].text, w)
		}
	}
}
