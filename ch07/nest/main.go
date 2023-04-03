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

func (p *Person) print() {
	p.Name = "zhangsa1"
	p.Age = 12
	fmt.Printf("name: %s age: %d \r\n", p.Name, p.Age)
}

// 交换两个值
func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {

	person := Person{"zhang", 19, 20}
	person.print()

	fmt.Println(person.Age)

	student := Student{Person{Name: "test1", Age: 18, Height: 172}, 2, "zhangs"}

	student.Person.Name = "test2"
	student.Age = 19
	student.Name = "zhangsa2"

	fmt.Println(student)

	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
