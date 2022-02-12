package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
	// X,Yのfloat64の型を持つ　少数整数16桁まで
}

// classにあたる。　Absという関数はvという名前のVertex型の要素を持つ
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// func Scale(v *Vertex, f float64)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
		// X:3,Y:4の初期値を設定、vのpointerを持つものは共通して3,4という値を持てるようになった（classと同じ）
	v := Vertex{3, 4}
// classの呼び出し
	v.Scale(10)
	fmt.Println(v.Abs())
}
