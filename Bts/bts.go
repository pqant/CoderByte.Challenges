package Bts

import "log"

func init() {

}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

type Tree struct {
	Node *Node
}

func (t *Tree) Insert(value int) *Tree {
	if t.Node == nil {
		t.Node = &Node{value: value}
	} else {
		t.Node.insert(value)
	}
	return t
}

func (n *Node) Exist(value int) bool {
	if n == nil {
		return false
	}
	if n.value == value {
		return true
	}
	if value <= n.value {
		return n.left.Exist(value)
	} else {
		return n.right.Exist(value)
	}
}

func (t *Tree) PrintNodes() {
	if t == nil {
		log.Printf("Tree is empty!\n")
	}
	if t.Node == nil {
		log.Printf("There is no node in the tree!")
	}
	t.Node.PrintNode("****")
}

func (n *Node) PrintNode(sign string) {
	if n == nil {
		return
	}
	log.Printf("%v Node value : %v\n", sign, n.value)
	n.left.PrintNode("<<<<")
	n.right.PrintNode(">>>>")
}

func (n *Node) Search(value int) *Node {
	if n == nil {
		return nil
	}
	if n.value == value {
		return n
	}
	if value <= n.value {
		return n.left.Search(value)
	} else {
		return n.right.Search(value)
	}
}

// TODO : Complete delete operation!
