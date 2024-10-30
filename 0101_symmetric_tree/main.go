package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	return check(root.Left, root.Right)
}

func check(l *TreeNode, r *TreeNode) bool {

	if (l == nil) || (r == nil) {
		return l == nil && r == nil
	}
	if l.Val != r.Val {
		return false
	}

	a1 := check(l.Left, r.Right)
	a2 := check(l.Right, r.Left)

	return a1 && a2
}

func main() {
	// root := &TreeNode{
	// 	Val: 1,
	// 	Left: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val:   3,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Right: &TreeNode{
	// 			Val:   4,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 2,
	// 		Left: &TreeNode{
	// 			Val:   4,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 		Right: &TreeNode{
	// 			Val:   3,
	// 			Left:  nil,
	// 			Right: nil,
	// 		},
	// 	},
	// }
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}

	println(isSymmetric(root))
}
