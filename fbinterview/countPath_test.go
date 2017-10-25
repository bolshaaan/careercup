package fbinterview

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// tests 2 realisations
func TestFindPath(t *testing.T) {

	for _, test := range []struct{ M, N, X, Y, K int }{
		{20, 20, 1, 1, 2},
		{20, 10, 1, 1, 2},
		{20, 10, 1, 1, 4},

		{20, 10, 1, 2, 4},
		{20, 20, 10, 10, 4},
	} {

		task := &Task{M: test.M, N: test.N, K: test.K}
		bf := task.FindPathBruteForce(test.X, test.Y, 0)
		pl := task.FindPathPoly(test.X, test.Y, 0)

		assert.Equal(t, bf, pl)
	}

}
