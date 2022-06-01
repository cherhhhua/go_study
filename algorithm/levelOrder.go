package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (res [][]int) {
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) < depth+1 {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], root.Val)
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}

	dfs(root, 0)
	return
}
