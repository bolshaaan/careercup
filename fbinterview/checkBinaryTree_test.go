package fbinterview

import (
	"testing"

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

}

func genListOfNodes(n int) (nl NodeList) {

	nl = make(NodeList, n)

	for i := 0; i < n; i++ {
		nl[i] = &Node{}
	}

	return
}
