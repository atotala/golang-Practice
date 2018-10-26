package main
import "fmt"

func foo(a int, b int)int{
	return a + b
}

func bar(a,b int)int{
	return a +b
}

func main(){

fmt.Println(foo(1,2))
fmt.Println(bar(3,4))
}
