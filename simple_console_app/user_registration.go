package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

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

	fmt.Println("Enter Last Name")
	fmt.Scan(&lastName)

	fmt.Println("Enter Email Address")
	fmt.Scan(&email)

	wg.Add(1)
	go sendOTPEmail(otp)

	fmt.Println("Enter Your Age")
	fmt.Scan(&age)

	fmt.Println("Enter Your City Name")
	fmt.Scan(&city)

	wg.Wait()

	fmt.Println("Enter OTP Code You Recieved")
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
	wg.Done()
}
