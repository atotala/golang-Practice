package main
import "fmt"

func main(){
	twoD := make([][]string, 3)

	for i := 0 ; i < 3; i++{
                twoD[i] = make([]string, i+1)
		for j := 0; j < i ; j++{
			twoD[i][j] = "a" 
		}
	}

	fmt.Println(twoD)
}
