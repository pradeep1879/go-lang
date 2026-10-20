package main

import "fmt"

// variables

func str(){
	// var name string = "golang"
	var isAdult bool = false

	name := "golang" // short hand syntax
	
	// var price float32 = 50.5
	price := 50.5

	fmt.Println(name);
	fmt.Println(isAdult);
	fmt.Println(price);
}

const age = 30 // also can be declared out side the function

func constants(){
	const name = "golang"
		// name = "Javascript" can't assigne once it is created with const

		const (
			port = 3000
			host = "localhost"
		)
	fmt.Println(port, host)	

	fmt.Println(name)
	fmt.Println(age)
}

// loop in golang (only construct in go for looping)
func loop(){
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// for{.            // infinity loop
	// 	fmt.Println("1")
	// }

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	// range
	for i := range 10 {
		fmt.Println(i)
	}
}

/// if else

func condition(){
	age := 16
	// if age >= 18 {
	// 	fmt.Println("Persion is adult")
	// } else {
	// 	fmt.Println("Persion is not an adult")
	// }

	if age >= 18 {
		fmt.Println("Persion is adult")
	} else if age >= 12 {
		fmt.Println("Persion is teenager")
	} else {
		fmt.Println("Persion is kid")
	}


	var role string = "admin"
	var hasPersmission bool = false

	if role == "admin" && hasPersmission {
		fmt.Println("yes")
	}else {
		fmt.Println("no")
	}

	// go doesn't ternary operator

}

func switchConst(){
	i := 5
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("default")
	}
}





func main() {
	fmt.Println("Hello world")
	// str()
	// constants()
	// loop()
	// condition()
	switchConst()
}