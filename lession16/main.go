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

/*
 本小节主要学习方法
 方法和函数比较类似但又有本质区别
 1、函数是独立的谁都可以调用的，方法是有关连状态的，是属于类的
 2、函数是通过接收值，完成特定的逻辑计算得到结果
 3、方法是有一个接收体
 4、要想使用方法，必须实例化一个类（对象），用对象使用类中的方法

*/

package main

import "fmt"

//定义一个结构体，也可以理解为类
type Stu struct {
	name string
	age  int16
	cj   //可以调用cj的所有方法和变量
}

type cj struct {
	chinese float64
	english float64
}

func (c *cj) setChinese(chinese float64) {
	c.chinese = chinese
}
func (c *cj) setEnglish(english float64) {
	c.english = english
}

func main() {

	s1 := Stu{} //实例化Stu类
	//调用方法一
	s1.setName("wuxiaoyong") //调用对像的方法
	s1.setAge(20)
	//调用了cj类中的方法
	s1.setChinese(40)
	s1.setEnglish(50)

	fmt.Println(s1)
}

//设置姓名方法
func (s *Stu) setName(name string) {
	s.name = name
}

//设置年龄方法
func (s *Stu) setAge(age int16) {
	s.age = age
}
