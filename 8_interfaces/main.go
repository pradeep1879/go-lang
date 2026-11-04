package main

import "fmt"



type payment struct {
	gatway stripe
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}
	
	// razorpayPaymentGw.pay(amount)
	p.gatway.pay(amount)
}

type razorpay struct {

}

func (r razorpay) pay(amount float32){
	// logic to make payment
	fmt.Println("Making payment using razorpay", amount)
}

type stripe struct {}

func (s stripe) pay(amount float32){
	
	fmt.Println("Making payment using stripe", amount)
}


func main(){
	stripePaymentGw := stripe{}
	newPayment := payment{
		gatway: stripePaymentGw,
	}
	newPayment.makePayment(100)
}