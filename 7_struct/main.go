package main

import (
	"fmt"
	"time"
)


type Order struct {
	id	string
	amount float32
	status	string
	createdAt	time.Time // nanosecond precision
}

// receiver type
func (o *Order) changeStatus(status string){
	o.status = status
}

func (o *Order) getAmount(amount float32){
	o.amount = amount
}


func main(){
	myOrder := Order{
		id: "1",
		amount: 50.00,
		status: "received",

	}
	myOrder.createdAt = time.Now()

	fmt.Println(myOrder.status)
	fmt.Println(myOrder.amount)

	fmt.Println("Order struct", myOrder)
	
	myOrder.changeStatus("confirmed")
	myOrder.getAmount(99.00)

	fmt.Println(myOrder.status)
	fmt.Println(myOrder.amount)

}