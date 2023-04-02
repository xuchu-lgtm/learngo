package main

import "fmt"

func main() {

	var courses []string
	fmt.Printf("%T \r\n", courses)

	courses = append(courses, "go")
	courses = append(courses, "grpc")
	courses = append(courses, "gin")
	fmt.Println(courses)

	allCourse := [5]string{"go", "gin", "grpc", "mysql", "elasticsearch"}
	courseSlice := allCourse[0:2]
	fmt.Printf("%T , %T \r\n", allCourse, courseSlice)
	fmt.Println(allCourse[:])
	fmt.Println(courseSlice)

	var allCourse2 []string
	allCourse2 = append(allCourse2, "ccc")
	fmt.Println(allCourse2)

	allCourse3 := make([]string, 0, 3)
	allCourse3 = append(allCourse3, "A")
	allCourse3 = append(allCourse3, "B")
	allCourse3 = append(allCourse3, "C")
	fmt.Println(allCourse3)

	courses4 := []string{"mysql", "go"}
	allCourse3 = append(allCourse3, courses4...)
	fmt.Println(allCourse3)

	allCourse5 := [5]string{"go", "gin", "grpc", "mysql", "elasticsearch"}

	delAllCourse5 := append(allCourse5[:3], allCourse5[4:]...)
	fmt.Println(delAllCourse5)

	courseSliceCopy := make([]string, len(delAllCourse5))
	copy(courseSliceCopy, delAllCourse5)
	fmt.Println(courseSliceCopy)
}
