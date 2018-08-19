/**
 * 功能说明：“模块”初始化文件
 *
 * @desc 模块名称
 * ---------------------------------------------------------------------
 * @author		wuxiaoyong <vip120@126.com>
 * @date		2018-08-19
 * @copyright	公司名称
 * ---------------------------------------------------------------------
 */

package main

import (
	"fmt"
)

func main() {
	var students = make(map[string]string)
	var students1 = make(map[string]string)

	students["name"] = "wuxiaoyong"
	students["age"] = "10"

	students1["name"] = "wuxiaoyong"
	students1["age"] = "10"
	//students=append(students,students1)

	fmt.Println(students, students1)

	//取一个map的元素的值，如果存在，返回值和是否索引成功
	if name, ok := students["name"]; ok {
		fmt.Println(name)
	}
	//删除一个map元素的值
	delete(students1, "name")
	fmt.Println(students1)
	//for name:=range students{
	//	fmt.Println(name,"的age 是：",students[name])
	//}

	//map和slice,func 一样，不支持==比较
	//但可以比较长度和map中的值是否相等

	fmt.Println(mapEqu(students, students1))

}

func mapEqu(x, y map[string]string) bool {

	if len(x) != len(y) {
		return false
	}
	for xi, xv := range x {
		if yv, ok := y[xi]; !ok || xv != yv {
			return false
		}
	}
	return true
}
