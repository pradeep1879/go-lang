package main

import "fmt"

func main(){
	nums := []int{5, 6, 7, 8}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }


	for i , num := range nums {
		fmt.Println(num, i)
	}

	m := map[string]string{"fname": "john", "lname": "doe"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// unicode code point rune
	// starting byte of rune
	// 300 -> 1 byte , 2 byte

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}
