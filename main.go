package main

import "fmt"

type IntSet struct {
	name  string
	myage int64
}

//指针接收者的方法
func (i *IntSet) String() string {
	return i.name
}

func main() {

	var s IntSet
	s.name = "wuxiaoyong"
	fmt.Println(s.String())
}
