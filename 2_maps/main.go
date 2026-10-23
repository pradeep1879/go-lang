package main

import (
	"fmt"
	"maps"
)

func main(){
	m := make(map[string]string)

	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["age"]) // IMP: if key value doesn't exits then its return zero
	fmt.Println(m["name"])
	fmt.Println(m["area"])
	fmt.Println(len(m))
	
	delete(m, "area")
	fmt.Println(m)
	
	clear(m)
	fmt.Println(m)


	m1 := map[string]int{"price": 99, "count": 3}
	m2 := map[string]int{"price": 99, "count": 3}

	fmt.Println(m1)

	v, ok := m1["count"]

	fmt.Println(v) // first will give the value and second will gave boolean

	if ok {
		fmt.Println("All ok")
	}else{
		fmt.Println("Not ok")
	}
	fmt.Println(maps.Equal(m1, m2))

	
}