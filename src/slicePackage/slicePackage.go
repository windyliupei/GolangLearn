package slicePackage

import "fmt"

//数组
//类型 [n]T 是一个有 n 个类型为 T 的值的数组。
//数组的长度是其类型的一部分，因此数组不能改变大小。
func ArrayFunc() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

//slice
//[]T 是一个元素类型为 T 的 slice。
func SliceFunc() {

	p := []int{2, 3, 5, 7, 11, 23}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}
}

//s[lo:hi] 表示从 lo 到 hi-1 的 slice 元素，含两端
func SliceFuncCut() {
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

func Slicenil() { //一个 nil 的 slice 的长度和容量是 0。
	var z []int
	fmt.Println(z, len(z), cap(z)) //容量
	if z == nil {
		fmt.Println("nil!")
	}
}

//The zero value of a slice is nil. The len and cap functions will both return 0 for a nil slice.
func AppendSlice() {
	var a []int

	printSlice("a", a)
	//如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组。
	a = append(a, 1)
	a = append(a, 2)

	printSlice("a", a)
}
