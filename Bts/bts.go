package Bts

import (
	"log"
	"sync"
)

func init() {

}

type Node struct {
	Value int
	left  *Node
	right *Node
}

func (n *Node) insert(value int) {
	if value <= n.Value {
		if n.left == nil {
			n.left = &Node{Value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{Value: value}
		} else {
			n.right.insert(value)
		}
	}
}

type Tree struct {
	Node *Node
	lock sync.RWMutex
}

func (t *Tree) Insert(value int) *Tree {
	t.lock.Lock()
	defer t.lock.Unlock()
	if t.Node == nil {
		t.Node = &Node{Value: value}
	} else {
		t.Node.insert(value)
	}
	return t
}

func (n *Node) Exist(value int) bool {
	if n == nil {
		return false
	}
	if n.Value == value {
		return true
	}
	if value <= n.Value {
		return n.left.Exist(value)
	} else {
		return n.right.Exist(value)
	}
}

func (t *Tree) PrintNodes() {
	if t == nil {
		log.Printf("Tree is nil!\n")
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
	log.Printf("%v Node value : %v\n", sign, n.Value)
	n.left.PrintNode("<<<<")
	n.right.PrintNode(">>>>")
}

func (n *Node) Search(value int) *Node {
	if n == nil {
		return nil
	}
	if n.Value == value {
		return n
	}
	if value <= n.Value {
		return n.left.Search(value)
	} else {
		return n.right.Search(value)
	}
}

func (t *Tree) MinNode() *Node {
	t.lock.RLock()
	defer t.lock.RUnlock()
	if t == nil && t.Node == nil {
		return nil
	}
	if t.Node.left == nil {
		return t.Node
	}
	temp := t.Node.left
	for {
		if temp.left != nil {
			temp = temp.left
		} else {
			return temp
		}
	}
}

func (t *Tree) MaxNode() *Node {
	t.lock.RLock()
	defer t.lock.RUnlock()
	if t == nil && t.Node == nil {
		return nil
	}
	if t.Node.right == nil {
		return t.Node
	}
	temp := t.Node.right
	for {
		if temp.right != nil {
			temp = temp.right
		} else {
			return temp
		}
	}
}


// TODO : Complete delete operation!
