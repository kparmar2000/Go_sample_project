package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	//"strings"
// )

// func main() {
// 	res, err := http.Get("http://www.google.com/robots.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer res.Body.Close()
// 	stream, err := ioutil.ReadAll(res.Body)
// 	//bot,err:= os.Readlink(res.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	//fmt.Printf("%s", stream)
// 	fmt.Println(string(stream))
// 	fmt.Println("start")
// 	defer func() {

// 		if err := recover(); err != nil {
// 			fmt.Printf("error occured % v ", err)
// 		}

// 		//fmt.Println(err)
// 	}()
// 	defer fmt.Println("defering this line")
// 	panic("something happened")
// 	fmt.Println("end")
// 	fmt.Println("hello it is working")
// }
