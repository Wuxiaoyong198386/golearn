/**
 * 功能说明：“模块”初始化文件
 *
 * @desc 模块名称
 * ---------------------------------------------------------------------
 * @author		wuxiaoyong <vip120@126.com>
 * @date		2018-09-03
 * @copyright	公司名称
 * ---------------------------------------------------------------------
 */

package main

import "fmt"

type Point struct {
	x int
	y int
}

func main(){
	p:=new(Point)
	p.x=20
	p.y=30
	fmt.Printf("%T",p)
	//fmt.Println(p)

	p1:=make([]int,0)
	fmt.Printf("%T",p1)
}
