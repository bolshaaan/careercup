// https://www.careercup.com/question?id=5183214233583616
//node {
//node * left, * right;
//}
//Given a list of node to determine whether the node in the list can form a valid binary tree. several points of judgment
//1. need to ensure that all left, right pointer point to the node inside list
//2. no cycle
//3. All node must be connected
//Boolean isValidTree(List<Node> list){}
//

package fbinterview

import "fmt"

type Node struct {
	left, right *Node
}

type NodeList []*Node
type MarkedNodes map[*Node]bool

func traceTree(node *Node, visitedNodes MarkedNodes) bool {
	if marked, ok := visitedNodes[node]; ok && marked {
		// we have already passed -- loop detected
		return false
	}

	var result bool = true
	visitedNodes[node] = true

	if node.left != nil {
		result = result && traceTree(node.left, visitedNodes)
	}

	if node.right != nil {
		result = result && traceTree(node.right, visitedNodes)
	}

	return result
}

// no cycle
// all nodes are connected
// all pointers are in the list
func IsValidTree(nl NodeList) bool {

	// empty binary tree
	if len(nl) == 0 {
		return true
	}

	// find node, that have no parents
	// make map nodes ( pointer )
	nodes := make(MarkedNodes)

	// fill map nodes
	for _, v := range nl {
		if v.left != nil {
			nodes[v.left] = true
		}
		if v.right != nil {
			nodes[v.right] = true
		}
	}

	// find root
	var root *Node
	for _, v := range nl {
		if _, ok := nodes[v]; !ok {
			if root != nil {
				fmt.Println("more then 1 root in struct, return false")
				return false
			}
			root = v
		}
	}

	if root == nil {
		fmt.Println("No root found !!!!")
		return false
	}

	visitedNodes := make(MarkedNodes)
	result := traceTree(root, visitedNodes)

	if result && len(visitedNodes) < len(nodes) {
		result = false
	}

	return result
}
