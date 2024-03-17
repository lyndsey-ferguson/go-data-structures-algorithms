package trees

import "github.com/lyndsey-ferguson/go-data-structures-algorithms/linkedlists"

/*
Implement a function to check if a binary tree is balanced. For the purpose of
this question, a balanced tree is defined to be a tree such that the heights of
the two subtrees of any node never differ by more than one.

input tree:

(13)

output list:

[1]

input tree:

     (25)
    / 
(13)
    \
	 (9)

output list:

[0] - [13]
[1] - [9]->[25]

input tree:

         (29)
        /  
     (25)
     /  \
    /    (21)
(13)
    \
     \
	 (9)

output list:

[0] - [13]
[1] - [9]->[25]
[2] - [21]->[29]


*/
func ListOfDepths[T comparable](root *Node[T], list []*linkedlists.Node[T], depth int) []*linkedlists.Node[T] {
	if root == nil {
		return list
	}
	n := linkedlists.CreateNode(root.data)
	if len(list) <= depth {
		list = append(list, n)
	} else {
		if list[depth] == nil {
			list[depth] = n
		} else {
			linkedlists.AppendNode(list[depth], n)
		}
	}
	list = ListOfDepths(root.left, list, depth+1)
	list = ListOfDepths(root.right, list, depth+1)

	return list
}
