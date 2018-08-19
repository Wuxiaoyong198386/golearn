/**
 * 功能说明：“模块”初始化文件
 *
 * @desc 模块名称
 * ---------------------------------------------------------------------
 * @author		wuxiaoyong <vip120@126.com>
 * @date		2018-08-14
 * @copyright	公司名称
 * ---------------------------------------------------------------------
 */

package main

import (
	"fmt"
)

func main() {

	s := [...]int{1, 2, 3, 4}
	fmt.Println(len(s))
	fmt.Println(s[0])

	s1 := make([]int, 3)
	fmt.Println(s1)

}
