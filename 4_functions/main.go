package main

import "fmt"


// func add(a , b int)int{
// 	return a + b
// }

func getLanguages() (string, string, bool, int){
	return "golang", "javascript", false, 99
}

func processIt1(fn func(a string) string) {
	fn("javscript")
}


func processIt2() func (a int) int {
	return func(a int) int {
		return 4
	}
}


// variadic functions

func variadicFunctions(nums ...int) int { // we can use interface for any types
	total := 0
	
	for _ , num := range nums {
		total = total + num
	}
	return total
}

func main(){


	// result := add(3, 5)
	// fmt.Println(result)

	lang1, lang2, lang3, _ := getLanguages()
	fmt.Println(lang1, lang2, lang3)

	fn := func(a string) string {
		return "golang"
	}

	processIt1(fn)

	fn2 := processIt2()
	fn2(4)
	
	result1 := variadicFunctions(3, 4, 5, 6, 7)
	fmt.Println(result1)

	nums := []int{8, 9, 10, 11}
	result2 := variadicFunctions(nums...)
	fmt.Println(result2)


}



