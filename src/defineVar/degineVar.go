package defineVar

import (
	"fmt"
)

func init() {

}

//变量定义
func DefineVar() {
	var c, python, java bool
	var i int
	fmt.Println(i, c, python, java)
}

var i, j int = 1, 2

//Define Variable way 2
func DefineVar2() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

//在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。
//函数外的每个语句都必须以关键字开始（`var`、`func`、等等），`:=` 结构不能使用在函数外。
func DefineVariabel3() {
	k := 3
	fmt.Println(k)
}
