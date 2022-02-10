package main
import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// 型定義
type Vertex struct {
	Lat, Long float64
}

// 宣言　mapでkey,valueの形を定義
var m map[string]Vertex

func main(){
	s := []int{2,3,5,7,11,13}
	fmt.Println(s)
	// 0 < x <= 4 なので　{3,5,7,11}
	s = s[:4]
	fmt.Println(s)
	// 0 < x <= 3 なので　{3,5,7}
	s = s[:3]
	fmt.Println(s)
	// 2 < x <= 3 なので　{7}
	s = s[2:]
	fmt.Println(s)

	// forEachと一緒、iは今の番号、vは数字
		for i, v := range pow {
		fmt.Printf("--------------\n")
		fmt.Printf("%d\n",i)
		fmt.Printf("%d\n",v)
		fmt.Printf("2**%d = %d\n", i, v)
	}

fmt.Printf("%v\n",m)
// 宣言　mapでkey,valueの形の配列を作成
	m = make(map[string]Vertex)
	fmt.Printf("%v\n",m)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	m["Mickel Jackson"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Mickel Jackson"])
	for key,value := range m{
		fmt.Println(key,value)
	}
}
