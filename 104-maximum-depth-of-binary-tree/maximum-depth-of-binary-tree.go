/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	levelCount:=0
    for len(queue)>0{
        levelSize:= len(queue)
        for i:=0; i<levelSize; i++{
               
            current:= queue[0]
            queue = queue[1:]

            if current.Left!=nil{
                queue = append(queue, current.Left)
            }

             if current.Right!=nil{
                queue = append(queue, current.Right)
            }

        }
        levelCount++
       
    }
return levelCount

	
}