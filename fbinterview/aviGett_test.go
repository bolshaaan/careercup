package fbinterview

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestNewFastString(t *testing.T) {

TEST:
	for _, tc := range []struct {
		str string
	}{
		{"Grumming"},
		{"L"},
		{"Very large sentence with big attitude to test this test case"},

		// empty test case
		{""},
	} {
		str := NewFastString(tc.str)

		cur := str
		for _, b := range tc.str {
			switch {
			case cur == nil:
				t.Errorf("Len is not equal!!!")
				continue TEST
			case b != cur.Val:
				t.Errorf("Symbols are not equal: exp: %s got: %s", string(b), string(cur.Val))
				continue TEST
			}

			cur = cur.Next
		}

	}
}

func TestNodeStr_FastStringInsert(t *testing.T) {
	for _, tc := range []struct {
		str, substr string
		pos         int
		res         string
	}{
		{"ABCDEF", "XXXX", 2, "ABXXXXCDEF"},

		//// at begin
		{"ABCDEF", "XXXX", 0, "XXXXABCDEF"},

		// at end
		{"ABCDEF", "XXXX", 6, "ABCDEFXXXX"},

		// check empty substring
		{"ABCDEF", "", 1, "ABCDEF"},

	} {

		//s.Print()

		newStr := NewFastString(tc.str).FastStringInsertString(tc.pos, tc.substr)

		assert.Equal(t, newStr.toSimpleStr(), tc.res)

		//newStr := s.FastStringInsertString(2, tc.substr)

		//DEBUG
		//newStr.Print()

	}

}
