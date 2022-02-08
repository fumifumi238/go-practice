package main
import "fmt"
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
}


