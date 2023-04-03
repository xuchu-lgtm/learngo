package main

import "fmt"

type MyWriter interface {
	Write(string) error
}

type fileWriter struct {
	filePath string
}

type writerCloser struct {
	MyWriter
}

type databaseWriter struct {
	host string
	port string
	db   string
}

func (fw *fileWriter) Write(string) error {
	fmt.Println("write string to file")
	return nil
}

func (dw *databaseWriter) Write(string) error {
	fmt.Println("write string to database")
	return nil
}

// 断言
func add(a, b interface{}) int {
	return a.(int) + b.(int)
}

// 非泛型的实现，比较麻烦
func add2(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi
	case int32:
		ai, _ := a.(int32)
		bi, _ := b.(int32)
		return ai + bi
	case int64:
		ai, _ := a.(int64)
		bi, _ := b.(int64)
		return ai + bi
	case string:
		ai, _ := a.(string)
		bi, _ := b.(string)
		return ai + bi
	default:
		panic("error")
	}
}

func main() {

	wc := &writerCloser{
		&databaseWriter{},
	}
	wc.Write("test")

	fmt.Println(add(1, 2))
}
