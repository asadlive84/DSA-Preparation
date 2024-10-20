package main

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    s1 := []*TreeNode{p}
    s2 := []*TreeNode{q}

    for len(s1) > 0 {
        n1 := s1[len(s1)-1]
        s1 = s1[:len(s1)-1]

        n2 := s2[len(s2)-1]
        s2 = s2[:len(s2)-1]

        if n1 == nil && n2 == nil {
            continue
        }
        if n1 == nil || n2 == nil || n1.Val != n2.Val {
            return false
        }

        s1 = append(s1, n1.Right, n1.Left)
        s2 = append(s2, n2.Right, n2.Left)
    }

    return true
}

func main() {
    // Creating Tree 1
    t1 := &TreeNode{Val: 1}
    t1.Left = &TreeNode{Val: 2}
    t1.Right = &TreeNode{Val: 3}

    // Creating Tree 2
    t2 := &TreeNode{Val: 1}
    t2.Left = &TreeNode{Val: 2}
    t2.Right = &TreeNode{Val: 3}

    // Checking if both trees are identical
    result := isSameTree(t1, t2)
    fmt.Println("Are the trees identical?", result) // Output: true
}
