package deferPackage

import "fmt"

//关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？因为 return 语句同样可以包含一些操作，而不是单纯地返回某个值）。
//关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源。

//defer
//defer 语句会延迟函数的执行直到上层函数返回。
//延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用
func DeferFun() {
	defer fmt.Println("World")

	fmt.Println("Hello ")
}

//defer 栈
//延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。
//阅读博文了解更多关于 defer 语句的信息。
func DeferFunc2() {

	fmt.Println("Defer Demo")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
