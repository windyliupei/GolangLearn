package funcPackage

import "fmt"

func FuncDefine() {
	//函数可以在函数中声明

	funcA := func() {
		fmt.Println("I am...")
	}

	funcA()
}

func FuncDefine2() {
	//函数可以在函数中声明

	funcA := func(x, y int) int {
		return x + y
	}

	fmt.Println(funcA(1, 2))
}

/*
函数的闭包
*http://go-tour-zh.appspot.com/moretypes/21
*/
func Closure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {

	sum := 0

	return func(x int) int {
		sum += x
		return sum

	}

}

func fibonacci() func(int) int {
	prev2 := 0
	prev1 := 1
	curr := 1

	return func(x int) int {
		if x == 0 {
			return 0
		} else if x == 1 {
			return 1
		} else {
			curr = prev2 + prev1
			prev2 = prev1
			prev1 = curr

			return curr
		}

	}
}

func Closure2() {

	pos := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}
}
