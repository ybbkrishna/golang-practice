package interviewPrep

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	output := [][]int{}
	if root == nil {
		return output
	}
	tempQ := []*TreeNode{}
	tempQ = append(tempQ, root, nil)
	level := 0
	for len(tempQ) != 0 {
		// pop node from the queue
		node := tempQ[0]
		tempQ = tempQ[1:]

		// breaking condition
		if node == nil && len(tempQ) == 0 {
			break
		}

		// level check
		if node == nil {
			level++
			tempQ = append(tempQ, nil)
			continue
		}

		if node.Left != nil {
			tempQ = append(tempQ, node.Left)
		}
		if node.Right != nil {
			tempQ = append(tempQ, node.Right)
		}

		if len(output) == level {
			output = append(output, []int{})
		}
		output[level] = append(output[level], node.Val)
	}
	return output
}
