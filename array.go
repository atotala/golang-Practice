package main
import "fmt"

func main(){
	var a [5]int
	a[4] =100

	b := [5]int{1,2,3,4,5}

	fmt.Println(a)
	fmt.Println(len(b))
	var twoD [2][4]int
	for i := 0 ; i< 2; i++{
		for j:=0 ; j< 4; j++{
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

}
