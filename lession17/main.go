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

/**
 本小节主要学习go中的接口interface
 1、和结构体很像，和java中的抽类很像
 2、接口是一种方法声明的集合，并没有具体实现方法。
 3、如果想要实现这个接口，就必须实现这个集合的所有方法，一个没有实现，就会报错
 4、接口和实现的方法并没有明确的关系
**/

package main

import "fmt"

type Stu interface {
	setName()
	setAge()
}

type data struct {
	name string
	age  int64
}

func (d *data) setName() {
	fmt.Println(d.name)
}
func (d *data) setAge() {
	fmt.Println(d.age)
}

func main() {

	d := &data{"wuxiaoyong", 40} //&取指针，所以上面的方法要用 * 。如果不用&，方法中也不要用*
	var i Stu
	i = d
	i.setName()

	i.setAge()
}
