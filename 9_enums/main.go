package main

import "fmt"



func changeOrderStatus(status string){
	fmt.Println("chaning order status to", status)
}


func main(){
	changeOrderStatus("confirmed")
}