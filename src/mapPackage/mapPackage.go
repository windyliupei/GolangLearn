package mapPackage

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

/*
map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。
*/

func MapInit() {

	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		40.8, -74.8,
	}

	fmt.Println(m["Bell Labs"])
}

func MapInit2() {

}
