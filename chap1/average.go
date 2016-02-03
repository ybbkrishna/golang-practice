/*
* @Author: bhargavkrishna
* @Date:   2016-02-02 14:02:59
* @Last Modified by:   bhargavkrishna
* @Last Modified time: 2016-02-03 10:20:30
 */

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	avg := 0.0
	for _, v := range a {
		avg += float64(v)
	}
	avg = avg / float64(len(a))
	fmt.Println(avg)
}
