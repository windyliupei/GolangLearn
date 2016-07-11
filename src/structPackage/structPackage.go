package structPackage

import "fmt"

//结构体
type Vertex struct {
	X int
	Y int
}

func StructFunc() {
	fmt.Println(Vertex{3, 2})

	v := Vertex{1, 3}
	v.X = 100
	fmt.Println(v.X)
}

func StructPoint() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func StructGrammar() {

	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}  // Y :0 被省略
		v3 = Vertex{}      //X:0 和 Y:0
		p  = &Vertex{1, 2} //类型为 *Vertex
	)

	fmt.Println(v1, p, v2, v3)

}
