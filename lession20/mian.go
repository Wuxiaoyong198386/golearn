/**
 * 功能说明：“模块”初始化文件
 *
 * @desc 模块名称
 * ---------------------------------------------------------------------
 * @author		wuxiaoyong <vip120@126.com>
 * @date		2018-08-26
 * @copyright	公司名称
 * ---------------------------------------------------------------------
 */
/*
类型断言


*/

package main

import "fmt"

func main() {
	//基本类型
	var i int
	var s float64
	i = 100
	//s=i  //类型不对，所以无法这样赋值,必须类型转换，如下行
	s = float64(i)
	fmt.Println(i)
	fmt.Println(s)

	//字符串不能转为数值型

	//定义一个mpa
	var m = make(map[string]interface{}, 0)
	m["key1"] = "value1"
	m["key2"] = "value2"
	fmt.Println(m)

	var ts = &ts{"wuxiaoyong", 35} // var ts=ts{"wuxiaoyong",35}
	//如果把ts中的值保存到map中？
	m["key3"] = *ts //注意*和&一般是对称出现		//m["key3"]=ts

	fmt.Println(m)

}

type ts struct {
	name string
	age  int
}
