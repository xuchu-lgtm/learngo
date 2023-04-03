package main

import "fmt"

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Height int    `json:"height"`
}

type Student struct {
	Person
	score float32
	Name  string
}

func main() {

	student := Student{Person{Name: "test1", Age: 18, Height: 172}, 2, "zhangs"}

	student.Person.Name = "test2"
	student.Age = 19
	student.Name = "zhangsa2"

	fmt.Println(student)

}
