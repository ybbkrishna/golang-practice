/*
* @Author: bhargavkrishna
* @Date:   2016-02-02 13:46:37
* @Last Modified by:   bhargavkrishna
* @Last Modified time: 2016-02-02 13:55:24
 */

package main

import "fmt"

func main() {
	str := "asSASA ddd dsjkdsjs dk"
	runes := []rune(str)
	fmt.Println(len(runes))
	copy(runes[4:7], []rune("abc"))
	fmt.Println(string(runes))
}
