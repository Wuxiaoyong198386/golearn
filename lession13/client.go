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
	"bufio"
	"fmt"
	"net"
	"time"
)

var quitSemaphore chan bool

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("error message:" + err.Error())
	} else {
		defer conn.Close()
		fmt.Println("connected!")
		go onMessageRecived(conn)
		b := []byte("time\n")
		conn.Write(b)
		<-quitSemaphore
	}

}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
		time.Sleep(time.Second)
		b := []byte(msg)
		conn.Write(b)
	}
}
