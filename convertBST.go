package lesson4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//ConvertBST:把二叉搜索树转换为累加树
/*parameter
root: 二叉搜索树根节点
return: 累加树根节点
*/
func ConvertBST(root *TreeNode) *TreeNode {

	sum := 0

	//先右后左中序遍历二叉搜索树
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {

		if node != nil {

			//遍历右子树
			dfs(root.Right)
			//累加右子树的值
			sum += node.Val
			//将累加值赋予当前节点
			node.Val = sum
			//遍历左子树
			dfs(node.Left)
		}
	}

	dfs(root)

	return root
}
