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
//GOMAXPROCS

package main

import (
	"fmt"
	"math"
	"time"
	"runtime"
	"sync"
)

func count() {

	var x int32
	for i := 0; i < math.MaxInt32; i++ {
		x +=1
	}
	fmt.Println(x)
}

func test1(n int){
	for i:=0;i<n;i++{
		count()
	}
}

func test2(n int){

	var wg sync.WaitGroup
	wg.Add(n)
	go func() {
		for i:=0;i<n;i++{
			count()
			wg.Done()
		}

	}()

	wg.Wait()

}


func main() {

	cpunum:=runtime.GOMAXPROCS(0)
	fmt.Println(time.Now())
	test1(cpunum)
	fmt.Println(time.Now())

	fmt.Println("---------------")

	fmt.Println(time.Now())
	test2(cpunum)
	fmt.Println(time.Now())
}
