package fbinterview

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
)

func TestIsValidTree(t *testing.T) {
	assert.True(t, IsValidTree(nil), "List Is Nil")

	var zeroList = make(NodeList, 0)
	assert.True(t, IsValidTree(zeroList), "List Is Zero List")

	// list 1 node
	oneNodeList := NodeList{&Node{left: nil, right: nil}}
	assert.True(t, IsValidTree(oneNodeList), "One Node List")

	// list 2 nodes
	nl := genListOfNodes(2)
	nl[0].left = nl[1]
	assert.True(t, IsValidTree(nl), "Yep, linked list")

	// false - cycle
	nl[1].right = nl[0]
	assert.False(t, IsValidTree(nl), "Nope")

	// 3 nodes, is cycled
	nl = genListOfNodes(3)
	nl[0].left = nl[1]
	nl[1].right = nl[2]
	nl[2].left = nl[1]

	// remove cycle
	nl[2].left = nil
	assert.True(t, IsValidTree(nl), "No cycles")

	// one node not connected
	nl[1].right = nil
	assert.False(t, IsValidTree(nl), "One node not connected")
}

func TestSplit(t *testing.T) {

	var tests = []struct {
		sep  string
		text string
		exp  int
	}{
		{sep: ":", text: "a:b:c", exp: 3},
	}

	for _, test := range tests {

		res := strings.Split(test.text, test.sep)
		if len(res) != test.exp {
			t.Errorf("Some shit: expected = %d, got = %s", len(res), test.exp)
		}
	}

}

func genListOfNodes(n int) (nl NodeList) {

	nl = make(NodeList, n)

	for i := 0; i < n; i++ {
		nl[i] = &Node{}
	}

	return
}
