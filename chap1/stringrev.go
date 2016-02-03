/*
* @Author: bhargavkrishna
* @Date:   2016-02-02 13:57:42
* @Last Modified by:   bhargavkrishna
* @Last Modified time: 2016-02-02 14:02:08
 */

package main

import "fmt"

func main() {
	str := "hello world"
	runes := []rune(str)
	for i, j := 0, len(str)-1; i <= j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes))
}
