package main

import "fmt"

// import _ "txt" 仅初始化使用
// import . "txt" 解压包，直接使用函数
// import t "txt" 别名

type MyWriter interface {
	Writer(string)
}

type MyReader interface {
	Reader() string
}

type MyReadWriter interface {
	MyWriter
	MyReader
	MyReadWriter()
}

type SreadWriter struct {
}

func (s *SreadWriter) Writer(s2 string) {
	fmt.Println(s2)
}

func (s *SreadWriter) Reader() string {
	//TODO implement me
	panic("implement me")
}

func (s *SreadWriter) MyReadWriter() {
	//TODO implement me
	panic("implement me")
}

func main() {
	var my MyReadWriter = &SreadWriter{}

	my.Writer("test")
}
