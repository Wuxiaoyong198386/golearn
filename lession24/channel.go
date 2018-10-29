/**
 * 功能说明：“模块”初始化文件
 *
 * @desc 模块名称
 * ---------------------------------------------------------------------
 * @author		wuxiaoyong <vip120@126.com>
 * @date		2018-09-09
 * @copyright	公司名称
 * ---------------------------------------------------------------------
 */

package main

import "fmt"

//func main(){
//
//	c:=make(chan int,3)
//
//	c<-1
//	c<-2
//	c<-3
//	fmt.Println("a test")
//
//	fmt.Println(<-c)
//	fmt.Println(<-c)
//	fmt.Println(<-c)
//
//}


func main(){

	done:=make(chan struct{})
	c:=make(chan int,3)

	go func() {
		defer close(done)
		for v:=range c{
			fmt.Println(v)
		}
	}()

	c<-1
	c<-2
	c<-3
	close(c)

	<-done

}