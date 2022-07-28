package main

import "fmt"

type Node struct {
	Key   string
	Left  *Node
	Right *Node
}

type NodeInt struct {
	Key   int
	Left  *NodeInt
	Right *NodeInt
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

	//z := breadthFirstValues(&a)
	//
	//fmt.Println(z)

	//x := treeIncludes(&a, "a")
	//
	//fmt.Println(x)

	three := NodeInt{Key: 3}
	eleven := NodeInt{Key: 11}
	four := NodeInt{Key: 4}
	two := NodeInt{Key: 2}
	four2 := NodeInt{Key: 4}
	one := NodeInt{Key: 1}

	three.Left = &eleven
	three.Right = &four
	eleven.Left = &four2
	eleven.Right = &two
	four.Right = &one

	//x := treeSum(&three)
	//
	//fmt.Println(x)

	y := smallestValue(&three)

	fmt.Println(y)

}
