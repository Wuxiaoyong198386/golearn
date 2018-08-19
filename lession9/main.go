package main

import "fmt"

//函数的基础知识
//1、匿名函数
//2、延时功能和函数
//3、中断及错误捉捕

func main() {

	////匿名函数方式一
	//a:= func() {
	//	fmt.Println("this is a demo")
	//}
	//a()
	////匿名函数方式二
	//func(){
	//	fmt.Println("this is a demo1")
	//}()
	//
	////匿名函数传参数及返回值
	//
	//b:= func() int{
	//	return 3
	//}()
	//fmt.Println(b)
	//
	//c:= func() int{
	//	return 3
	//}
	//fmt.Println(c())

	//e:=func(x,j int) int{
	//	return x+j
	//}(1,2)
	//
	//fmt.Println(e)
	//
	//f:=func(x,j int) int{
	//	return x+j
	//}
	//
	//fmt.Println(f(1,1))

	//延时执行

	//a:=1
	//b:=2
	//// defer 如果放在函数中，return 之前的最后一条代码执行
	//defer fmt.Println(a+1,b+1) //最后执行此行代码
	//
	//fmt.Println(a,b)
	// recover 必须放在 defer 语句中
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("this a error")

}
