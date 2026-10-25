package main

import "fmt"

// func changeNumber(num int){
// 	num = 5
// 	fmt.Println("In changeNumber", num)
// }


func changeNumber(num *int){
	*num = 5
	fmt.Println("In changeNumber", *num)
}

func main(){
	num := 1
	changeNumber(&num)

	fmt.Println("After changeNumber in main", num)
}