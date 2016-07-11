package typeOperation

import "fmt"

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
