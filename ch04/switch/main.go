package main

import "fmt"

func main() {

	day := "星期一"
	switch day {
	case "星期一":
		fmt.Println("Monday")
	case "星期二":
		fmt.Println("Tuesday")
	default:
		fmt.Println("UnKnown")
	}

	score := 69
	switch {
	case score >= 60 && score < 70:
		fmt.Println("D")
	case score >= 70 && score < 80:
		fmt.Println("C")
	}

	score = 60
	switch score {
	case 50, 60, 80:
		fmt.Println("Wow")
	default:
		fmt.Println("default")
	}
}
