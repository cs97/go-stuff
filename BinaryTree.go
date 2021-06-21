package main

import (
 	"fmt"
)

type node struct {
    data int
    l *node
    r *node
}

func insert(n, x *node) {
	if n.data == x.data {
		return
	}
	
	if n.data > x.data {
		if n.l == nil {
			n.l = x
		} else {
			insert(n.l, x)
		}
	} else {
		if n.r == nil {
			n.r = x
		} else {
			insert(n.r, x)
		}
	}
}

func member(n *node, i int) bool {
	if n.data == i {
		return true
	} else {
		if n.data > i {
			if n.l == nil {
				return false
			} else {
				return member(n.l, i)
			}
		} else {
			if n.r == nil {
				return false
			} else {
				return member(n.r, i)
			}
		}
	}
}

func main() {
	tmp :=  &node{data: 69}
	tmp2 := &node{data: 666}
	tmp3 := &node{data: 13}
	insert(tmp, tmp2)
	insert(tmp, tmp3)	
	
	fmt.Println(tmp.data)
	fmt.Println(tmp.r.data)
	fmt.Println(tmp.l.data)	
	
	fmt.Println(member(tmp, 666))
	fmt.Println(member(tmp, 99))
	
}
