/*
* @Author: bhargavkrishna
* @Date:   2016-02-01 11:25:00
* @Last Modified by:   bhargavkrishna
* @Last Modified time: 2016-02-01 11:26:46
 */

package main

import "fmt"

func main() {
	var i int = 0
	var arr [10]int
	for i = 0; i < 10; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
}
