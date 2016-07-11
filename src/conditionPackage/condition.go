package condition

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

//if 语句除了没有了 `( )` 之外（甚至强制不能使用它们），看起来跟 C 或者 Java 中的一样，而 `{ }` 是必须的
func IfCondition(i int) {
	if i > 0 {
		fmt.Println("Input parameter > 0")
	} else { //注意else和'{','}'必须得这样写在同一行！
		fmt.Println("Input parameter < 0")
	}
}

//if 的便捷语句
// 跟 for 一样，`if` 语句可以在条件之前执行一个简单的语句。
//由这个语句定义的变量的作用域仅在 if 范围之内。
func IfExpress(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

//一个结构体（`struct`）就是一个字段的集合。
//除非以 fallthrough 语句结束，否则分支会自动终止。
func SwitchFunc() {
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

func SwitchFunc2() {

	fmt.Println("When's Saturday")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorry.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
