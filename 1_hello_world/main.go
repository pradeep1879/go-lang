package main

import (
	"fmt"
	"slices"
)

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

// arrays
func arrays(){
	// var arr [4]int

	// arr[0] = 1

	// val := [3]int{1, 2, 3}

	// fmt.Println(len(arr));
	// fmt.Println((arr[0]));
	// fmt.Println((val));
	
	
	// // 2d- array
	// num := [2][2]int{{3, 4}, {5,6}}
	// fmt.Println((num));

	// fixed size, that is predictable
	// memory optimazation
	// contant time access



	// SLICES
	// --> dynamic
	// --> most used construct in go
	// --> useful method

	// var slc []int // uninitialized slice in nil
	// fmt.Println(slc == nil)
	// fmt.Println(len(slc))

	var nums1 = make([]int, 2, 5) // 5 = initial capacity , 2 = starting two element will be zero, so put 0 instead of 2
	fmt.Println(nums1 == nil)
	fmt.Println(cap(nums1)) // maximum numbers of elements can fit
	
	// nums1 = append(nums1, 1)
	// nums1 = append(nums1, 2)
	// nums1 = append(nums1, 3)
	// nums1 = append(nums1, 4)
	// fmt.Println(nums1) // 
	// fmt.Println(cap(nums1)) // 

	var num3 = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(num3[:2]) // start:ending(excluded)
	fmt.Println(num3[1:]) // start:ending(excluded)

	var num4 = []int{1, 2, 3}
	var num5 = []int{1, 2, 3, 4}

	fmt.Println(slices.Equal(num4, num5))


	// 2d slices
	var arr2d = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2d)
}

func main() {
	fmt.Println("Hello world")
	// str()
	// constants()
	// loop()
	// condition()
	// switchConst()
	arrays()
}