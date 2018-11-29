package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/*
四则运算
 */

func myeval1(a, b int, op string) int  {
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
		panic("unsupported operation:" +  op) //panic是中断执行，并非容错
	}

}




/*
定义除运算，返回商和余数，go语言函数可以多返回值
 */
func div(a, b int) (q, r int) {
	//可以给函数的多个返回值值起名字，仅用于非常简单的函数，对于调用者而言是没有感知的
	//q = a / b
	//r = a % b
	//return q, r
	return a / b, a % b
}

/*
go 语言中，多返回值得函数除了可以返回多个值，还可以同时返回值和错误信息
 */
func myeval2(a, b int, op string) (int, error)  {
	switch op {
	case "+":
		return a+b, nil
	case "-":
		return a-b, nil
	case "*":
		return a*b, nil
	case "/":
		// 函数中有多个返回值，但是就想返回一个，那么可以用 _ 做占位符
		q, _  := div(3, 2)  //只返回商
		return q,nil
	default:
		return 0, fmt.Errorf("unsupported operation : %s", op)

	}

}


//go语言中函数式一等公民，函数中也可以放函数
/**
op :形参，表示一个函数接受两个int类型参数，返回int类型
a  :形参，参数a
b  :形参, 参数b

 */
func apply(op func(int, int) int, a, b int) int  {
	//用反射的方法获取函数名
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n",opName, a, b)

	return op(a,b)
}

/*
定义一个子函数用于作为函数式编程的入参
 */

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

/*
go语言函数中没有默认参数，函数重载的用法,操作符重载，唯一有一个可变参数列表
 */
func sum(numbers ...int)int {
	sum := 0
	for i := range numbers{
		sum += numbers[i]
	}

	return sum
}




func main()  {
	//fmt.Println(myeval1(3,4, "*"))
	//
	//q, r := div(13,2)
	//fmt.Println(q, r)

	////判断在函数的执行过程中是否发生了错误
	//if result, err := myeval2(3,4,"$"); err != nil{
	//	fmt.Println("Error:",err)
	//}else {
	//	fmt.Println(result)
	//}

	//传入匿名函数
	//fmt.Println(apply(func(a int, b int) int {
	//	return int(math.Pow(float64(a),float64(b)))
	//},3, 4))

	//传入定义好的函数
	//pow 表示函数名
	//fmt.Println(apply(pow,3, 4))


	fmt.Println(sum(1, 2, 3, 4))



}

/*
要点：
返回值类型写在最后面
可返回多个值
函数作为参数
没有默认参数，可选参数
 */