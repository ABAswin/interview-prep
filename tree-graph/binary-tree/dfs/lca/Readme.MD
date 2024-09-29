Learnings from this excercise


The idea here is to return non nil nodes from the left , right and determine the LCA.

A beautiful idea , incase if found in very iterations the address of 2 that node is returned to the previous levels if already determined. (case of 7,4)

At the end incase of 0,6 there has to be a place where the left and right subtree returns non nil values which happens to be the maximum root of the binary tree. 