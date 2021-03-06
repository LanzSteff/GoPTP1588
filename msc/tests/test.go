package main

import (
	"fmt"
	"alex/ptp1588boundaryclock/communication"
	"alex/ptp1588boundaryclock/datasets"
	_"time"
)

func main() {
	i := 4
	s := "bla " + fmt.Sprintf("%v", i) + " end"
	fmt.Println(s)
	str := "8080"
	str += ":"
	fmt.Println(str)
	var iTest uint32 = 3231748232
	fmt.Println("test")
	fmt.Println(iTest)
	first := uint8((iTest >> 24) & 0xc0ff)
	second := uint8(iTest >> 16)
	third := uint8(iTest >> 8)
	fourth := uint8(iTest)
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
	fmt.Println(0xff)

	var b1 [3]byte
	var b2 [3]byte

	b1[0], b1[1], b1[2] = 0, 1, 2
	b2[0], b2[1], b2[2] = 3, 4, 5

	fmt.Printf("b1 %+v", b1)
	fmt.Printf("b2 %+v", b2)

	var msg communication.MessageType
	msg = communication.Announce

	b1[0] = byte(msg)
	fmt.Println(msg)

	vers := new(datasets.PortDS)
	vers.VersionNumber = 128

	b1[1] = vers.VersionNumber
	fmt.Println(b1)

	zahl := 0
	zahl += 1
	fmt.Println(zahl)
	zahl += (1<<1)
	zahl += (1<<2)
	fmt.Println(zahl)

	test := new(datasets.DefaultDS)
	fmt.Println(test)
	fmt.Println(*test)
	fmt.Println(&test)

	//time.Sleep(100 * time.Millisecond)

/*	car1, car2 := new(stuff.Car), new(stuff.Car)
	car1.Id = 10
	car2.Id = 20
	go stuff.TestCar(*car1, 1, 500)
	go stuff.TestCar(*car2, 2, 10)*/

	//time.Sleep(10 * time.Second)
	fmt.Println("Exit")

	var signed int8
	var unsigned uint8

	signed = -127
	fmt.Println(signed)
	unsigned = uint8(signed)
	fmt.Println(unsigned)

	fmt.Println(int8(unsigned))
}
