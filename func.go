package main

import "fmt"

func myeval(a, b int, op string) int  {
	switch op {
	case "+":
		return a+b
	case "-":
		return a-b
	case "*":
		return a*b
	case "/":
		return a/b
	default:
		panic("unsupported operation:" +  op)
	}
	
}

func main()  {
	fmt.Println(myeval(3,4, "*"))

}