package main

import (
	"fmt"

	"github.com/hiabhi-cpu/collections/stack"
)

func toPostFix(infix []string) (postfix []string) {
	// opr := ""
	stk := stack.NewStack()
	for _, r := range infix {
		if r == "+" || r == "-" || r == "*" || r == "/" || r == "(" || r == ")" {
			if r == ")" {
				for stk.Peek() != "(" {
					postfix = append(postfix, stk.Pop().(string))
				}
				stk.Pop()
				continue
			}
			stk.Push(r)
		} else {
			postfix = append(postfix, r)
			// fmt.Println(r)
		}
	}

	for stk.Len() != 0 {
		postfix = append(postfix, stk.Pop().(string))
	}

	fmt.Println(stk.Len())
	return postfix
}
