/*
* @Author: bhargavkrishna
* @Date:   2016-02-02 13:35:33
* @Last Modified by:   bhargavkrishna
* @Last Modified time: 2016-02-02 13:41:47
 */

package main

import "fmt"

func main() {
	x := "A"
	a := 'A'
	for i := 0; i < 10; i++ {
		fmt.Println(x)
		x = x + "A"
	}
}
