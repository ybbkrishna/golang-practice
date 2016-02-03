/*
* @Author: bhargavkrishna
* @Date:   2016-02-01 11:15:34
* @Last Modified by:   bhargavkrishna
* @Last Modified time: 2016-02-01 11:17:47
 */

package main

import "fmt"

func main() {
	var i int = 1
hello:
	if i > 10 {
		goto end
	}
	fmt.Println(i)
	i++
	goto hello
end:
}
