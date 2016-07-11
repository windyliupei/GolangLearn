package constPackage

import "fmt"

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
