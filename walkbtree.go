package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkTree(t, ch)
	close(ch)
}

func walkTree(t *tree.Tree, ch chan int){
	if t != nil {
		if &t.Left != nil {
			walkTree(t.Left, ch)
		} else {
			return
		}
		ch <- t.Value
		if &t.Right != nil {
			walkTree(t.Right, ch)
		} else {
			return
		}
		
		return
	} else {
		return
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(tree.New(1), ch1)
	go Walk(tree.New(1), ch2)
	for v1 := range ch1 {
		if v1 != <-ch2 {
			return false	
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Print(v)
	}
	fmt.Println()
	fmt.Println(Same(tree.New(1), tree.New(1)))
}
