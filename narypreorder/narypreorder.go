package narypreorder

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	ret := []int{root.Val}
	for _, child := range root.Children {
		ret = append(ret, preorder(child)...)
	}

	return ret
}

func preorder_iter(root *Node) []int {
	if root == nil {
		return nil
	}

	//max input node count = 10000
	stack := make([]*Node, 10000)
	var ret []int
	sp, iret := 0, 0

	stack[sp] = root
	sp++
	for sp > 0 {
		sp--
		c := stack[sp]
		ret = append(ret, c.Val)
		iret++

		for i := len(c.Children) - 1; i >= 0; i-- {
			stack[sp] = c.Children[i]
			sp++
		}
	}

	return ret
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type NodeTrack struct {
	ParentNode *Node
	ChildIndex int
}

type NodeStack struct {
	stack []*NodeTrack
}

func MakeStack() *NodeStack {
	return &NodeStack{
		stack: make([]*NodeTrack, 0),
	}
}

func (s *NodeStack) Push(node *NodeTrack) {
	s.stack = append(s.stack, node)
}

func (s *NodeStack) Pop() *NodeTrack {
	if len(s.stack) == 0 {
		return nil
	}

	stackLen := len(s.stack)
	val := s.stack[stackLen-1]
	s.stack = s.stack[:stackLen-1]
	return val
}

func preorder_leet0ms(root *Node) []int {
	if root == nil {
		return nil
	}

	res := []int{root.Val}

	stack := MakeStack()
	current := root
	childIndex := 0

	for {
		for i := childIndex; i < len(current.Children); i++ {
			child := current.Children[i]
			res = append(res, child.Val)

			if child.Children != nil {
				stack.Push(&NodeTrack{ParentNode: current, ChildIndex: i})
				current = child
				i = -1
			}
		}

		prevLevel := stack.Pop()

		if prevLevel == nil {
			break
		}

		current = prevLevel.ParentNode
		childIndex = prevLevel.ChildIndex + 1
	}

	return res
}
