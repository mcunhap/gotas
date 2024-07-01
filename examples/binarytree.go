package examples

import (
	"fmt"
	"gotas/trees/binarytree"
)

func BinaryTreeExample() {
	//        1
	//     /	  \
	//    2	     3
	//   / \    / \
	//  4  5   6  7

	root := binarytree.NewNode[int](1)

	two, _ := root.InsertLeft(2)
	three, _ := root.InsertRight(3)
	two.InsertLeft(4)
	two.InsertRight(5)
	three.InsertLeft(6)
	three.InsertRight(7)

	fmt.Printf("Searching value 5: %t\n", root.Search(5))
	fmt.Printf("Searching value 10: %t\n", root.Search(10))

	fmt.Printf("InOrder traversal: %v\n", root.InOrder())
	fmt.Printf("PreOrder traversal: %v\n", root.PreOrder())
}
