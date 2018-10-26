package main
import "fmt"

func f1(a *int){
	*a = 0
}

func f2(a int){
	a = 10
}

func main(){
	i := 5
	f2(i)
	fmt.Println(i)
	f1(&i)
	fmt.Println(i)
}
