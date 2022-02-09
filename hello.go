package main
import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

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
}
