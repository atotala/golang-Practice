package main
import "fmt"

func fact(a int)int{
	if a == 0{
		return 1
	}else{
		return a*fact(a-1)
	}

}

func main(){

fmt.Println(fact(5))
}
