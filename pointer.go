package main

import "fmt"

/*
go 语言的指针不能运算
 */

func point1()  {
	var a int = 2
	var pa *int = &a  //pa 是int类型指针，保存变量a的地址
	*pa = 3	          //使用指针变量时，要在变量前加*
	fmt.Println(a)
}



/*
在go语言中只有值传递

值传递：在函数调用参数传递时，函数外部变量的值被复制一份到函数内部，
      函数内部变量的改变不会引起函数外部变量值得改变
引用传递：在函数调用参数传递时，函数外部变量的引用被传递到函数内部，
      函数内部变量值改变，函数外部也会改变 如 java python,注意
      一些语言的内建类型的传递也是值传递

 */



 /*
 go 语言中可以用指针实现引用传递
  */

  var a int = 2
  var pa = &a
func pointer2(pa *int)  {
	*pa = 3
}


/*
交换两个变量的值，方法1
 */
func swap1(a, b *int)  {
	*a,*b = *b, *a
}

/**
交换量个变量的值，方法2
 */

func swap2(a, b int) (int,int) {
	return b,a
}

func main()  {
	//point1();

	//pointer2(pa)
	//fmt.Println(a)  //此时a的值在函数中被赋值成了3

	b := 3
	c := 4
	//swap1(&b, &c)
	b,c = swap2(b, c)
	fmt.Println(b, c)



}