package main

import "fmt"

type car struct {
	company    string
	model      string
	cc         int
	frontwheel wheel
	rearwheel  wheel
}
type wheel struct {
	radius   int
	material string
}
type rect struct {
	width  int
	height int
}

var myCar = car{
	company: "honda",
	cc:      1800, model: "civic",
	frontwheel: wheel{10, "alloy"},
	rearwheel:  wheel{20, "metal"}}

func rectarea(r rect) int {
	return r.height * r.width
}
func (r rect) rectarea1() int {
	return r.height * r.width
}
func Function() {
	fmt.Println(myCar)
}
