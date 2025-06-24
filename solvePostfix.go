	package main

	import (
		"errors"
		"fmt"
		"strconv"

		"github.com/hiabhi-cpu/collections/stack"
	)

	func solvePostfix(postfix []string) (string, error) {

		stk := stack.NewStack()

		for _, r := range postfix {
			if isOpr(r) {
				a := stk.Pop().(string)
				b := stk.Pop().(string)
				res, err := solve(b, a, r)
				if err != nil {
					return "0", errors.New("Something happended")
				}
				stk.Push(res)
			} else {
				stk.Push(r)
			}
			fmt.Println(stk.Peek())
		}

		return stk.Pop().(string), nil
	}

	func solve(a, b, opr string) (string, error) {
		aInt, err := strconv.ParseInt(a, 10, 0)
		if err != nil {
			fmt.Println("Error converting:", err)
			return "0", err
		}

		bInt, err := strconv.ParseInt(b, 10, 0)
		if err != nil {
			fmt.Println("Error converting:", err)
			return "0", err
		}
		if opr == "+" {
			return fmt.Sprint(aInt + bInt), nil
		} else if opr == "-" {
			return fmt.Sprint(aInt - bInt), nil
		} else if opr == "*" {
			fmt.Println(opr)
			return fmt.Sprint(aInt * bInt), nil
		} else if opr == "/" {
			return fmt.Sprint(aInt / bInt), nil
		} else {
			return "0", errors.New("Give correct symbol")
		}
	}

	func isOpr(r string) bool {
		if r == "+" || r == "-" || r == "*" || r == "/" {
			return true
		}
		return false
	}
