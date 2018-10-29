package main
import "fmt"

func main(){
	type record struct{
		name string
		age int
	}

	r1 := record{"anand", 25}
	fmt.Println(r1)

	r2 := &record{"gourav", 24}

	fmt.Println(r2)

	fmt.Println(r2.name)
	fmt.Println(r2.age)
}
