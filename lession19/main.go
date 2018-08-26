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
 本节学习接口及接口的嵌套

*/

package main

import "fmt"

type stringer interface {
	string() string
}

type tester interface {
	stringer //接口的嵌套
	test()
}

type data struct{}

func (d *data) string() string {
	fmt.Println("aaaa")
	return "aaa"
}

//通过指针接值
func (d *data) test() {
	fmt.Println("bbbb")
}

func main() {

	var a data
	var b stringer = &a //通过指针传值
	b.string()          //只有这一个方法

	var c tester = &a //通过指针传值
	c.test()          //自己的方法
	c.string()        //继承过来的方法

	var s stringer = &a
	pp(s)
}

func pp(s stringer) {
	s.string()
}
