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
	
}


