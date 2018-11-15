package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

/**
定义方法实现勾股定义
 */

 func triangle(){
 	var a, b int =  3,4
 	var c int
 	c = int(math.Sqrt(float64(a * a + b * b )))
 	fmt.Println(c)
 }

 //定义方法实现文件读写，使用go语言的if新特性
func readFiles() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename);err == nil {
		//string将字节码转为字符串
		fmt.Println(string(contents))
	} else {
		fmt.Println("cannot print file contents", err)
	}
}

 /**
 实现文件的读取
  */
func readfiles2() {
	const filename = "abc.txt"
	//注意go语言的函数是可以返回多个值得
	contents, err := ioutil.ReadFile(filename)
	//如果返回的错误不为空
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Printf("%s\n",contents)
	}

}

 /**
 由于go是强类型语言，有时候使用起来会过度严谨让编程不畅
 我们可以考虑常量，常量本身是没有类型的，但是常量可以根据
 使用情况成为任何类型，下面我们从写triangle（）方法，用常
 量定义a, b
  */

 func main(){
 	//triangle()
 	readFiles()
 	readfiles2()
 }