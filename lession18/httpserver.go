/**
 * 功能说明：“模块”初始化文件
 *
 * @desc 模块名称
 * ---------------------------------------------------------------------
 * @author		wuxiaoyong <vip120@126.com>
 * @date		2018-08-31
 * @copyright	公司名称
 * ---------------------------------------------------------------------
 */

package main

import (
	"fmt"
	"net/http"
)

func getnsqs_get(w http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()
	get_act := query["msg"][0]
	w.Write([]byte(get_act))
	fmt.Println(get_act)

}

func getnsq_post(w http.ResponseWriter, req *http.Request) {

	buf := make([]byte, 1024)
	n, err := req.Body.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buf[0:n]))

}

type Handlers struct{}

func main() {
	http.HandleFunc("/getnsq", getnsqs_get)
	http.HandleFunc("/getnsqpost", getnsq_post)
	http.ListenAndServe(":8001", nil)
}
