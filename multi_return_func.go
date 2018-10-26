package main
import "fmt"

func foo() (int, int, int){
	return 1,2,3
}

func main(){
	_, _, x := foo()
	fmt.Println(x)
}
