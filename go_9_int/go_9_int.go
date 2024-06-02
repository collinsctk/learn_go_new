package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	// ---------------%T表示类型, %v表示值-----------------
	year := 2024
	fmt.Printf("Type %T for %v\n", year, year)

	a := "text"
	fmt.Printf("Type %T for %v\n", a, a)

	b := 42
	fmt.Printf("Type %T for %v\n", b, b)

	c := 3.14
	fmt.Printf("Type %T for %v\n", c, c)

	d := true
	fmt.Printf("Type %T for %v\n", d, d)

	// ---------------16进制-----------------
	var red, green, blue uint8 = 0, 141, 255
	var red16, green16, blue16 = 0x00, 0x8D, 0xD5

	fmt.Printf("Red: %d, Green: %d, Blue: %d\n", red, green, blue)
	fmt.Printf("Red: %d, Green: %d, Blue: %d\n", red16, green16, blue16)

	// ---------------打印16进制-----------------
	fmt.Printf("%02x%02x%02x\n", red, green, blue)

	// ---------------整数环绕-----------------
	var number_unit8 uint8 = 255
	number_unit8++
	fmt.Printf("%d\n", number_unit8)

	var number_int8 int8 = 127
	number_int8++
	fmt.Printf("%d\n", number_int8)

	// ---------------打印每一个bit-----------------
	var number_bit uint8 = 3
	fmt.Printf("%08b\n", number_bit)
	number_bit++
	fmt.Printf("%08b\n", number_bit)

	// ---------------打印时间-----------------
	future := time.Unix(12622780800, 0)
	fmt.Println(future)

	now := time.Now()
	fmt.Println(now)

	// -----------------------很大的数字-------------------------
	// ---------------一般超过int64的数字都用big包-----------------
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)
	fmt.Println(lightSpeed, secondsPerDay)
}
