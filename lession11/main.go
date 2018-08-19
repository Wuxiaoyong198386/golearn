package main

import (
	"fmt"
)

//数组：
//类型，长度，值，只有相同类型和长度的数组才可以比较
func main() {

	//定义空数组
	var arr1 = [2]int{}

	arr1[0] = 10
	arr1[1] = 11
	fmt.Println(arr1)
	//定义+赋值
	var arr2 = [2]int{20, 21}
	fmt.Println(arr2)

	//数组的比较
	if arr1 == arr2 {
		fmt.Println("ture")
	} else {
		fmt.Println("false")
	}

	//数组的遍历
	s3 := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	for i, v := range s3 {
		fmt.Printf("%d:%d\n", i, v)
	}

	//指针数组   数组指针
	// &取变量的内存地址，也就是指针
	s4 := [3]string{"a", "b", "c"}
	fmt.Printf("%p:%v", &s4, s4)

	s5 := [2]*string{} //指针类型的数组
	s6 := "wu"
	s7 := "xiaoyong"

	s5[0] = &s6
	s5[1] = &s7

	fmt.Println(s5)

	//把数组传弟这给函数要注意
	//1、数组复制，包括地址
	//2、如果要改变数组的值，需把数组指针传弟过去

	s8 := [3]int{1, 2, 3}
	testArray(&s8)
	testArray1(s8) //并没有改变值
	fmt.Println(s8)
	fmt.Println(s8)

	//定义不定长度的数组用...
	var s9 = [...]int{1, 2, 3}
	fmt.Println(s9)

}

func testArray(a *[3]int) {
	a[0] = 11
}
func testArray1(a [3]int) {
	a[0] = 22
}
