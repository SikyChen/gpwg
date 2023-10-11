package main

import (
	"fmt"
	"math/rand"
)

func main() {
	func1()
	func2()
	func3()
	func4()
	func5()
}

func func1() {
	const lightSpeed = 299792
	var distance = 56000000

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000

	fmt.Println(distance/lightSpeed, "seconds")

	fmt.Println(5 / 2)
}

func func2() {
	const hoursPerDay = 24
	var speed = 100800
	var distance = 96300000

	fmt.Println(distance/speed/hoursPerDay, "days")
}

func func3() {
	var hoursPerDay, minutesPerHours = 24, 60

	fmt.Printf("一天有%v小时，每小时有%v分钟\n", hoursPerDay, minutesPerHours)
}

func func4() {
	var a, b = 1, 2
	a *= 2
	b++

	fmt.Println("a is", a, ", b is", b)
}

func func5() {
	// rand number
	var num = rand.Intn(10) + 1
	fmt.Println("Rand number is", num)

	num = rand.Intn(10) + 1
	fmt.Println("Rand number is", num)

	// 介于 56,000,000 至 401,000,000 之间的数
	var distance = rand.Intn(401000000-56000000+1) + 56000000
	fmt.Println("Rand distance is", distance)
}

func lesson() {
	const distance = 56000000
	var days = 28
	var hoursPerDay = 24

	var speed = distance / days / hoursPerDay
	fmt.Println("Speed is", speed)
}
