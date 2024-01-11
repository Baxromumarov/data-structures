package main

import "fmt"

/*

type BinaryTreeNode struct {
	Val   int64
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (t *BinaryTreeNode) print() {
	if t != nil {
		// Traverse the left subtree
		t.Left.print()

		// Print the value of the current node
		fmt.Println(t.Val)

		// Traverse the right subtree
		t.Right.print()
	}
}

func (t *BinaryTreeNode) findElement(val int64) bool {
	if t == nil {
		return false
	}

	if t.Val == val {
		return true
	}
	return t.Right.findElement(val) || t.Left.findElement(val)
}

func (t *BinaryTreeNode) DeleteElement(val int64) {
	if t == nil {
		return
	}
	if val == t.Val {

	}
	if val < t.Val {
		t.Left.DeleteElement(val)
	} else if val > t.Val {
		t.Right.DeleteElement(val)
	}

}

func (t *BinaryTreeNode) insert(data int64) {
	if t == nil {
		t = &BinaryTreeNode{Val: data}
	} else if data <= t.Val {
		if t.Left == nil {
			t.Left = &BinaryTreeNode{Val: data, Left: nil, Right: nil}
		} else {
			t.Left.insert(data)
		}
	} else {
		if t.Right == nil {
			t.Right = &BinaryTreeNode{Val: data, Left: nil, Right: nil}
		} else {
			t.Right.insert(data)
		}
	}
	// fmt.Println(t)
}

func DFS(root *TreeNode) {
	if root == nil {
		return
	}

	// Pre-order traversal: process the current node before its children
	fmt.Printf("%d ", root.Value)

	DFS(root.Left)
	// In-order traversal: process the current node in between its left and right children
	// fmt.Printf("%d ", root.Value)
	DFS(root.Right)
	// Post-order traversal: process the current node after its children
	// fmt.Printf("%d ", root.Value)
}
func main() {
	var newTree = &BinaryTreeNode{}
	newTree.insert(10)
	newTree.insert(15)
	newTree.insert(4)
	newTree.insert(12)
	newTree.insert(8)
	newTree.insert(21)
	newTree.insert(5)

	dfs(newTree)
	// newTree.print()

	// fmt.Println(newTree.findElement(13))

}
*/

type Node struct {
	Val   int64
	Left  *Node
	Right *Node
}

// Inorder Traversal (Left-Root-Right)
// Preorder Traversal (Root-Left-Right)
// Postorder Traversal (Left-Right-Root)
// link traversal: https://en.wikipedia.org/wiki/Tree_traversal
// 	 1
// 	/  \
// 	2   3
// / \
// 4  5
// Therefore, the Depth First Traversals of this Tree will be:

// 1.Inorder: 4 2 5 1 3
// 2.Preorder: 1 2 4 5 3
// 3.Postorder: 4 5 2 3 1

// 1.Inorder: 4 2 5 1 3
// Time Complexity: O(N)
// Auxiliary Space: O(log N)
func printInOrder(node *Node) {
	if node == nil {
		return

	}

	printInOrder(node.Left)

	/* then print the data of node */
	fmt.Print(node.Val, " ")

	/* now recur on right child */
	printInOrder(node.Right)

}

// 2.Preorder: 1 2 4 5 3
// Time Complexity: O(N)
// Auxiliary Space: O(log N)
func printPreOrder(node *Node) {
	if node == nil {
		return

	}

	fmt.Print(node.Val, " ")
	printPreOrder(node.Left)
	printPreOrder(node.Right)

}

// 3.Postorder: 4 5 2 3 1
// Time Complexity: O(N)
// Auxiliary Space: O(log N)
func printPostOrder(node *Node) {
	if node == nil {
		return

	}

	printPostOrder(node.Left)
	printPostOrder(node.Right)
	fmt.Print(node.Val, " ")

}

func DFS(root *Node) {
	if root == nil {
		return
	}

	// Pre-order traversal: process the current node before its children
	fmt.Printf("%d ", root.Val)

	DFS(root.Left)
	// In-order traversal: process the current node in between its left and right children
	// fmt.Printf("%d ", root.Value)
	DFS(root.Right)
	// Post-order traversal: process the current node after its children
	// fmt.Printf("%d ", root.Value)
}

func BFS(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		// Dequeue a node
		node := queue[0]
		queue = queue[1:]

		// Process the node
		fmt.Printf("%d ", node.Val)

		// Enqueue the left child
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		// Enqueue the right child
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
	node := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
		},
	}
	// printPostOrder(node)
	// DFS(node)
	BFS(node)
}
