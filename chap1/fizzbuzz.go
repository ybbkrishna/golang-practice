/*
* @Author: bhargavkrishna
* @Date:   2016-02-01 11:32:37
* @Last Modified by:   bhargavkrishna
* @Last Modified time: 2016-02-02 13:30:10
 */

package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("Fizz Buzz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		}
	}
}
