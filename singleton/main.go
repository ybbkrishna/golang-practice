package main

import (
	"fmt"

	"github.com/ybbkrishna/golang-practice/singleton/serv1"
	"github.com/ybbkrishna/golang-practice/singleton/serv2"
	"github.com/ybbkrishna/golang-practice/singleton/serv3"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	for i := 0; i < 10000; i++ {
		go testing()
	}
	// go func() {
	// 	defer func() {
	// 		if err := recover(); err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	}()
	// 	fmt.Printf("%p\n", serv3.GetService())
	// 	panic("lll1")
	// }()
	// go func() {
	// 	defer func() {
	// 		if err := recover(); err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	}()
	// 	fmt.Printf("%p\n", serv3.GetService())
	// 	panic("lll12")
	// }()
	// go func() {
	// 	// time.Sleep(2 * time.Second)
	// 	fmt.Printf("%p\n", serv3.GetService())
	// }()
	// go func() {
	// 	// time.Sleep(time.Second * 10)
	// 	fmt.Printf("%p\n", serv3.GetService())
	// }()
	select {}
}

func testing() {
	if serv1.GetService() == nil {
		fmt.Println("lll")
	}
	if serv2.GetService() == nil {
		fmt.Println("lll")
	}
	if serv3.GetService() == nil {
		fmt.Println("lll")
	}
	// fmt.Println(serv1.GetService())
	// fmt.Println(serv2.GetService())
	// fmt.Println(serv3.GetService())
}
