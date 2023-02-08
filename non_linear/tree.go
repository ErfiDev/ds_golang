package main

import "fmt"

/*
	In computer science, a tree is a widely used abstract
	data type that represents a hierarchical tree structure
	with a set of connected nodes. Each node in the tree
	can be connected to many children (depending on the type
	of tree), but must be connected to exactly one parent,
	except for the root node, which has no parent.
*/

// implementing DOM tree

type Node struct {
	tag      string
	id       string
	text     string
	children []*Node
}

type DOM struct {
	root *Node
}

func (d *DOM) getElementById(id string) *Node {
	queue := make([]*Node, 0)
	queue = append(queue, d.root)

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		if n.id == id {
			return n
		}

		if len(n.children) > 0 {
			queue = append(queue, n.children...)
		}
	}

	return nil
}

func (d *DOM) append(parent *Node, child *Node) {
	ok := d.getElementById(child.id)
	if ok == nil {
		parent.appendNode(child)
	}
}

func (n *Node) appendNode(node *Node) {
	if node != nil {
		n.children = append(n.children, node)
	}
}

func print(n *Node, level int) {
	if n == nil {
		return
	}

	if level > 0 {
		str := ""
		for i := 0; i < level; i++ {
			str += " "
		}

		fmt.Printf("%s%s\n", str, n.tag)
	} else {
		fmt.Printf(n.tag + "\n")
	}

	level++

	lenChildren := len(n.children)
	if lenChildren > 0 {
		for _, c := range n.children {
			print(c, level)
		}
	}
}

func main() {
	header := Node{
		tag:  "header",
		id:   "header",
		text: "this is the header xxx",
	}

	h1 := Node{
		tag:  "h1",
		id:   "h1",
		text: "hey",
	}

	div := Node{
		tag: "div",
		id:  "container",
	}

	p := Node{
		tag:  "p",
		id:   "p",
		text: "paragraph",
	}

	p1 := Node{
		tag:  "p",
		id:   "p",
		text: "second paragraph",
	}
	footer := Node{
		tag:  "footer",
		id:   "footer",
		text: "this is the footer xxx",
	}

	body := Node{
		tag:      "body",
		children: []*Node{},
	}

	head := Node{
		tag: "head",
		id:  "head",
	}

	title := Node{
		tag:  "title",
		text: "page title",
		id:   "title",
	}

	html := Node{
		tag:      "html",
		children: []*Node{&head, &body},
	}

	dom := DOM{
		root: &html,
	}

	dom.append(&head, &title)
	dom.append(&body, &header)
	dom.append(&body, &h1)
	dom.append(&body, &div)
	dom.append(&body, &footer)
	dom.append(&div, &p)
	dom.append(&div, &p1)

	print(&html, 0)
}
