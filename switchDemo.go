package main

import "fmt"

/**
实现四则运算
switch有条件语句时
 */
func eval(a, b int, op string ) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

/**
实现一个判分函数
switch 语句无条件语句
 */
func grade(score int) string {
	g := ""
	switch  {
	case score < 60:
		g = "D"
	case score < 70:
		g = "C"
	case score < 80:
		g = "B"
	case score < 100:
		g = "A"
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score %d", score))
	}
	return g

}

func main()  {
	 result := eval(1,3, "+")
	 fmt.Printf("%d\n",result)  //格式化输出
	 fmt.Println(result)                //非格式化输出

	 re := grade(89)
	 fmt.Printf(re)

}

