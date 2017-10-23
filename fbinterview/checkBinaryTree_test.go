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
	nodeList := genListOfNodes(2)
	nodeList[0].left = nodeList[1]
	assert.True(t, IsValidTree(nodeList), "Yep, linked list")

}

func genListOfNodes(n int) (nl NodeList) {

	nl = make(NodeList, n)

	for i := 0; i < n; i++ {
		nl[i] = &Node{}
	}

	return
}
