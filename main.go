package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hrllo")
	if len(os.Args) != 2 {
		fmt.Println("Give correct argumenrs")
		return
	}
	exp := strings.Split(os.Args[1], " ")
	// fmt.Println(exp[0])
	// res, err := getResultInt(exp[0], exp[2], exp[1])
	// if err != nil {
	// 	fmt.Println("No correct expression")
	// 	return
	// }
	// fmt.Println(res)

	postFix := toPostFix(exp)
	fmt.Println(postFix)
}

func getResultInt(a, b, opr string) (int64, error) {

	aInt, err := strconv.ParseInt(a, 10, 0)
	if err != nil {
		fmt.Println("Error converting:", err)
		return 0, err
	}

	bInt, err := strconv.ParseInt(b, 10, 0)
	if err != nil {
		fmt.Println("Error converting:", err)
		return 0, err
	}

	if opr == "+" {
		return aInt + bInt, nil
	} else if opr == "-" {
		return aInt - bInt, nil
	} else if opr == "*" {
		fmt.Println(opr)
		return aInt * bInt, nil
	} else if opr == "/" {
		return aInt / bInt, nil
	} else {
		return 0, errors.New("Give correct symbol")
	}

}
