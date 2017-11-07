package fbinterview

import "testing"

func TestLongestSubstring(t *testing.T) {
	for _, tc := range []struct {
		Str       string
		UniqChars int
		Result    string
	}{
		{"aaaa", 1, "aaaa"},
		{"aabbb", 1, "bbb"},
		{"aaaabbbb", 2, "aaaabbbb"},
		{"asdfrttt", 3, "frttt"},
		{"", 3, ""},
	} {

		if res := LongestSubstring(tc.Str, tc.UniqChars); res != tc.Result {

			t.Fatalf("input: %s %d , expected: %s,  got: %s", tc.Str, tc.UniqChars, tc.Result, res)

		}

	}
}
