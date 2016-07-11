package baseType

import (
	"fmt"
	"math/cmplx"
)

//基本类型
var (
	Tobe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func BaseType() {
	const f = "%T(%v)\n"

	fmt.Println(f, Tobe, Tobe)
	fmt.Println(f, MaxInt, MaxInt)
	fmt.Println(f, z, z)
}
