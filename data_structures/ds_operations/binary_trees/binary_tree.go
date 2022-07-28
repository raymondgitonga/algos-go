package main

import "fmt"

type Node struct {
	Key   string
	Left  *Node
	Right *Node
}

func main() {
	a := Node{Key: "a"}
	b := Node{Key: "b"}
	c := Node{Key: "c"}
	d := Node{Key: "d"}
	e := Node{Key: "e"}
	f := Node{Key: "f"}

	a.Left = &b
	a.Right = &c
	b.Left = &d
	b.Right = &e
	c.Right = &f

	//	 a
	//    /  \
	//b		c
	// / \    |
	//d  e     f
	//x := depthFirstValues(&a)
	//fmt.Println(x)
	//
	//y := depthFirstValuesRecursion(&a)
	//fmt.Println(y)

	z := breadthFirstValues(&a)

	fmt.Println(z)
}
