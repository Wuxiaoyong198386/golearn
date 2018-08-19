package main

import (
	"fmt"

	"github.com/Wuxiaoyong198386/golearn/lession7/pkg"
)

//函数知识点
//1、函数的使用无需前置声明。
//2、函数内部不能嵌套命名函数，但可以有匿名函数
//3、函数内的局部变量可以返回，局部变量放在盏中，昨时存放，当返回时，会从盏中复制到堆中。
//4、函数名第一个字符大小写很重要，大写可以被其它包引用，小写只能被当前包引用
//5、函数形参不能有默认值，且实参和形参必须一一对应
//6、函数都是传值，而不是指针。

func main() {
	sayHello()
	var myName1 *int = sayMyName()
	fmt.Println(myName1, *myName1)
	pkg.SayHello()

	var a1 []int = []int{1, 2, 3, 4, 5}
	//showNum("aaaa",a1...)//直接把数组打散成n个元素
	//showNum("aaaa",a1[:]...)//把数组转为切片
}

func showNum(x string, arr ...int) {
	fmt.Println(arr)
}

func sayHello() {
	fmt.Println("helloworld")
	//匿名函数
	a := func() { fmt.Println("function in function") }
	a()
}

func sayMyName() *int {
	myName := 100
	return &myName
}
