package rbtree

import (
	"fmt"
)

type RBTreeNode struct {
	key    int64
	color  byte
	parent *RBTreeNode
	left   *RBTreeNode
	right  *RBTreeNode
	data   interface{}
}

type RBTree struct {
	root     *RBTreeNode
	sentinel *RBTreeNode
}

func (tree *RBTree) Init() {
	var sentinel RBTreeNode
	sentinel.color = 'b'
	tree.root = &sentinel
	tree.sentinel = &sentinel
}

func NewRBTree() *RBTree {
	var tree RBTree
	var sentinel RBTreeNode
	sentinel.color = 'b'
	tree.root = &sentinel
	tree.sentinel = &sentinel
	return &tree
}

func (tree *RBTree) Search(key int64) *RBTreeNode {
	curpos := tree.root
	for {
		if curpos == nil {
			return nil
		}

		if curpos.key == key {
			return curpos
		}

		if curpos.key > key {
			curpos = curpos.left
		} else {
			curpos = curpos.right
		}
	}
}

func (tree *RBTree) Insert(key int64, value interface{}) *RBTreeNode {
	var insertnode RBTreeNode
	var node = &insertnode
	node.key = key
	node.data = value
	node.left = tree.sentinel
	node.right = tree.sentinel
	node.color = 'r'

	if tree.root == tree.sentinel {
		node.parent = nil
		node.color = 'b'
		tree.root = node
		return node
	}

	curpos := tree.root
	for {
		if curpos.key > node.key {
			if curpos.left == tree.sentinel {
				curpos.left = node
				node.parent = curpos
				break
			} else {
				curpos = curpos.left
			}
		} else if curpos.key < node.key {
			if curpos.right == tree.sentinel {
				curpos.right = node
				node.parent = curpos
				break
			} else {
				curpos = curpos.right
			}
		} else {
			curpos.data = node.data
			return node
		}
	}

	var uncle *RBTreeNode
	for node != tree.root && node.color == 'r' && node.parent.color == 'r' {
		if node.parent == node.parent.parent.left {
			uncle = node.parent.parent.right
			if uncle.color == 'r' {
				node.parent.color = 'b'
				uncle.color = 'b'
				node.parent.parent.color = 'r'
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					tree.LeftRotate(node)
				}
				node.parent.color = 'b'
				node.parent.parent.color = 'r'
				tree.RightRotate(node.parent.parent)
			}
		} else {
			uncle = node.parent.parent.left
			if uncle.color == 'r' {
				node.parent.color = 'b'
				uncle.color = 'b'
				node.parent.parent.color = 'r'
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					tree.RightRotate(node)
				}
				node.parent.color = 'b'
				node.parent.parent.color = 'r'
				tree.LeftRotate(node.parent.parent)
			}
		}

	}

	tree.root.color = 'b'

	return node
}

func (tree *RBTree) Delete(key int64) {
	node := tree.Search(key)
	if node == nil {
		return
	}

	var temp *RBTreeNode
	var subst *RBTreeNode
	var w *RBTreeNode
	if node.left == tree.sentinel {
		temp = node.right
		subst = node
	} else if node.right == tree.sentinel {
		temp = node.right
		subst = node
	} else {
		subst = tree.Min()
		if subst.left != tree.sentinel {
			temp = subst.left
		} else {
			temp = subst.right
		}
	}

	if subst == tree.root {
		tree.root = temp
		temp.color = 'b'

		node.parent = nil
		node.left = nil
		node.right = nil

		return
	}

	color := subst.color

	if subst == subst.parent.left {
		subst.parent.left = temp
	} else {
		subst.parent.right = temp
	}

	if subst == node {
		temp.parent = subst.parent
	} else {
		if subst.parent == node {
			temp.parent = subst
		} else {
			temp.parent = subst.parent
		}

		subst.left = node.left
		subst.right = node.right
		subst.parent = node.parent
		subst.color = node.color

		if node == tree.root {
			tree.root = subst
		} else {
			if node == node.parent.left {
				node.parent.left = subst
			} else {
				node.parent.right = subst
			}
		}

		if subst.left != tree.sentinel {
			subst.left.parent = subst
		}
		if subst.right != tree.sentinel {
			subst.right.parent = subst
		}
	}

	node.parent = nil
	node.left = nil
	node.right = nil

	if color == 'r' {
		return
	}

	for temp != tree.root && temp.color == 'b' {
		if temp == temp.parent.left {
			w = temp.parent.right
			if w.color == 'r' {
				w.color = 'b'
				temp.parent.color = 'r'
				tree.LeftRotate(temp.parent)
				w = temp.parent.right
			}

			if w.left.color == 'b' && w.right.color == 'b' {
				w.color = 'r'
				temp = temp.parent
			} else {
				if w.right.color == 'b' {
					w.left.color = 'b'
					w.color = 'r'
					tree.RightRotate(w)
					w = temp.parent.right
				}

				w.color = temp.parent.color
				temp.parent.color = 'b'
				w.right.color = 'b'
				tree.LeftRotate(temp.parent)
				temp = tree.root
			}
		} else {
			w = temp.parent.left
			if w.color == 'r' {
				w.color = 'b'
				temp.parent.color = 'r'
				tree.RightRotate(temp.parent)
				w = temp.parent.left
			}

			if w.left.color == 'b' && w.right.color == 'b' {
				w.color = 'r'
				temp = temp.parent
			} else {
				if w.left.color == 'b' {
					w.right.color = 'b'
					w.color = 'r'
					tree.LeftRotate(w)
					w = temp.parent.left
				}

				w.color = temp.parent.color
				temp.parent.color = 'b'
				w.left.color = 'b'
				tree.RightRotate(temp.parent)
				temp = tree.root
			}
		}
	}

	temp.color = 'b'
}

func (tree *RBTree) Min() *RBTreeNode {
	curpos := tree.root
	for {
		if curpos.left != tree.sentinel {
			curpos = curpos.left
		} else {
			return curpos
		}
	}
}

func (tree *RBTree) Max() *RBTreeNode {
	curpos := tree.root
	for {
		if curpos.right != tree.sentinel {
			curpos = curpos.right
		} else {
			return curpos
		}
	}
}

func (tree *RBTree) IsEmpty() bool {
	if tree.root == tree.sentinel {
		return true
	} else {
		return false
	}
}

func (tree *RBTree) LeftRotate(node *RBTreeNode) {
	var temp *RBTreeNode

	temp = node.right
	node.right = temp.left

	if temp.left != tree.sentinel {
		temp.left.parent = node
	}

	temp.parent = node.parent

	if node == tree.root {
		tree.root = temp
	} else if node == node.parent.left {
		node.parent.left = temp
	} else {
		node.parent.right = temp
	}

	temp.left = node
	node.parent = temp
}

func (tree *RBTree) RightRotate(node *RBTreeNode) {
	var temp *RBTreeNode

	temp = node.left
	node.left = temp.right

	if temp.right != tree.sentinel {
		temp.right.parent = node
	}

	temp.parent = node.parent

	if node == tree.root {
		tree.root = temp
	} else if node == node.parent.right {
		node.parent.right = temp
	} else {
		node.parent.right = temp
	}

	temp.right = node
	node.parent = temp
}

func (tree *RBTree) TreeDeep() int {
	return tree.Deep(tree.root)
}

func (tree *RBTree) Deep(node *RBTreeNode) int {
	var getdeep func(node *RBTreeNode) int
	getdeep = func(node *RBTreeNode) int {
		if node == nil {
			return 0
		}
		if node.left == nil && node.right == nil {
			return 1
		}

		ldeep := getdeep(node.left)
		rdeep := getdeep(node.right)
		if ldeep > rdeep {
			return ldeep + 1
		} else {
			return rdeep + 1
		}
	}

	return getdeep(node) - 1
}

func (tree *RBTree) PrintTree() {
	var printnode func(node *RBTreeNode, indent int, flag byte)
	printnode = func(node *RBTreeNode, indent int, flag byte) {
		if node != tree.sentinel {
			for i := 0; i < indent; i++ {
				print("  ")
			}
			fmt.Print(fmt.Sprintf("%d(%c)(%c)\n", node.key, node.color, flag))
			printnode(node.left, indent+1, 'L')
			printnode(node.right, indent+1, 'R')
		}
	}
	printnode(tree.root, 1, 'G')
}

func (node *RBTreeNode) Key() int64 {
	if node == nil {
		return -1
	} else {
		return node.key
	}
}

func (node *RBTreeNode) Value() interface{} {
	if node == nil {
		return nil
	} else {
		return node.data
	}
}

func (tree *RBTree) PrevOf(node *RBTreeNode) *RBTreeNode {
	if node == nil || node == tree.sentinel {
		return nil
	}

	var getmax func(node *RBTreeNode) *RBTreeNode
	getmax = func(node *RBTreeNode) *RBTreeNode {
		tmp := node
		for {
			if tmp.right != tree.sentinel {
				tmp = tmp.right
			} else {
				return tmp
			}
		}
	}

	if node.left == tree.sentinel {
		tmp := node
		for {
			if tmp.parent != nil && tmp == tmp.parent.left {
				tmp = tmp.parent
			} else {
				break
			}
		}
		return tmp.parent
	}

	return getmax(node.left)
}

func (tree *RBTree) NextOf(node *RBTreeNode) *RBTreeNode {
	if node == nil || node == tree.sentinel {
		return nil
	}

	var getmin func(node *RBTreeNode) *RBTreeNode
	getmin = func(node *RBTreeNode) *RBTreeNode {
		tmp := node
		for {
			if tmp.left != tree.sentinel {
				tmp = tmp.left
			} else {
				return tmp
			}
		}
	}

	if node.right == tree.sentinel {
		tmp := node
		for {
			if tmp.parent != nil && tmp == tmp.parent.right {
				tmp = tmp.parent
			} else {
				break
			}
		}
		return tmp.parent
	}

	return getmin(node.right)
}
