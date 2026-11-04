// package main

// import "fmt"

// type payment struct {
// 	gatway razorpay
// }

// func (p payment) makePayment(amount float32) {
// 	// razorpayPaymentGw := razorpay{}

// 	// razorpayPaymentGw.pay(amount)
// 	p.gatway.pay(amount)
// }

// type razorpay struct {

// }

// func (r razorpay) pay(amount float32){
// 	// logic to make payment
// 	fmt.Println("Making payment using razorpay", amount)
// }

// type stripe struct {}

// func (s stripe) pay(amount float32){

// 	fmt.Println("Making payment using stripe", amount)
// }

// type fakePayment{}

// func (f fakePayment) pay(amount float32){
// 	fmt.Println("Making payment using fake payment gatway", am)
// }

// func main(){
// 	// stripePaymentGw := stripe{}
// 	razorpayPaymentGw := razorpay{}
// 	newPayment := payment{
// 		gatway: razorpayPaymentGw,
// 	}
// 	newPayment.makePayment(100)
// }

// interfaces

package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gatway paymenter
}

type razorpay struct {

}

func (r razorpay) pay(amount float32){
	// logic to make payment
	fmt.Println("Making payment using razorpay", amount)
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{}

	// razorpayPaymentGw.pay(amount)
	p.gatway.pay(amount)
}


type paypal struct {}

func (p paypal) pay(amount float32){
	fmt.Println("Making payment using paypal", amount)
}




func main(){
	// razorpayPaymentGw := razorpay{}

	paypalGw := paypal{}

	newPayment := payment {
		gatway: paypalGw,
	}

	newPayment.makePayment(100)
}