package main

import "fmt"

func main() {
	var a uint8
	var b uint8
	var c uint16

	a = 0x12
	b = 0x34
	c = 0x0000

	c |= uint16(a << 8)
	c |= uint16(b << 0)

	fmt.Printf("c is 0x%04X\n", c)
}
