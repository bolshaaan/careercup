package fbinterview

import (
	"testing"
)

func TestPrintSentence(t *testing.T) {

	for _, tc := range []struct {
		text   string
		kw     int
		result string
	}{
		{"H", 26, "R8, T"},
		{"HI", 26, "R8, T, R1, T"},
		{"HELLO", 13, "R8, T, L3, T, R7, T, T, D1, L10, T"},
	} {
		if res := PrintSentence(tc.text, tc.kw); res != tc.result {
			t.Fatalf("in: %s %d, exp: %s recv: %s ", tc.text, tc.kw, tc.result, res)
		}
	}
}
