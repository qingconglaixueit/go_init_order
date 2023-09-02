// @Author Bing 
// @Date 2023/9/2 12:27:00 
// @Desc
package main

import "fmt"
import _ "ttt/p1"
import _ "ttt/p2"

var (
	_ = checkConst()
	c = varInit(4)
	d = varInit(5)
)

const (
	a = 2
	b = 3
)

func init() {

	fmt.Println("main : init")

}

func checkConst() int {
	if a == 2 {
		fmt.Println("main : const a")
	}
	if b == 3 {
		fmt.Println("main : const b")
	}
	return 0
}

func varInit(x int) int {
	fmt.Println("main : var ", x)
	return x
}

func main() {
	fmt.Println("main : main ")
}
