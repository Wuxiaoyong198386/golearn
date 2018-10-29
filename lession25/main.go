/**
 * 功能说明：“模块”初始化文件
 *
 * @desc 模块名称
 * ---------------------------------------------------------------------
 * @author		wuxiaoyong <vip120@126.com>
 * @date		2018-10-28
 * @copyright	公司名称
 * ---------------------------------------------------------------------
 */

/*
 1、读取txt文件，两种模式
 	a、一行一行的读取
 	b、一次读取所有文本
 2、写入txt文件

 3、读取图片，根据不同的类型相对应的转为base64位


*/

package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"io/ioutil"
)

func main() {

	//读取模式
	filepath := "test.txt"
	readType := 2
	content, err := ReadFile(filepath, readType)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(content[0])

	//写入模式
	filename:="test1.txt"
	writeType:=1  //1:追加模式 2、替换模式
	result,err:=WriteFile(filename,writeType)
	if result{

	}

}

func WriteFile(filepath string,writerType int) (bool,error){

	//判断文件是否存在，不存在就创建
	isExists:=IsExists(filepath)
	if !isExists{
		os.Create()
	}

	return true,nil
}


func IsExists(filepath string)bool{
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func ReadFile(filepath string, readType int) ([]string, error) {
	//判断文件是否存在，可用函数实现
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("error msg:" + filepath + "文件不存在")
		}
	}
	//开始读取文件
	//readtype:=1	//1、行读取模式 2、全部模式

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println()
		return nil, errors.New("error msg:打开文件文件！")
	}
	defer file.Close()
	var content []string
	if readType == 1 {
		r := bufio.NewReader(file)
		for {
			c, err := readline(r)
			if err != nil {
				break
			}
			content = append(content, c)
		}
	}else{
		content,err:=ioutil.ReadAll(file)
		if err!=nil{
			return nil,err
		}
		c:=make([]string,0)
		c=append(c,string(content))
		return c,nil


	}
	return content, nil
}

//读取一行文本
func readline(r *bufio.Reader) (string, error) {

	line, isprefix, err := r.ReadLine()
	for isprefix && err == nil {
		var bs []byte
		bs, isprefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return string(line), err
}
