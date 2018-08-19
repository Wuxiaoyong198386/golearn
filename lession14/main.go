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

type cj struct {
	chinese float64
	englist float64
}

type sudents struct {
	myname string
	myage  int16
	cj
}

func main() {

	s1 := sudents{}
	s1.myname = "wuxiaoyong"
	s1.myage = 36
	s1.cj.chinese = 20.1
	s1.cj.englist = 40.3
	fmt.Println(s1)

	s2 := sudents{"xiaoming", 29, cj{30.3, 50.4}}
	fmt.Println(s2)

}
