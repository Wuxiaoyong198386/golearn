package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串相加

	var s1 string = "wuxiaoyong "
	var s2 string = "boy"

	fmt.Println(s1 + s2)

	//定义的字符串，不赋值默认为空

	var s3 string
	fmt.Println(s3 == "") //结果为true
	//fmt.Println(s3==nil) //不等于nil 且不能这样比较

	//索引字符串

	var s4 string = "abcdefg"
	for i := 0; i < len(s4); i++ {
		fmt.Printf("%c\n", s4[i])
	}

	for i, v := range s4 {
		fmt.Printf("%d:%c\n", i, v)
	}

	//字符串有特殊字符，需要反斜杠

	var s5 string = "wuxiaoyong \"boy"
	fmt.Println(s5)

	//替换字符串的字符一，此方法好笨

	var s6 string = "abcdefg"
	s7 := []byte(s6)
	s7[0] = 98
	fmt.Printf("%c", s7)

	//替换字符串中的字符

	var s8 string = "wuxiaowuyongwu"
	fmt.Println(strings.Replace(s8, "wu", "hahaha", 2))

}
