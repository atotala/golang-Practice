package main
import "fmt"

func foo(nums ...int) (int){
	sum := 0
	for _, num := range nums{
		sum = sum + num
	}

	return sum
}

func main(){

	s := make([]int, 3)
	s[0] = 1
	s[1] = 2
	s[2] = 3

fmt.Println(foo(1,2,3,4))
fmt.Println(foo(s...))
}
