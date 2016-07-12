package rangePackage

import "fmt"

//http://go-tour-zh.appspot.com/moretypes/12
func RangAppendSlice() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2** %d = %d \n", i, v)
	}
}

func RangAppendSlice2() {
	pow := make([]int, 10)

	for i := range pow {

		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	fmt.Printf("d%\n", 1<<3)
}
