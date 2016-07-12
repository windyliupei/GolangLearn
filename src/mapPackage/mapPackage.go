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

func LoopKeyValue() {

	var m = map[string]Vertex{
		"Bell Labs": Vertex{1, 2},
		"Google":    Vertex{},
	}

	//fmt.Println(m)

	//loop values of map
	for _, value := range m {
		fmt.Println(value)
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

	for key, _ := range m {
		fmt.Println(key)
	}

}

func DeleteItemFromMap() {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{1, 2},
		"Google":    Vertex{},
	}

	delete(m, "Google")

	fmt.Println(m)

}

func GetItem(key string) Vertex {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{1, 2},
		"Google":    Vertex{9, 9},
	}

	if value, ok := m[key]; ok {
		return value
	}

	return Vertex{0, 0}

}
