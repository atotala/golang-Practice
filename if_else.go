package main
import "fmt"

func main(){
	i := 5
	str := "ANAND"

	if i%5 == 0{
		fmt.Println("Divisble by 5")
	}

	if i%2 == 0{
		fmt.Println("Divisible by 2")
	} else{
		fmt.Println("Not Divisible by 2")
	}

	if str == "ABCD"{
		fmt.Println(str)
	}else if str == "EFGH"{
		fmt.Println(str)
	}else{
		fmt.Println(str)
	}


}
