package main

//import fmt "fmt" // Package implementing formatted I/O.
//import mat "math" //倒入包，并以'mat'的名称使用这个包

//倒入包方式2
import (
	"fmt"
	"math"
	"math/cmplx"
	"runtime"
	"time"

	"defineVar"
)

//import math "math"

func main() {

	defineVar.DefineVar()
	defineVar.DefineVar2()
	defineVar.DefineVar()
	//baseType()
	//TypeConvert()
	//TypeGuess()
	//ConstFun()
	//forloop()
	//orloop2()
	//ifCondition(1)
	//fmt.Println(ifExpress(2, 4, 10000))
	//switchFunc()
	//switchFunc2()
	//deferFun()
	//deferFunc2()
	//pointFun()
	//structFunc()
	//structPoint()
	//structGrammar()
	//arrayFunc()
	//sliceFunc()
	//sliceFuncCut()
	//CreateSlice()
	//slicenil()
	appendSlice()

}

var i, j int = 1, 2

//基本类型
var (
	Tobe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func baseType() {
	const f = "%T(%v)\n"

	fmt.Println(f, Tobe, Tobe)
	fmt.Println(f, MaxInt, MaxInt)
	fmt.Println(f, z, z)
}

//类型转换
func TypeConvert() {

	var i int = 3
	var f float32 = float32(i)
	var u uint = uint(f)

	fmt.Println(i, f, u)

}

//类型推导
func TypeGuess() {
	i := 32
	fmt.Printf("i Type is %T \n", i) //printf 带有格式的输出
}

//常亮设置
//Big 不可以直接修改
const (
	Big   = 1 << 99
	Small = Big >> 99
)

func ConstFun() {

	fmt.Println(needFloat(Big))

}
func needFloat(x float64) float64 {
	return x * 0.1
}

//for 循环
//基本的 for 循环除了没有了 `( )` 之外（甚至强制不能使用它们），看起来跟 C 或者 Java 中做的一样，而 `{ }` 是必须的。

func forloop() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//跟 C 或者 Java 中一样，可以让前置、后置语句为空。
//for 是 Go 的 “while” 基于此可以省略分号：C 的 while 在 Go 中叫做 `for`。
func forloop2() {

	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}

//死循环
//如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环。
func deadLoop() {
	for {
	}
}

//if 语句除了没有了 `( )` 之外（甚至强制不能使用它们），看起来跟 C 或者 Java 中的一样，而 `{ }` 是必须的
func ifCondition(i int) {
	if i > 0 {
		fmt.Println("Input parameter > 0")
	} else { //注意else和'{','}'必须得这样写在同一行！
		fmt.Println("Input parameter < 0")
	}
}

//if 的便捷语句
// 跟 for 一样，`if` 语句可以在条件之前执行一个简单的语句。
//由这个语句定义的变量的作用域仅在 if 范围之内。
func ifExpress(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

//一个结构体（`struct`）就是一个字段的集合。
//除非以 fallthrough 语句结束，否则分支会自动终止。
func switchFunc() {
	fmt.Println("GO Runs on")

	switch os := runtime.GOOS; os {
	case "drawin": //os == drawin
		{
			fmt.Println("OS X")
		}
	case "linux": //os == linux
		{
			fmt.Println("Linux")
		}
	default: //os == others
		{
			fmt.Println("Others OS")
		}
	}
}

func switchFunc2() {

	fmt.Println("When's Saturday")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorry.")
	case today + 2:
		fmt.Println("In tow days.")
	default:
		fmt.Println("Too far away.")
	}
}

//defer
//defer 语句会延迟函数的执行直到上层函数返回。
//延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用
func deferFun() {
	defer fmt.Println("World")

	fmt.Println("Hello ")
}

//defer 栈
//延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。
//阅读博文了解更多关于 defer 语句的信息。
func deferFunc2() {

	fmt.Println("Defer Demo")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}

/*
 Go 具有指针。 指针保存了变量的内存地址。
 类型 *T 是指向类型 T 的值的指针。其零值是 `nil`。
 var p *int
 & 符号会生成一个指向其作用对象的指针。
	i := 42
	p = &i
* 符号表示指针指向的底层的值。
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
这也就是通常所说的“间接引用”或“非直接引用”。
*/

func pointFun() {
	i, j = 42, 2701

	p := &i //& 符号会生成一个指向其作用对象的指针。
	fmt.Println(*p)

	*p = 21 //* 符号表示指针指向的底层的值。
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

}

//结构体
type Vertex struct {
	X int
	Y int
}

func structFunc() {
	fmt.Println(Vertex{3, 2})

	v := Vertex{1, 3}
	v.X = 100
	fmt.Println(v.X)
}

func structPoint() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func structGrammar() {

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}  // Y :0 被省略
		v3 = Vertex{}      //X:0 和 Y:0
		p  = &Vertex{1, 2} //类型为 *Vertex
	)

	fmt.Println(v1, p, v2, v3)

}

//数组
//类型 [n]T 是一个有 n 个类型为 T 的值的数组。
//数组的长度是其类型的一部分，因此数组不能改变大小。
func arrayFunc() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

//slice
//[]T 是一个元素类型为 T 的 slice。
func sliceFunc() {

	p := []int{2, 3, 5, 7, 11, 23}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

//s[lo:hi] 表示从 lo 到 hi-1 的 slice 元素，含两端
func sliceFuncCut() {
	p := []int{2, 3, 5, 7, 11, 13}

	fmt.Println("P == ", p)
	fmt.Println("p[1:4]", p[1:4]) //从0开始， 【1】 是3，【4 -1】是7

	//省略下标代表从 0 开始
	fmt.Println("p[:3]", p[:3]) //从0开始， 【1】 是3，【3-1】是5

	//省略上标代表到 len(s) 结束
	fmt.Println("p[4：]", p[4:])
}

func CreateSlice() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

}
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x) //cap()函数返回的是数组切片分配的空间大小
}

func slicenil() { //一个 nil 的 slice 的长度和容量是 0。
	var z []int
	fmt.Println(z, len(z), cap(z)) //容量
	if z == nil {
		fmt.Println("nil!")
	}
}

//The zero value of a slice is nil. The len and cap functions will both return 0 for a nil slice.
func appendSlice() {
	var a []int

	printSlice("a", a)
	//如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组。
	a = append(a, 1)
	a = append(a, 2)

	printSlice("a", a)
}
