package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var firstName string
	var lastName string
	var email string
	var age int8
	var city string
	var otp int = rand.Intn(9999)
	var otpTyped int

	fmt.Println("Enter First Name")
	fmt.Scan(&firstName)

	fmt.Println("Enter First Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter Email Name")
	fmt.Scan(&email)
	sendOTPEmail(otp)

	fmt.Println("Enter Age Name")
	fmt.Scan(&age)

	fmt.Println("Ënter City")
	fmt.Scan(&city)

	fmt.Println("Enter OTP code Recieved")
	fmt.Scan(&otpTyped)

	if otp == otpTyped {
		fmt.Println("Registration Successfully Completed")
	} else {
		fmt.Println("Registration failed")
	}

}

func sendOTPEmail(otp int) {
	time.Sleep(10 * time.Second)
	fmt.Println("Your OTP code is : ", otp)
	return
}
