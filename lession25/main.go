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

 4、base64位存为图片


*/

package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"io/ioutil"
	"encoding/base64"
	"log"
	"path"
)

func main() {

	//文本读取模式
	/*
	filepath := "test.txt"
	readType := 1
	content, err := ReadFile(filepath, readType)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("文件读取成功！内容如下：\n"+content[0])

	//文本写入模式
	filename:="test1.txt"
	writeType:=2  //1:追加模式 2、替换模式
	filecontent:=[]byte("golang aaaa文本文件创建、写入操作\n")
	if _,err:=WriteFile(filename,filecontent,writeType);err!=nil{
		fmt.Println(err)
	}
	fmt.Println("文件写入成功！")
	*/
	//图片base64读取
	filepath:="qiniu_20181024.jpg"
	base64,err:=ImageToBase64(filepath)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(base64)
	//图片base64转存图片

}

func ImageToBase64(filepath string)(string,error){

	if isE:=IsExists(filepath);!isE{
		return "",errors.New("图片文件不存在！")
	}
	const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	var coder = base64.NewEncoding(base64Table)
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	filenamewithsuffix:=path.Base(filepath)
	fileSuffix:=string([]byte(path.Ext(filenamewithsuffix))[1:4])
	return "data:image/"+fileSuffix+";base64,"+string([]byte(coder.EncodeToString(data))),nil
}

func Base64ToImage(base64 string)error{



	return nil
}


func WriteFile(filepath string,content []byte,writerType int) (bool,error){
	var f *os.File
	var err error
	//判断文件是否存在，不存在就创建
	//isExists:=IsExists(filepath)

	if writerType==1{ //追加模式
		f, err = os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	}else{  //替换模式
		f, err = os.OpenFile(filepath, os.O_WRONLY | os.O_CREATE, 0644)

	}
	if err!=nil{
		return false,errors.New("创建文件失败！")
	}
	defer f.Close()
	if _,_err:=f.Write(content);_err!=nil{
		return false,errors.New("文件写入失败！")
	}
	return true,nil
}

//判断文件是否存在
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
