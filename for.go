package main
import "fmt"

func main(){
	i := 0

	for i <=3{
		fmt.Println(i)
		i++
	}

	fmt.Println("Second Loop")

	for j := 0; j< 5; j++{
		fmt.Println(j)
	}

	fmt.Println("Third Loop")

	for{
		fmt.Println("Infinite Loop")
		break
	}


}
