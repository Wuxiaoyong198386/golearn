package main

import "fmt"

//常量定义和账值
//1、常量可以定义但可以不使用，也不会出错
//2、常量的值不可以改变
//3、变量可以用&得到内存地址，常量却没有内存地址

const myname string = "wuxiaoyong"

//类似枚举类型  0 1 2 3
const (
	x = iota //必须是iota，默认为在当前的位置，从0开始 如果直接赋其它值，三个值会相同
	y
	z
	a = iota //值为3
)

func main() {

	fmt.Println(myname)
	fmt.Println(x, y, z, a)
}
