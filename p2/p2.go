// @Author Bing 
// @Date 2023/9/2 12:32:00 
// @Desc
package p2

import "fmt"

var (

	_ = checkConst()
	c = varInit(4)
	d = varInit(5)
)

const(

	a = 2
	b = 3
)


func init(){

	fmt.Println("p2 : init")

}


func checkConst() int{
	if a == 2{
		fmt.Println("p2 : const a")
	}
	if b == 3{
		fmt.Println("p2 : const b")
	}
	return 0
}

func varInit(x int)int{
	fmt.Println("p2 : var ",x)
	return x
}