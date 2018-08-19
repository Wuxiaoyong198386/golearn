/**
 * 功能说明：“模块”初始化文件
 *
 * @desc 模块名称
 * ---------------------------------------------------------------------
 * @author		wuxiaoyong <vip120@126.com>
 * @date		2018-08-18
 * @copyright	公司名称
 * ---------------------------------------------------------------------
 */

package main

import (
	"fmt"
	"net"
)

func main() {
	demain := "www.baidu.com"
	ip := net.ParseIP(name)
	mask := ip.DefaultMask()

	network := ip.Mask(mask)
	fmt.Println(network.String())
}
