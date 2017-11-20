package fbinterview

import (
	"testing"

	"os"

	"log"

	"runtime/pprof"

	"github.com/magiconair/properties/assert"
)

func TestT9Autocorrect(t *testing.T) {

	f, err := os.Create("/Users/aleksandr/memprof.mprof")
	defer f.Close()

	if err != nil {
		log.Fatal("Something: ", err)
	}

	for _, tc := range []struct {
		dictWordsCase []string
		numbers       []int
		result        string
	}{

		{[]string{"aaa"}, []int{2, 2, 2}, "aaa"},
		{[]string{"dab", "why"}, []int{3, 2, 2}, "dab"},
		{[]string{"tapping", "tapping", "tasting"}, []int{8, 2, 7, 7, 4, 6, 4}, "tapping"},
		{[]string{"testing", "tapping", "tasting"}, []int{8, 3, 7, 8, 4, 6, 4}, "testing"},
		{[]string{"testing", "tapping", "tasting"}, []int{8, 2, 7, 8, 4, 6, 4}, "tasting"},
	} {
		dictWords = tc.dictWordsCase
		trieTree = MakeTrie()

		TraverseTreeDump(trieTree)

		res := T9Autocorrect(tc.numbers)
		assert.Equal(t, res, tc.result)
	}

	pprof.WriteHeapProfile(f)
}
