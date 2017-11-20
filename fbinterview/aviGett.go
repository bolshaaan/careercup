package fbinterview

import (
	"fmt"
)

type NodeStr struct {
	Next *NodeStr
	Prev *NodeStr
	Val  int32
}

func (str *NodeStr) Print() {
	fmt.Println("IN PRINT: ", str.toSimpleStr())
}

func (str *NodeStr) toSimpleStr() string {
	simpleStr := ""

	cur := str
	for cur != nil {
		simpleStr += string(cur.Val)
		cur = cur.Next
	}

	return simpleStr
}

func (str *NodeStr) FastStringInsertString(pos int, substr string) *NodeStr {
	if pos < 0 {
		return nil
	}

	if len(substr) == 0 {
		return str
	}

	cur := str

	for cur.Next != nil && pos > 0 {
		cur = cur.Next
		pos--
	}

	switch {
	case cur.Next == nil && pos == 1:
		// insert to the end
		cur.Next = NewFastString(substr)
		return str
	case cur.Next == nil:
		// very big position, to think, if can error insert?
		return nil
	}

	// Debug
	//fmt.Printf("Char at position %s\n", string(cur.Val))

	fastSubstr := NewFastString(substr)

	stored := cur
	//fmt.Println("Stored: ", string(stored.Val))
	start := str

	// cur.Prev can be nil - Must do something
	if cur.Prev == nil {
		cur = fastSubstr
		start = fastSubstr
	} else {
		cur.Prev.Next = fastSubstr
		fastSubstr.Prev = cur.Prev
		cur = cur.Prev.Next
	}

	for cur.Next != nil {
		cur = cur.Next
	}

	// Debug
	//fmt.Printf("Char at stored next %s\n", string(stored.Next.Val))
	//fmt.Printf("Char at cur  %s\n", string(cur.Val))

	cur.Next = stored

	str = start

	return start
}

func NewFastString(tmpStr string) *NodeStr {
	node := &NodeStr{}
	head := node

	for i, b := range tmpStr {
		node.Val = b

		if i < len(tmpStr)-1 {
			node.Next = &NodeStr{}
			node.Next.Prev = node
		}

		node = node.Next //go further
	}

	return head
}

func (str *NodeStr) FastStringInsert(pos int, substr *NodeStr) *NodeStr {

	return nil
}
