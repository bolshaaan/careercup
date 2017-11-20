package fbinterview

import (
	"fmt"

	"go4.org/sort"
)

type NodeT struct {
	Siblings map[int]*NodeT
	Words    []string
}

var trieTree *NodeT

var dictWords = []string{
	"hello", "bugiboy", "armen", "kocharov",
}

var NumMap map[int]int = MakeNumMap()

// 26 letters
func MakeNumMap() map[int]int {

	res := make(map[int]int)

	// 2 empty strings --> offset for 2
	temp := []string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}

	for i := 2; i <= 9; i++ {
		for _, b := range temp[i] {
			res[int(b)] = i
		}
	}

	return res
}

func MakeTrie() *NodeT {

	start := new(NodeT)

	// every word
	for _, w := range dictWords {

		cur := start // search from start

		// every byte in word
		for _, b := range w {

			// find elem  in Node
			if cur.Siblings == nil {
				cur.Siblings = make(map[int]*NodeT)
			}

			bb := NumMap[int(b)]

			if _, ok := cur.Siblings[bb]; !ok {
				cur.Siblings[bb] = new(NodeT)
			}

			cur = cur.Siblings[bb]

		}

		// add word to result node
		if cur.Words == nil {
			cur.Words = []string{}
		}

		cur.Words = append(cur.Words, w)
		sort.Strings(cur.Words)
	}

	return start
}

func init() {
	trieTree = MakeTrie()
}

func T9Autocorrect(taps []int) string {
	cur := trieTree

	for _, tap := range taps {
		if cur.Siblings == nil {
			return "no siblings"
		}

		if _, ok := cur.Siblings[tap]; !ok {
			return "have siblings, no tap"
		}

		cur = cur.Siblings[tap]
	}

	if cur.Words == nil {
		return "no words"
	}

	return cur.Words[0]
}

func TraverseTreeDump(node *NodeT) {
	cur := node

	if cur.Words != nil {
		fmt.Printf(" -> Words == %v", cur.Words)
	}

	if cur.Siblings != nil {
		for k, v := range cur.Siblings {
			fmt.Printf(" -> %d", k)
			TraverseTreeDump(v)
		}
	}

}
