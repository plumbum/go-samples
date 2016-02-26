package main

import "fmt"

func main() {

	a := 5

	// Классический выбор
	fmt.Println("Classic switch")
	switch a {
	case 5: fmt.Println("Is 5")
	case 6: fmt.Println("Is 6")
	case 7: fmt.Println("Is 7")
	}

	// Условия. Заменяет конструкции if {} else if {}
	fmt.Println("Conditions in case")
	switch {
	case a < 5: fmt.Println("less 5")
	case a < 6: fmt.Println("less 6")
	case a < 7: fmt.Println("less 7")
	}

	// Условия могут быть сложные
	b := true
	fmt.Println("2 Conditions in case")
	switch {
	case a < 5 && b: fmt.Println("less 5")
	case a < 6 && !b: fmt.Println("less 6")
	case a < 7 && b: fmt.Println("less 7")
	default: fmt.Println("default")
	}
}
