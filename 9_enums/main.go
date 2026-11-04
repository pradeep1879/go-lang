package main

import "fmt"

type OrderStatus string

const (
	// Received OrderStatus = "received"  //Untyped int.
	Confirmed = "confirmed"
	Prepared	= "prepared"
	Delivered = "delivered"
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("chaning order status to", status)
}


func main(){
	changeOrderStatus(Confirmed)
}