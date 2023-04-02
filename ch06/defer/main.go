package main

import "fmt"

// defer是有能力修改返回值的，因为在函数执行结束之前defer是先执行的
func deferReturn() (ret int) {
	defer func() {
		ret++
	}()
	return 10
}

func main() {
	//连接数据库，打开文件、使用锁、无论如何最后都要记得关闭数据库、关闭文件、解锁

	ret := deferReturn()
	fmt.Printf("ret = %d \r\n", ret)
}
