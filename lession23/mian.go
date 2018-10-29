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
//goroutine同步
package main

import (
	"fmt"
	"time"
	"sync"
)

func main()  {

	//第一种方法  channel

	fmt.Println(time.Now())
	goexit:=make(chan int)  //channel 通道类型
	go func() {
		fmt.Println("start test1")
		time.Sleep(time.Second*2)
		fmt.Println("start test2")
		goexit<-1			//goroutine结束，通道写值
	}()

	//time.Sleep(time.Second*3)

	<-goexit  				//如果此值为空，一直在阻塞，等待有值写入
	fmt.Println(time.Now())


	//第二种方法
	fmt.Println("-----------------")
	fmt.Println(time.Now())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("start test3")
		time.Sleep(time.Second*2)
		fmt.Println("start test4")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(time.Now())



}
