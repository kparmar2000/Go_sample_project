package main

import "fmt"

// type structr struct {
// 	sName  string
// 	sage   int
// 	semail string
// }

// func function2() {
// 	//var sn, sa, se = s1.function() //error
// 	s1 := structr{
// 		sName: "karan", sage: 23, semail: "karan@gmail.com"} //error

// 	fmt.Println(s1.sName)
// 	fmt.Println(s1.sage)
// 	fmt.Println(s1.semail)
// 	// fmt.Println(sn)
// 	// fmt.Println(sa)
// 	// fmt.Println(se)

// }
// func main() {
// 	//s1.function();
// 	function2()
// }

type structr struct {
	sName  string
	sage   int
	semail string
}

func (s1 structr) function() (string, int, string) {
	s1.sName = "karan"
	s1.sage = 23
	s1.semail = "karanparmar@gmail.com"
	return s1.sName, s1.sage, s1.semail
}
func function2(s1 structr) {
	var sn, sa, se = s1.function() //error

	fmt.Println(sa)
	fmt.Println(sn)
	fmt.Println(se)

}
func main() {
	s1 := structr{}
	s1.function()
}
